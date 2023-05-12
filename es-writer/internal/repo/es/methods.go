package es

import (
	"es-writer/pkg/errs"
	"github.com/caarlos0/env/v8"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sirupsen/logrus"
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

	return &res, nil
}
