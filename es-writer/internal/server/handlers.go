package server

import (
	"encoding/json"
	"es-writer/internal/logic"
	"es-writer/internal/models"
	"es-writer/internal/repo/cache"
	"es-writer/pkg/errs"
	"es-writer/pkg/tools/den"
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

	err = s.es.CreateIndex(models.IndexName, body.Bytes())
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

	var vs []models.FileInfo
	e = json.Unmarshal(bs, &vs)
	if e != nil {
		s.logs <- errs.NewError(logrus.InfoLevel, e.Error()).Wrap(source)
		return c.SendStatus(http.StatusInternalServerError)
	}
	bs = nil
	vs = vs[6469:]
	// getKeys(vs)
	log.Println("len vs", len(vs))

	err := logic.InsertDocuments(vs, "vacancy", s.es.InsertDocument)
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

	var vs []models.FileInfo
	e = json.Unmarshal(bs, &vs)
	if e != nil {
		s.logs <- errs.NewError(logrus.InfoLevel, e.Error()).Wrap(source)
		return c.SendStatus(http.StatusInternalServerError)
	}
	bs = nil
	//vs = vs[:2000]

	// Создание пачек по 1000 записей
	batches := logic.SplitIntoBatches(vs, 1000)

	err := logic.InsertBatchOfDocuments(batches, models.IndexName, s.es.InsertBatchOfDocuments)
	if err != nil {
		s.logs <- err.Wrap(source)
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
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

func (s *Server) fillRedis(c *fiber.Ctx) error {
	const source = "server.fillRedis"

	err := s.cache.FlushAll()
	if err != nil {
		s.logs <- err.Wrap(source)
		return c.SendStatus(http.StatusInternalServerError)
	}

	bs, e := os.ReadFile(filename)
	if e != nil {
		s.logs <- errs.NewError(logrus.InfoLevel, e.Error()).Wrap(source)
		return c.SendStatus(http.StatusInternalServerError)
	}

	var vs []models.FileInfo
	e = json.Unmarshal(bs, &vs)
	if e != nil {
		s.logs <- errs.NewError(logrus.InfoLevel, e.Error()).Wrap(source)
		return c.SendStatus(http.StatusInternalServerError)
	}

	batches := logic.SplitIntoBatches(vs, 100)

	for idx, batch := range batches {
		infoKVMap := make(map[string]string)
		keywordsKVMap := make(map[string][]string)

		for _, vacancy := range batch {
			info, err := den.EncodeJson(vacancy)
			if err != nil {
				s.logs <- err.Wrap(source)
				return c.SendStatus(http.StatusInternalServerError)
			}

			infoKVMap[vacancy.ID] = info.String()

			keywords := logic.GetKeywords(vacancy)

			keywordsKVMap[vacancy.ID] = keywords
		}

		err = s.cache.SetValueInPipeline(cache.Info, infoKVMap)
		if err != nil {
			s.logs <- err.Wrap(source)
			return c.SendStatus(http.StatusInternalServerError)
		}

		err = s.cache.AddValueInPipeline(cache.Keywords, keywordsKVMap)
		if err != nil {
			s.logs <- err.Wrap(source)
			return c.SendStatus(http.StatusInternalServerError)
		}

		log.Println("batch", idx, "added")
	}

	return c.SendStatus(http.StatusOK)
}
