package models

import "os"

var JWTSalt = "salt"

func init() {
	salt, ok := os.LookupEnv("JWT_SALT")
	if ok {
		JWTSalt = salt
	}
}

type BaseConfig struct {
	Host       string `env:"HOST" envDefault:":1234"`
	ServerName string `env:"SERVER_NAME" envDefault:"TestServer"`
	SentryDSN  string `env:"SENTRY_DSN" envDefault:"https://906c97ffd8b940e6a1223e353d1c11eb@sentry.devops.rusklimat.net/15"`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"trace"`
	TLSKey     string `env:"TLS_KEY" envDefault:""`
	TLSPem     string `env:"TLS_PEM" envDefault:""`
}
