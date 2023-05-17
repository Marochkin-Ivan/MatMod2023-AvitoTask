package es

import (
	"bytes"
	"context"
	"es-writer/pkg/errs"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/sirupsen/logrus"
	"io"
	"log"
)

func (c *Client) CreateIndex(name string, req []byte) *errs.Error {
	const source = "es.CreateIndex"

	res, err := c.c.Indices.Create(name, c.c.Indices.Create.WithBody(bytes.NewReader(req)))
	if err != nil {
		return errs.NewError(logrus.ErrorLevel, err.Error()).Wrap(source)
	}
	defer res.Body.Close()

	if res.IsError() {
		return errs.NewError(logrus.ErrorLevel, res.String()).Wrap(source)
	}

	return nil
}

func (c *Client) InsertDocument(name string, req []byte, id string) *errs.Error {
	const source = "es.Insert"

	res, err := c.c.Index(name, bytes.NewReader(req), c.c.Index.WithDocumentID(id))
	if err != nil {
		return errs.NewError(logrus.ErrorLevel, err.Error()).Wrap(source)
	}
	defer res.Body.Close()
	b, _ := io.ReadAll(res.Body)
	log.Println(string(b))

	return nil
}

func (c *Client) InsertBatchOfDocuments(req esapi.BulkRequest) *errs.Error {
	const source = "es.InsertBatchOfDocuments"

	// Выполнение пакетного запроса индексации
	res, err := req.Do(context.Background(), c.c)
	if err != nil {
		return errs.NewError(logrus.ErrorLevel, err.Error()).Wrap(source)
	}
	defer res.Body.Close()

	// Проверка статуса ответа
	if res.IsError() {
		return errs.NewError(logrus.ErrorLevel, err.Error()).Wrap(source)
	}

	b, _ := io.ReadAll(res.Body)
	log.Println(string(b))

	return nil
}
