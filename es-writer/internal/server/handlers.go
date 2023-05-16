package server

import (
	"es-writer/internal/models"
	"es-writer/pkg/errs"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (s *Server) ping(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

type getListReq struct {
}

func (s *Server) getList(c *fiber.Ctx) error {
	const source = "server.getList"

	var req getListReq
	if err := c.QueryParser(&req); err != nil {
		s.logs <- errs.NewError(logrus.InfoLevel, err.Error()).Wrap(source)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// req.Validate()

	// list, err := s.app.GetList(req)
	// if err != nil {
	// 	s.logs <- errs.NewError(logrus.InfoLevel, err.Error()).Wrap(source)
	// 	return c.Status(http.StatusInternalServerError).SendString(err.Error())
	// }

	return c.Status(http.StatusOK).SendString(models.DefaultList)

	//return c.SendStatus(http.StatusOK)
}

type getDetailReq struct {
}

func (s *Server) getDetail(c *fiber.Ctx) error {
	const source = "server.getDetail"

	var req getDetailReq
	if err := c.QueryParser(&req); err != nil {
		s.logs <- errs.NewError(logrus.InfoLevel, err.Error()).Wrap(source)
		return c.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// req.Validate()

	// detail, err := s.app.GetDetail(req)
	// if err != nil {
	// 	s.logs <- errs.NewError(logrus.InfoLevel, err.Error()).Wrap(source)
	// 	return c.Status(http.StatusInternalServerError).SendString(err.Error())
	// }

	return c.Status(http.StatusOK).SendString(models.DefaultInfo)

	//return c.SendStatus(http.StatusOK)
}
