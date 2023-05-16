package es

import "github.com/elastic/go-elasticsearch/v8"

type cfg struct {
	Addr string `env:"ES_ADDR" envDefault:"http://178.216.97.177:9200"`
	User string `env:"ES_USER" envDefault:"elastic"`
	Pass string `env:"ES_PASS" envDefault:"$DcZd[Ht94Q4e7$$"`
}

type Client struct {
	c   *elasticsearch.Client
	cfg cfg
}
