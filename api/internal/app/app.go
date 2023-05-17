package app

import (
	"api/internal/repo/cache"
	"api/internal/repo/es"
	"api/internal/server"
	"api/pkg/errs"
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

	cacheCl, err := cache.New()
	if err != nil {
		lg.Fatal()
	}

	esCli, cusErr := es.New()
	if cusErr != nil {
		lg.Fatal(cusErr.Error())
	}

	logChan := make(errs.LogChan, 1000)
	fiberLg := errs.NewFiberLogger(logChan)
	s := server.NewServer(
		server.WithApp(server.NewFiberApp(fiberLg)),
		server.WithConfig(cfg),
		server.WithLogChan(logChan),
		server.WithElasticSearch(esCli),
		server.WithCache(cacheCl),
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
