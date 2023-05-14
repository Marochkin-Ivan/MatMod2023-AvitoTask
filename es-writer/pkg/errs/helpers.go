package errs

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
)

var values = map[string]logrus.Level{"trace": logrus.TraceLevel, "debug": logrus.DebugLevel, "info": logrus.InfoLevel, "warn": logrus.WarnLevel, "error": logrus.ErrorLevel, "panic": logrus.PanicLevel}

func ParseToLogrus(lvl string) (logrus.Level, *Error) {
	const source = "ParseToLogrus"

	lvl = strings.ToLower(lvl)
	val, found := values[lvl]
	if !found {
		errText := fmt.Sprintf("unknown logging level %s.  level name's %v", lvl, values)
		return 0, NewError(logrus.FatalLevel, errText).Wrap(source)
	}
	return val, nil
}

func LogWatcher(lch LogChan, logger *logrus.Logger, wg *sync.WaitGroup) {
	defer wg.Done()
	for l := range lch {
		consolePrint(l, logger)

		// send to sentry
		if l.GetLevel() <= logrus.WarnLevel {

			// if log from handler
			if h, ok := l.err.(HandlerLog); ok {
				if l.GetLevel() == logrus.WarnLevel {
					sendHandlerLogToSentry(sentry.LevelInfo, h)
				} else {
					sendHandlerLogToSentry(sentry.LevelError, h)
				}

				continue
			}

			// if default log
			sentry.CaptureEvent(l.sentryEvent)
		}
	}
}

func consolePrint(l *Error, logger *logrus.Logger) {
	switch l.GetLevel() {
	case logrus.TraceLevel:
		logger.Traceln(l.Error())
	case logrus.DebugLevel:
		logger.Debugln(l.Error())
	case logrus.InfoLevel:
		logger.Infoln(l.Error())
	case logrus.WarnLevel:
		logger.Warnln(l.Error())
	case logrus.ErrorLevel:
		logger.Errorln(l.Error())
	case logrus.FatalLevel:
		logger.Fatalln(l.Error())
	case logrus.PanicLevel:
		logger.Panicln(l.Error())
	}
}

func RequestBody4Sentry(requestBody []byte) map[string]any {
	res := make(map[string]any)
	res["request"] = string(requestBody)
	return res
}

func LocalLogWatcher(ch, globalCh LogChan, funcName string, funcType SentryCategory, input map[string]any, mostCriticalError *Error, wt *sync.WaitGroup) {
	defer wt.Done()

	canBeMostCritical := &Error{level: logrus.TraceLevel}
	var hadE = false

	for e := range ch {
		if e == nil {
			continue
		}

		if e.GetLevel() <= canBeMostCritical.GetLevel() {
			canBeMostCritical = e
			hadE = true
		}

		globalCh <- e.
			WrapWithSentry(funcName, funcType, input)
	}

	if hadE {
		mostCriticalError = canBeMostCritical
	} else {
		mostCriticalError = nil
	}
}
