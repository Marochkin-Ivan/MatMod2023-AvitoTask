package server

import (
	"es-writer/internal/models"
	"es-writer/internal/repo/es"
	"es-writer/pkg/errs"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	cfg  Config
	logs errs.LogChan
	a    *fiber.App
	es   *es.Client
}

type Config struct {
	models.BaseConfig
	// special fields for server ...
}
