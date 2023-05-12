package es

import (
	"bytes"
	"es-writer/pkg/errs"
	"github.com/sirupsen/logrus"
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

func (c *Client) InsertDocument(name string, req []byte) *errs.Error {
	const source = "es.Insert"

	res, err := c.c.Index(name, bytes.NewReader(req))
	if err != nil {
		return errs.NewError(logrus.ErrorLevel, err.Error()).Wrap(source)
	}
	defer res.Body.Close()

	return nil
}
