package server

import (
	"es-writer/internal/models"
	"es-writer/pkg/errs"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	cfg  Config
	logs errs.LogChan
	a    *fiber.App
}

type Config struct {
	models.BaseConfig
	// special fields for server ...
}
