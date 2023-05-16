package cache

import (
	"api/pkg/errs"
	"context"
	"github.com/sirupsen/logrus"
)

// GetValue возвращает string
func (c Client) GetValue(dbID int, key string) (string, *errs.Error) {
	const source = "GetValue"

	if dbID < 0 || dbID >= tableCount {
		return "", errs.NewError(logrus.ErrorLevel, "unknown DB").
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("dbID", dbID))
	}

	rdb := c.cs[dbID]

	val := rdb.Get(context.Background(), key).Val()
	if val == "" {
		return "", errs.NewError(logrus.WarnLevel, "cache get error: key ["+key+"] isn't exist in "+DBName[dbID]+" DB").
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("dbId, key", dbID, key))
	}

	return val, nil
}

// GetArray возвращает []string
func (c Client) GetArray(dbID int, key string) ([]string, *errs.Error) {
	const source = "GetValue"

	if dbID < 0 || dbID >= tableCount {
		return nil, errs.NewError(logrus.ErrorLevel, "unknown DB").
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("dbID", dbID))
	}

	rdb := c.cs[dbID]

	val := rdb.SMembers(context.Background(), key).Val()
	if val == nil {
		return nil, errs.NewError(logrus.WarnLevel, "cache get error: key ["+key+"] isn't exist in "+DBName[dbID]+" DB").
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("dbId, key", dbID, key))
	}

	return val, nil
}
