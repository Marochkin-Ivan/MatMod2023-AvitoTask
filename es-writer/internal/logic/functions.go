package logic

import (
	"bytes"
	"es-writer/internal/models"
	"es-writer/internal/server"
	"es-writer/pkg/errs"
	"es-writer/pkg/tools/den"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gofiber/fiber/v2"
	"log"
	"strings"
)

// SplitIntoBatches Функция для разделения среза на пачки по указанному размеру
func SplitIntoBatches(slice []server.FileInfo, batchSize int) [][]server.FileInfo {
	var batches [][]server.FileInfo

	for batchSize < len(slice) {
		slice, batches = slice[batchSize:], append(batches, slice[0:batchSize:batchSize])
	}
	batches = append(batches, slice)

	return batches
}

func InsertDocuments(docs []server.FileInfo, indexName string, insert func(name string, req []byte, id string) *errs.Error) *errs.Error {
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

func InsertBatchOfDocuments(batches [][]server.FileInfo, indexName string, insert func(req esapi.BulkRequest) *errs.Error) *errs.Error {
	const source = "logic.InsertBatchOfDocuments"

	for idx, batch := range batches {
		body, err := CreateBulkRequestBody(batch)
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

// CreateBulkRequestBody Функция для создания тела пакетного запроса индексации
func CreateBulkRequestBody(vacancies []server.FileInfo) ([]byte, *errs.Error) {
	const source = "logic.CreateBulkRequestBody"

	var body []byte

	for _, vacancy := range vacancies {
		metaData := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": models.IndexName,
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

func GetQueryParams(c *fiber.Ctx) map[string]string {
	v := make(map[string]string)

	c.Request().URI().QueryArgs().VisitAll(func(key, value []byte) {
		v[string(key)] = string(value)
	})

	return v
}

func GetKeywords(vacancy server.FileInfo) []string {
	keywords := make(map[string]struct{})

	specialisations := strings.Split(vacancy.ProfessionalSphereName, ",")

	// Удаление начальных и конечных пробелов в каждой специализации
	for i := range specialisations {
		specialisations[i] = strings.TrimSpace(specialisations[i])
	}

	// Добавление специализаций в мапу (избежание повторений)
	for _, specialisation := range specialisations {
		if specialisation != "" {
			keywords[specialisation] = struct{}{}
		}
	}

	return mapToKeySlice(keywords)
}

func mapToKeySlice(m map[string]struct{}) []string {
	var res []string

	for key := range m {
		res = append(res, key)
	}

	return res
}
