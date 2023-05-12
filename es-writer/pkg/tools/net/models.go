package net

import (
	"golang.org/x/time/rate"
	"net/http"
)

// limiterRoundTripper - структура, реализующая интерфейс http.RoundTripper
type limiterRoundTripper struct {
	rt  http.RoundTripper // Основной http.RoundTripper
	lim *rate.Limiter     // Лимитер для контроля скорости отправки запросов
}
