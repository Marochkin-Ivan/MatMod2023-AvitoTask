package cache

import "github.com/go-redis/redis/v8"

const (
	DetailInfoTable = iota // all info about vacancy
	KeyWordsTable
	EventsTable

	tableCount // must be last in list
)

var DBName = map[int]string{
	DetailInfoTable: "DetailInfo",
	KeyWordsTable:   "KeyWords",
	EventsTable:     "Events",
}

type Client struct {
	cs  []*redis.Client
	cfg cfg
}

type cfg struct {
	Addr string `env:"REDIS_ADDR" envDefault:"178.216.99.30:6379"`
	User string `env:"REDIS_USER" envDefault:""`
	Pass string `env:"REDIS_PASS" envDefault:"94?gqFWn1XhI5(+^"`
}
