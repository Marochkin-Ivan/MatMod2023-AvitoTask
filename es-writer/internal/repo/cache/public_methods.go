package cache

import (
	"context"
	"es-writer/pkg/errs"
	"github.com/sirupsen/logrus"
)

// AddValue добавляет множество значений к массиву |ключ(string) : значение(string)|
func (c Connections) AddValue(dbID int, key string, val []string) *errs.Error {
	const source = "AddValue"

	if dbID < 0 || dbID >= tablesCount {
		return errs.NewError(logrus.ErrorLevel, "unknown DB").
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("dbID", dbID))
	}

	rdb := c[dbID]

	e := rdb.SAdd(context.Background(), key, val).Err()
	if e != nil {
		return errs.NewError(logrus.ErrorLevel, e.Error()).
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("key, val", key, val))
	}

	return nil
}

// SetValue создает элемент в указанной таблице |ключ(string) : значение(string)|
func (c Connections) SetValue(dbID int, key string, val string) *errs.Error {
	const source = "SetValue"

	if dbID < 0 || dbID >= len(c) {
		return errs.NewError(logrus.ErrorLevel, "unknown DB").
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("dbID", dbID))
	}

	rdb := c[dbID]

	err := rdb.Set(context.Background(), key, val, 0).Err()
	if err != nil {
		return errs.NewError(logrus.ErrorLevel, "cache set error: "+err.Error()).
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("dbID, key, val", dbID, key, val))
	}

	return nil
}
