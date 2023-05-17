package cache

import (
	"context"
	"es-writer/pkg/errs"
	"github.com/sirupsen/logrus"
	"log"
	"strconv"
)

// AddValueInPipeline добавляет множество значений к массиву |ключ(string) : значение(string)|
func (c Connections) AddValueInPipeline(dbID int, KVMap map[string][]string) *errs.Error {
	const source = "AddValueInPipeline"

	if dbID < 0 || dbID >= tablesCount {
		return errs.NewError(logrus.ErrorLevel, "unknown DB").
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("dbID", dbID))
	}

	rdb := c[dbID]
	pipe := rdb.TxPipeline()
	defer log.Println(pipe.Close())

	for key, values := range KVMap {
		e := pipe.SAdd(context.Background(), key, values).Err()
		if e != nil {
			return errs.NewError(logrus.ErrorLevel, e.Error()).
				WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("key", key))
		}
	}

	_, err := pipe.Exec(context.Background())
	if err != nil {
		return errs.NewError(logrus.ErrorLevel, err.Error()).
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("KVMap", KVMap))
	}

	return nil
}

// SetValueInPipeline создает элемент в указанной таблице |ключ(string) : значение(string)|
func (c Connections) SetValueInPipeline(dbID int, KVMap map[string]string) *errs.Error {
	const source = "SetValue"

	if dbID < 0 || dbID >= len(c) {
		return errs.NewError(logrus.ErrorLevel, "unknown DB").
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("dbID", dbID))
	}

	rdb := c[dbID]
	pipe := rdb.TxPipeline()
	defer log.Println(pipe.Close())

	for key, value := range KVMap {
		pipe.Set(context.Background(), key, value, 0)
	}

	_, err := pipe.Exec(context.Background())
	if err != nil {
		return errs.NewError(logrus.ErrorLevel, err.Error()).
			WrapWithSentry(source, errs.SentryCategoryCache, errs.InputToSentryData("KVMap", KVMap))
	}

	return nil
}

func (c Connections) Ping() *errs.Error {
	for _, dbID := range c {
		err := dbID.Ping(context.Background()).Err()
		if err != nil {
			return errs.NewError(logrus.ErrorLevel, "cache ping error: "+err.Error()).
				WrapWithSentry("Ping", errs.SentryCategoryCache, nil)
		}
	}

	return nil
}

// FlushAll очищает все таблицы в этом клиенте редиса
func (c Connections) FlushAll() *errs.Error {
	const source = "FlushAll"

	for i, rdb := range c {
		status := rdb.FlushDB(context.Background())
		if status.Err() != nil {
			return errs.NewError(logrus.WarnLevel, "DB: "+strconv.Itoa(i)+" "+status.Err().Error()).
				WrapWithSentry(source, errs.SentryCategoryCache, nil)
		}
	}

	return nil
}
