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
	SentryDSN  string `env:"SENTRY_DSN" envDefault:""`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"trace"`
	TLSKey     string `env:"TLS_KEY" envDefault:""`
	TLSPem     string `env:"TLS_PEM" envDefault:""`
}

type BaseEventRequest struct {
	UserID    string `json:"user_id"`
	VacancyID string `json:"vacancy_id"`
	Type      string `json:"type"`
}

type SearchEventRequest struct {
	BaseEventRequest
	SearchPhrases string `json:"search_phrases"`
}

type RedisEvent struct {
	VacancyID string `json:"vacancy_id"`
	Type      string `json:"type"`
}

type RedisSearch struct {
	Phrases []string `json:"phrases"`
}
