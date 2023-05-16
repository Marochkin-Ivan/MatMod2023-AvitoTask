package es

import "github.com/elastic/go-elasticsearch/v8"

type cfg struct {
	Addr string `env:"ES_ADDR" envDefault:""`
	User string `env:"ES_USER" envDefault:""`
	Pass string `env:"ES_PASS" envDefault:""`
}

type Client struct {
	c   *elasticsearch.Client
	cfg cfg
}
