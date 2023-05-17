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

	for _, vacancy := range vs {
		info, err := den.EncodeJson(vacancy)
		if err != nil {
			s.logs <- err.Wrap(source)
			return c.SendStatus(http.StatusInternalServerError)
		}

		err = s.cache.SetValue(cache.Info, vacancy.ID, info.String())
		if err != nil {
			s.logs <- err.Wrap(source)
			return c.SendStatus(http.StatusInternalServerError)
		}

		keywords := logic.GetKeywords(vacancy)

		err = s.cache.AddValue(cache.Keywords, vacancy.ID, keywords)
		if err != nil {
			s.logs <- err.Wrap(source)
			return c.SendStatus(http.StatusInternalServerError)
		}
	}

	return c.SendStatus(http.StatusOK)
}
