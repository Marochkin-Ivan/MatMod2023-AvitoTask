package cache

import (
	"context"
	"events-adapter/pkg/errs"
	"github.com/sirupsen/logrus"
)

// AddValue добавляет элемент к указанной таблице |ключ(string) : значение([]string)|
func (c Connections) AddValue(dbID int, key string, val string) *errs.Error {
	const source = "AddValue"

	if dbID < 0 || dbID >= tablesCount {
		return errs.NewError(logrus.ErrorLevel, "unknown DB").
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("dbID", dbID))
	}

	rdb := c[dbID]

	e := rdb.SAdd(context.Background(), key, val)
	if e != nil {
		return errs.NewError(logrus.ErrorLevel, "cache add error: "+e.Err().Error()).
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("dbID, key, val", dbID, key, val))
	}

	return nil
}
