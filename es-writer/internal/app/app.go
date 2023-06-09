package app

import (
	"es-writer/internal/repo/cache"
	"es-writer/internal/repo/es"
	"es-writer/internal/server"
	"es-writer/pkg/errs"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

func Start() {
	var lg = logrus.New()
	lg.Formatter = new(logrus.JSONFormatter)

	cfg, cusErr := server.GetConfig()
	if cusErr != nil {
		lg.Fatal(cusErr)
	}

	lvl, cusErr := errs.ParseToLogrus(cfg.LogLevel)
	if cusErr != nil {
		lg.Fatal(cusErr.Error())
	}
	lg.SetLevel(lvl)

	esCli, cusErr := es.New()
	if cusErr != nil {
		lg.Fatal(cusErr.Error())
	}

	conns, cusErr := cache.New()
	if cusErr != nil {
		lg.Fatal(cusErr.Error())
	}
	defer cache.Close(conns)

	logChan := make(errs.LogChan, 1000)
	fiberLg := errs.NewFiberLogger(logChan)
	s := server.NewServer(
		server.WithApp(server.NewFiberApp(fiberLg)),
		server.WithConfig(cfg),
		server.WithLogChan(logChan),
		server.WithElasticSearch(esCli),
		server.WithCache(conns),
	).SetupHandlers()

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go errs.LogWatcher(logChan, lg, wg)

	connectionsClosed := make(chan struct{})
	go s.GracefulShutdown(connectionsClosed)

	if err := s.Listen(); err == http.ErrServerClosed {
		logChan <- errs.NewError(logrus.FatalLevel, err.Error()) // wrap not required
	} else {
		<-connectionsClosed // wait 4 success close connections
	}

	// close db connections

	close(logChan)
	wg.Wait()
}
