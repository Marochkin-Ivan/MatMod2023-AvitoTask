package cache

import "github.com/go-redis/redis/v8"

type Connections []*redis.Client

type config struct {
	addr     string `env:"REDIS_ADDR" envDefault:""`
	user     string `env:"REDIS_USER" envDefault:""`
	password string `env:"REDIS_PASS" envDefault:""`
}

const (
	// Info
	//
	// key: vacancy_id (string)
	//
	//value:
	Info = iota

	// Keywords
	//
	//key: vacancy_id (string)
	//
	//value:
	Keywords

	tablesCount
)

var DBName = map[int]string{
	0: "Info",
	1: "Keywords",
}
