package net

import (
	"compress/gzip"
	"es-writer/pkg/errs"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// SendRequest создает и отправляет запрос с переданными параметрами.
//
// Возвращает: тело ответа, ошибку
func SendRequest(method string, url string, body io.Reader, token string, contentType string, client http.Client) ([]byte, *errs.Error) {
	const source = "SendRequest"

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errs.NewError(logrus.ErrorLevel, err.Error()).
			WrapWithSentry(source, errs.SentryCategoryHttp, errs.InputToSentryData("method,url", method, url))
	}

	if token != "" {
		req.Header.Set("Authorization", token)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
		req.Header.Set("Accept", contentType)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errs.NewError(logrus.ErrorLevel, err.Error()).
			WrapWithSentry(source, errs.SentryCategoryHttp, errs.InputToSentryData("method,url", method, url))
	}
	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	res, err := io.ReadAll(reader)
	if err != nil {
		return nil, errs.NewError(logrus.ErrorLevel, err.Error()).
			WrapWithSentry(source, errs.SentryCategoryHttp, errs.InputToSentryData("method,url", method, url))
	}

	return res, nil
}

// SendRequestWithStatus создает и отправляет запрос с переданными параметрами.
//
// Возвращает: тело ответа, статус код ответа, ошибку
func SendRequestWithStatus(method string, url string, body io.Reader, token string, contentType string, client http.Client) ([]byte, int, *errs.Error) {
	const source = "SendRequestWithStatus"

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, 0, errs.NewError(logrus.ErrorLevel, err.Error()).
			WrapWithSentry(source, errs.SentryCategoryHttp, errs.InputToSentryData("method,url", method, url))
	}

	if token != "" {
		req.Header.Set("Authorization", token)
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
		req.Header.Set("Accept", contentType)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, errs.NewError(logrus.ErrorLevel, err.Error()).
			WrapWithSentry(source, errs.SentryCategoryHttp, errs.InputToSentryData("method,url", method, url))
	}
	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		defer reader.Close()
	default:
		reader = resp.Body
	}

	res, err := io.ReadAll(reader)
	if err != nil {
		return nil, 0, errs.NewError(logrus.ErrorLevel, err.Error()).
			WrapWithSentry(source, errs.SentryCategoryHttp, errs.InputToSentryData("method,url", method, url))
	}

	return res, resp.StatusCode, nil
}
