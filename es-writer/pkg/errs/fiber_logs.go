package errs

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

type Logger4Fiber struct {
	ch LogChan
}

type HandlerLog struct {
	Status   int    `json:"status"`
	Duration any    `json:"duration"`
	Method   string `json:"method"`
	Path     string `json:"path"`
	Response any    `json:"resp"`
}

func (h HandlerLog) toMessage() string {
	return fmt.Sprintf("%s %s %d with duration %v resp: %s", h.Method, h.Path, h.Status, h.Duration, h.Response)
}

func (h HandlerLog) Error() string {
	return h.toMessage()
}

func (e Logger4Fiber) Write(b []byte) (int, error) {
	const source = "logger4Fiber.Write"

	ns := strings.ReplaceAll(string(b), "\"{", "{")
	ns = strings.ReplaceAll(ns, "}\"", "}")
	var h HandlerLog
	err := json.Unmarshal([]byte(ns), &h)
	if err != nil {
		e.ch <- NewError(logrus.WarnLevel, err.Error()).Wrap(source)
		return 0, err
	}

	var efkLvl logrus.Level

	switch h.Status {
	case http.StatusInternalServerError, http.StatusInsufficientStorage, http.StatusBadGateway, http.StatusGatewayTimeout:
		efkLvl = logrus.ErrorLevel

	case http.StatusNotFound, http.StatusBadRequest, http.StatusConflict, http.StatusForbidden:
		efkLvl = logrus.WarnLevel

	default:
		efkLvl = logrus.InfoLevel
	}

	e.ch <- &Error{level: efkLvl, err: h}
	return 0, nil
}

func sendHandlerLogToSentry(lvl sentry.Level, h HandlerLog) {
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetLevel(lvl)
		scope.AddBreadcrumb(&sentry.Breadcrumb{
			Type:      "http",
			Category:  "URL Endpoints",
			Message:   h.toMessage(),
			Level:     lvl,
			Timestamp: time.Now(),
		}, 10)

		sentry.CaptureMessage(h.toMessage())
		scope.Clear()
	})
}

func NewFiberLogger(ch LogChan) Logger4Fiber {
	return Logger4Fiber{ch: ch}
}
