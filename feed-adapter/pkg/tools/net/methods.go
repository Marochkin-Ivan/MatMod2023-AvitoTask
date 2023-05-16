package net

import (
	"context"
	"golang.org/x/time/rate"
	"net/http"
)

// RoundTrip - метод структуры limiterRoundTripper, реализующий интерфейс http.RoundTripper
func (lrt *limiterRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if err := lrt.lim.Wait(context.Background()); err != nil {
		return nil, err
	}

	return lrt.rt.RoundTrip(req)
}

// NewRoundTripperWithLimiter возвращает обертку над переданным интерфейсом http.RoundTripper, который
// позволяет установить ограничение на RPS для обращений к сторонним сервисам
func NewRoundTripperWithLimiter(rt http.RoundTripper, lim *rate.Limiter) http.RoundTripper {
	return &(limiterRoundTripper{rt: rt, lim: lim})
}
