package es

import (
	"api/pkg/errs"
	"bytes"
	"context"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/sirupsen/logrus"
	"io"
)

func (c *Client) Search(queryByte []byte) ([]byte, *errs.Error) {
	const source = "es.CreateIndex"

	searchRequest := esapi.SearchRequest{
		Index: []string{"vacancy"},
		Body:  bytes.NewReader(queryByte),
	}

	// Выполнение запроса на поиск
	res, e := searchRequest.Do(context.Background(), c.c)
	if e != nil {
		return nil, errs.NewError(logrus.ErrorLevel, e.Error()).Wrap(source)
	}
	defer res.Body.Close()

	// Обработка результатов поиска
	if res.IsError() {
		return nil, errs.NewError(logrus.ErrorLevel, res.Status()).Wrap(source)
	}

	body, e := io.ReadAll(res.Body)
	if e != nil {
		return nil, errs.NewError(logrus.ErrorLevel, e.Error()).Wrap(source)
	}

	return body, nil
}
