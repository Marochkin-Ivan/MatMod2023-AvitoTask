package server

import (
	"bytes"
	"encoding/json"
	"es-writer/pkg/errs"
	"es-writer/pkg/tools/den"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
)

const filename = "rr_vacancies_normal_format.json"

func (s *Server) ping(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (s *Server) create(c *fiber.Ctx) error {
	const source = "server.create"

	body, err := den.EncodeJson(createIndexRequest)
	if err != nil {
		return err.Wrap(source)
	}

	err = s.es.CreateIndex(indexName, body.Bytes())
	if err != nil {
		return err.Wrap(source)
	}

	return c.SendStatus(http.StatusOK)
}

func (s *Server) insert(c *fiber.Ctx) error {
	const source = "server.insert"

	bs, e := os.ReadFile(filename)
	if e != nil {
		s.logs <- errs.NewError(logrus.InfoLevel, e.Error()).Wrap(source)
		return c.SendStatus(http.StatusInternalServerError)
	}

	var vs []FileInfo
	e = json.Unmarshal(bs, &vs)
	if e != nil {
		s.logs <- errs.NewError(logrus.InfoLevel, e.Error()).Wrap(source)
		return c.SendStatus(http.StatusInternalServerError)
	}
	bs = nil
	vs = vs[6469:]
	// getKeys(vs)
	log.Println("len vs", len(vs))

	err := InsertDocuments(vs, "vacancy", s.es.InsertDocument)
	if err != nil {
		s.logs <- err.Wrap(source)
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

func (s *Server) insertBatch(c *fiber.Ctx) error {
	const source = "server.insertBatch"

	bs, e := os.ReadFile(filename)
	if e != nil {
		s.logs <- errs.NewError(logrus.InfoLevel, e.Error()).Wrap(source)
		return c.SendStatus(http.StatusInternalServerError)
	}

	var vs []FileInfo
	e = json.Unmarshal(bs, &vs)
	if e != nil {
		s.logs <- errs.NewError(logrus.InfoLevel, e.Error()).Wrap(source)
		return c.SendStatus(http.StatusInternalServerError)
	}
	bs = nil
	//vs = vs[:2000]

	// Создание пачек по 1000 записей
	batches := splitIntoBatches(vs, 1000)

	err := InsertBatchOfDocuments(batches, indexName, s.es.InsertBatchOfDocuments)
	if err != nil {
		s.logs <- err.Wrap(source)
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

// Функция для разделения среза на пачки по указанному размеру
func splitIntoBatches(slice []FileInfo, batchSize int) [][]FileInfo {
	var batches [][]FileInfo

	for batchSize < len(slice) {
		slice, batches = slice[batchSize:], append(batches, slice[0:batchSize:batchSize])
	}
	batches = append(batches, slice)

	return batches
}

func InsertDocuments(docs []FileInfo, indexName string, insert func(name string, req []byte, id string) *errs.Error) *errs.Error {
	const source = "logic.InsertDocuments"

	for idx, doc := range docs {
		tmp := doc.ToESDoc()
		req, err := den.EncodeJson(tmp)
		if err != nil {
			return err.Wrap(source)
		}

		if err := insert(indexName, req.Bytes(), doc.ID); err != nil {
			return err.Wrap(source)
		}
		log.Println("inserted", idx)
	}

	return nil
}

func InsertBatchOfDocuments(batches [][]FileInfo, indexName string, insert func(req esapi.BulkRequest) *errs.Error) *errs.Error {
	const source = "logic.InsertBatchOfDocuments"

	for idx, batch := range batches {
		body, err := createBulkRequestBody(batch)
		if err != nil {
			return err.Wrap(source)
		}

		// Создание пакетного запроса индексации
		bulkRequest := esapi.BulkRequest{
			Index: indexName,
			Body:  bytes.NewReader(body),
		}

		if err := insert(bulkRequest); err != nil {
			return err.Wrap(source)
		}
		log.Println("inserted batch", idx)
	}

	return nil
}

const indexName = "vacancy"

// Функция для создания тела пакетного запроса индексации
func createBulkRequestBody(vacancies []FileInfo) ([]byte, *errs.Error) {
	const source = "logic.createBulkRequestBody"

	var body []byte

	for _, vacancy := range vacancies {
		metaData := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": indexName,
			},
		}

		metaDataBytes, err := den.EncodeJson(metaData)
		if err != nil {
			return nil, err.Wrap(source)
		}

		tmp := vacancy.ToESDoc()
		vacancyBytes, err := den.EncodeJson(tmp)
		if err != nil {
			return nil, err.Wrap(source)
		}

		body = append(body, metaDataBytes.Bytes()...)
		body = append(body, []byte("\n")...)
		body = append(body, vacancyBytes.Bytes()...)
		body = append(body, []byte("\n")...)
	}

	return body, nil
}

func getQueryParams(c *fiber.Ctx) map[string]string {
	v := make(map[string]string)

	c.Request().URI().QueryArgs().VisitAll(func(key, value []byte) {
		v[string(key)] = string(value)
	})

	return v
}

func (s *Server) test(c *fiber.Ctx) error {
	v := make(map[string]string)

	c.Request().URI().QueryArgs().VisitAll(func(key, value []byte) {
		v[string(key)] = string(value)
	})

	for k, v := range v {
		log.Println(k, v)
	}
	return c.SendStatus(http.StatusOK)
}
