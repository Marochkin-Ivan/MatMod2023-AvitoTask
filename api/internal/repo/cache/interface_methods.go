package cache

import (
	"api/pkg/errs"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

func (c *Client) Get(table int, key string) (string, *errs.Error) {
	const source = "cache.Get"

	if table >= tableCount || table < 0 {
		return "", errs.NewError(logrus.ErrorLevel, "unexpected table id").
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("table", table))
	}

	rdb := c.cs[table]

	// get the value and update last used timestamp
	val, err := rdb.Do(rdb.Context(), "GET", key).Result()
	if err != nil {
		return "", errs.NewError(logrus.ErrorLevel, err.Error()).
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("table, key", table, key))
	}

	return fmt.Sprintf("%v", val), nil
}

func (c *Client) Set(table int, key, val string) *errs.Error {
	const source = "cache.Get"

	if table >= tableCount || table < 0 {
		return errs.NewError(logrus.ErrorLevel, "unexpected table id").
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("table", table))
	}

	rdb := c.cs[table]

	err := rdb.Set(context.Background(), key, val, 0).Err()
	if err != nil {
		return errs.NewError(logrus.ErrorLevel, err.Error()).
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("table, key, val", table, key, val))
	}

	return nil
}

func (c *Client) Delete(table int, key string) *errs.Error {
	const source = "cache.Delete"

	if table >= tableCount || table < 0 {
		return errs.NewError(logrus.ErrorLevel, "unexpected table id").
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("table", table))
	}

	rdb := c.cs[table]

	err := rdb.Del(context.Background(), key).Err()
	if err != nil {
		return errs.NewError(logrus.ErrorLevel, err.Error()).
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("table, key", table, key))
	}

	return nil
}

func (c *Client) GetKeys(table int) ([]string, *errs.Error) {
	const source = "cache.GetKeys"

	if table >= tableCount || table < 0 {
		return nil, errs.NewError(logrus.ErrorLevel, "unexpected table id").
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("table", table))
	}

	rdb := c.cs[table]

	keys, err := rdb.Keys(context.Background(), "*").Result()
	if err != nil {
		return nil, errs.NewError(logrus.ErrorLevel, err.Error()).
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("table", table))
	}

	return keys, nil
}
