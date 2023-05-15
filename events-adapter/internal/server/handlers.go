package server

import (
	"events-adapter/internal/models"
	"events-adapter/internal/repo/cache"
	"events-adapter/internal/repo/logic"
	"events-adapter/pkg/errs"
	"events-adapter/pkg/tools/den"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (s *Server) ping(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (s *Server) event(c *fiber.Ctx) error {
	const source = "event"

	var req models.BaseEventRequest
	err := den.DecodeJson(&req, c.Body())
	if err != nil {
		s.logs <- err.WrapWithSentry(source, errs.SentryCategoryHandler, errs.InputToSentryData("req", req))

		return c.SendStatus(http.StatusBadRequest)
	}

	ev := logic.CreateRedisEvent(req)
	encoded, err := den.EncodeJson(ev)
	if err != nil {
		s.logs <- err.WrapWithSentry(source, errs.SentryCategoryHandler, errs.InputToSentryData("req", req))

		return c.SendStatus(http.StatusInternalServerError)
	}

	err = s.cache.AddValue(cache.UserEvents, req.UserID, encoded.String())
	if err != nil {
		s.logs <- err.WrapWithSentry(source, errs.SentryCategoryHandler, errs.InputToSentryData("req", req))

		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

func (s *Server) search(c *fiber.Ctx) error {
	const source = "search"

	var req models.SearchEventRequest
	err := den.DecodeJson(&req, c.Body())
	if err != nil {
		s.logs <- err.WrapWithSentry(source, errs.SentryCategoryHandler, errs.InputToSentryData("req", req))

		return c.SendStatus(http.StatusBadRequest)
	}

	searchEv := logic.CreateRedisSearch(req)
	encoded, err := den.EncodeJson(searchEv)
	if err != nil {
		s.logs <- err.WrapWithSentry(source, errs.SentryCategoryHandler, errs.InputToSentryData("req", req))

		return c.SendStatus(http.StatusInternalServerError)
	}

	err = s.cache.AddValue(cache.UserSearches, req.UserID, encoded.String())
	if err != nil {
		s.logs <- err.WrapWithSentry(source, errs.SentryCategoryHandler, errs.InputToSentryData("req", req))

		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}
