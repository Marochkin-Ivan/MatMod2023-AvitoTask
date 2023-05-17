package cache

import (
	"api/pkg/errs"
	"context"
	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"log"
)

func New() (*Client, *errs.Error) {
	const source = "cache.New"

	var cfg cfg
	e := env.Parse(&cfg)
	if e != nil {
		return nil, errs.NewError(logrus.FatalLevel, e.Error()).
			WrapWithSentry(source, errs.SentryCategoryFunc, nil)
	}

	var cs []*redis.Client
	for i := 0; i < tableCount; i++ {
		cs = append(cs,
			redis.NewClient(&redis.Options{
				Addr:     cfg.Addr,
				Username: cfg.User,
				Password: cfg.Pass,
				DB:       i,
			}))
		log.Println(cs[i].Ping(context.Background()))
	}

	return &Client{cs, cfg}, nil
}
