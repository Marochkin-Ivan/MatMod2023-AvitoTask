package server

import (
	"errors"
	"es-writer/internal/repo/es"
	"es-writer/pkg/errs"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func NewServer(opts ...func(*Server)) *Server {
	s := new(Server)

	for _, o := range opts {
		o(s)
	}
	return s
}

func WithApp(a *fiber.App) func(*Server) {
	return func(s *Server) {
		s.a = a
	}
}

func WithLogChan(l errs.LogChan) func(*Server) {
	return func(s *Server) {
		s.logs = l
	}
}

func WithConfig(cfg Config) func(*Server) {
	return func(s *Server) {
		s.cfg = cfg
	}
}

func WithElasticSearch(cli *es.Client) func(*Server) {
	return func(s *Server) {
		s.es = cli
	}
}

func WithCache(c Setter) func(*Server) {
	return func(s *Server) {
		s.cache = c
	}
}

func (s *Server) SetupHandlers() *Server {
	v1 := s.a.Group("/api/v1/search")

	v1.Get("/ping", s.ping)

	// for example
	v1.Get("/create/index", s.create)
	v1.Get("/insert/document", s.insert)
	v1.Get("/insert/documents", s.insertBatch)
	v1.Get("/fill/cache", s.fillRedis)

	v1.Get("/test", s.test)

	return s
}

// Listen запускает соответсвующую конфигурацию для http\https
func (s *Server) Listen() error {
	var err error

	if (s.cfg.TLSPem != "" && s.cfg.TLSKey == "") || (s.cfg.TLSPem == "" && s.cfg.TLSKey != "") {
		return errors.New("(*Server).Listen() error: cfg.TLSPem or cfg.TLSKey doesn't have value")
	}

	switch s.cfg.TLSPem != "" && s.cfg.TLSKey != "" {
	case true:
		err = s.a.ListenTLS(s.cfg.Host, s.cfg.TLSPem, s.cfg.TLSKey)

	case false:
		err = s.a.Listen(s.cfg.Host)
	}

	return err
}

// GracefulShutdown - перехватывает syscall'ы и выполняет штатную остановку приложения
func (s *Server) GracefulShutdown(connectionsClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	<-sigint

	if err := s.a.Shutdown(); err != nil {
		s.logs <- errs.NewError(logrus.FatalLevel, err.Error()) // wrap not required
	}

	connectionsClosed <- struct{}{}
}
