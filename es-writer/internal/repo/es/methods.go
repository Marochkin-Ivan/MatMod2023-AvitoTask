package es

import (
	"context"
	"es-writer/pkg/errs"
	"github.com/caarlos0/env/v8"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func New() (*Client, *errs.Error) {
	const source = "es.New"

	var cfg cfg
	e := env.Parse(&cfg)
	if e != nil {
		return nil, errs.NewError(logrus.ErrorLevel, e.Error()).Wrap(source)
	}

	var res Client
	res.c, e = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{cfg.Addr},
		Username:  cfg.User,
		Password:  cfg.Pass,
	})
	if e != nil {
		return nil, errs.NewError(logrus.ErrorLevel, e.Error()).Wrap(source)
	}

	res.cfg = cfg

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ping, e := res.c.Info(res.c.Info.WithContext(ctx))
	if e != nil {
		return nil, errs.NewError(logrus.ErrorLevel, e.Error()).Wrap(source)
	}
	defer ping.Body.Close()

	// Проверьте статус-код ответа
	if ping.IsError() {
		return nil, errs.NewError(logrus.ErrorLevel, ping.Status()).Wrap(source)
	}

	log.Println("Elastic connection Status:", ping.Status())

	return &res, nil
}
