package cache

import (
	"events-adapter/pkg/errs"
	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func getConfig() (config, *errs.Error) {
	const source = "getConfig"

	var cfg config
	err := env.Parse(&cfg)
	if err != nil {
		return config{}, errs.NewError(logrus.ErrorLevel, err.Error()).Wrap(source)
	}

	return cfg, nil
}

func New() (Connections, *errs.Error) {
	const source = "NewConnections"

	c, err := getConfig()
	if err != nil {
		return Connections{}, err.WrapWithSentry(source, errs.SentryCategoryFunc, nil)
	}

	cons := make(Connections, tablesCount)
	for idx := range cons {
		cons[idx] = redis.NewClient(&redis.Options{
			Username: c.user,
			Password: c.password,
			Addr:     c.addr,
			DB:       idx,
		})
	}

	return cons, nil
}

func Close(conns Connections) {
	for _, conn := range conns {
		conn.Close()
	}
}
