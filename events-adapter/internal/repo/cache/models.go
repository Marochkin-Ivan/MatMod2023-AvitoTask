package cache

import "github.com/go-redis/redis/v8"

type Connections []*redis.Client

type config struct {
	addr     string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	user     string `env:"REDIS_USER" envDefault:""`
	password string `env:"REDIS_PASS" envDefault:""`
}

const (
	// UserEvents
	//
	// key: user_id (string)
	//
	// value: timestamp(int64)
	UserEvents = iota

	// UserSearches
	//
	// key: user_id (string)
	//
	// value: timestamp(int64)
	UserSearches

	tablesCount
)

var DBName = map[int]string{
	0: "UserEvents",
	1: "UserSearches",
}
