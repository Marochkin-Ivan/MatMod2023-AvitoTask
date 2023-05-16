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
	// value: models.RedisEvent (json-string)
	UserEvents = 2

	tablesCount
)

var DBName = map[int]string{
	2: "UserEvents",
}
