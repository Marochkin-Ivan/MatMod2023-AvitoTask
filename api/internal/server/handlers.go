package server

import (
	"api/pkg/errs"
	"api/pkg/tools/den"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (s *Server) ping(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (s *Server) search(c *fiber.Ctx) error {
	const source = "search"

	queryParams := getQueryParams(c)

	searchReq := createSearchRequest(queryParams["q"])

	for param, value := range queryParams {
		if filterFunc, exist := filtersTypeMap[param]; exist {
			searchReq.withFilter(filterFunc, param, value)
		}
	}

	b, _ := json.Marshal(searchReq)
	s.logs <- errs.NewError(logrus.DebugLevel, string(b)).Wrap(source)

	encodedSearchReq, err := den.EncodeJson(searchReq)
	if err != nil {
		s.logs <- errs.NewError(logrus.ErrorLevel, err.Error()).Wrap(source)

		return c.SendStatus(http.StatusInternalServerError)
	}

	resBytes, err := s.es.Search(encodedSearchReq.Bytes())
	if err != nil {
		s.logs <- errs.NewError(logrus.ErrorLevel, err.Error()).Wrap(source)

		return c.SendStatus(http.StatusInternalServerError)
	}

	var res ElasticResponse
	err = den.DecodeJson(&res, resBytes)
	if err != nil {
		s.logs <- errs.NewError(logrus.ErrorLevel, err.Error()).Wrap(source)

		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(res.ToResponse())
}
