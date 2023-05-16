package errs

import (
	"errors"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type LogChan chan *Error
type Error struct {
	level       logrus.Level
	err         error
	sentryEvent *sentry.Event
}

/*
Wrap (deprecated) Оборачивает ошибку в переданное название функции, т.е.:

err = models.NewError(someLevel, "can't read file file.txt")

err.Wrap("ReadFile")

fmt.Println(err.Err.Error()) -> in ReadFile: cant read file file.txt
*/
func (e *Error) Wrap(callerName string) *Error {
	e.err = errors.New(" in " + callerName + ": " + e.err.Error())
	return e
}

// popStacktrace удаляет последний вызов функции из стека
func popStacktrace(s *sentry.Stacktrace) *sentry.Stacktrace {
	s.Frames = s.Frames[:len(s.Frames)-1]
	return s
}

func getCallerName(s *sentry.Stacktrace) string {
	return s.Frames[len(s.Frames)-1].Function
}

// NewError конструктор ошибки от уровня и строки
func NewError(level logrus.Level, err string) *Error {
	var event *sentry.Event

	switch level {
	case logrus.WarnLevel:
		event = sentry.NewEvent()
		event.Message = err
		event.Timestamp = time.Now()
		event.Level = sentry.LevelInfo

	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		event = sentry.NewEvent()

		trace := popStacktrace(sentry.NewStacktrace())

		event.Type = "error in function " + getCallerName(trace)
		event.Message = err
		event.Level = sentry.LevelError
		event.Timestamp = time.Now()

		event.Exception = append(event.Exception, sentry.Exception{
			Type:       err,
			Value:      err,
			Stacktrace: trace,
		})

	default:
		event = nil
	}

	return &Error{level: level, err: errors.New(err), sentryEvent: event}
}

/*
InputToSentryData на вход получает имена переменных(в виде строки через запятую), которые передаются в функцию и их значения.
На выход получаем мапу, которая помещается в хлебные крошки для sentry

Example

	func functionName(req, meta string) *models.Error {
		const source = "functionName"

		// вызов функции с параметрами
		err := otherFunction("a", "b", "c")
		if err != nil {
			// если вернулась ошибка, то оборачиваем и передаем дальше
			return err.WrapWithSentry(source, models.SentryCategoryFunc, models.InputToSentryData("req,meta", req,meta))
		}

		return nil
	}
*/
func InputToSentryData(names string, values ...any) map[string]any {
	if values == nil {
		return nil
	}

	names = strings.ReplaceAll(names, " ", "")
	names = strings.ReplaceAll(names, "\t", "")
	names = strings.ReplaceAll(names, "\n", "")
	ns := strings.Split(names, ",")

	res := make(map[string]any)
	for idx, name := range ns {
		if idx <= len(values)-1 {
			res[name] = values[idx]
		}
	}

	return res
}

type SentryCategory uint8

const (
	SentryCategoryFunc = iota
	SentryCategoryHandler
	SentryCategoryHttp
	SentryCategoryDB
	SentryCategoryCache
)

var sentryCategories = []struct{ Type, Category string }{
	{"default", "function"},
	{"http", "handler"},
	{"http", "http"},
	{"transaction", "db"},
	{"transaction", "cache"},
}

/*
WrapWithSentry оборачивает ошибку в переданное имя функции. При этом добавляя хлебные крошки для sentry
Для хлебных крошек устанавливается тип и категорию, соответствующие переданной SentryCategory (добавляются только для ошибок
критичнее Warn)

В WrapWithSentry можно передавать входные параметры (map[string]any) с которыми была запущена функция, чтобы
они добавились в хлебные крошки (для удобства использовать InputToSentryData)
*/
func (e *Error) WrapWithSentry(callerName string, sentryCategory SentryCategory, input map[string]any) *Error {
	e.err = errors.New("in " + callerName + ": " + e.err.Error())

	if e.level <= logrus.WarnLevel {
		var sentryLvl = sentry.LevelInfo
		if e.level <= logrus.ErrorLevel {
			sentryLvl = sentry.LevelError
		}
		e.sentryEvent.Breadcrumbs = append(e.sentryEvent.Breadcrumbs, &sentry.Breadcrumb{
			Type:      sentryCategories[sentryCategory].Type,
			Category:  sentryCategories[sentryCategory].Category,
			Message:   callerName + " with input: ",
			Data:      input,
			Level:     sentryLvl,
			Timestamp: time.Now(),
		})
	}

	return e
}

func (e *Error) GetLevel() logrus.Level {
	if e == nil {
		return logrus.TraceLevel
	}
	return e.level
}

func (e *Error) GetErr() error {
	if e == nil {
		return nil
	}
	return e.err
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return e.err.Error()
}

func (e *Error) Is(err error) bool {
	if e == nil {
		return false
	}
	return strings.Contains(e.err.Error(), err.Error())
}
