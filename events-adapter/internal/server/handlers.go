package server

import (
	"events-adapter/internal/repo/cache"
	"events-adapter/internal/repo/logic"
	"events-adapter/pkg/errs"
	"events-adapter/pkg/tools/den"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (s *Server) ping(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (s *Server) event(c *fiber.Ctx) error {
	const source = "event"

	userID := c.Cookies("user_id")
	s.logs <- errs.NewError(logrus.DebugLevel, userID).Wrap(source)

	vacancyID := c.Query("vacancy_id")
	s.logs <- errs.NewError(logrus.DebugLevel, vacancyID).Wrap(source)

	eventType := c.Query("type")
	s.logs <- errs.NewError(logrus.DebugLevel, eventType).Wrap(source)

	ev := logic.CreateRedisEvent(vacancyID, eventType)
	encoded, err := den.EncodeJson(ev)
	if err != nil {
		s.logs <- err.WrapWithSentry(
			source,
			errs.SentryCategoryHandler,
			errs.InputToSentryData("vacancyID, type", vacancyID, eventType),
		)

		return c.SendStatus(http.StatusInternalServerError)
	}

	err = s.cache.AddValue(cache.UserEvents, userID, encoded.String())
	if err != nil {
		s.logs <- err.WrapWithSentry(
			source,
			errs.SentryCategoryHandler,
			errs.InputToSentryData("vacancyID, type", vacancyID, eventType),
		)

		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}
