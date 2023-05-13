package server

import (
	"api/internal/models"
	"api/pkg/errs"
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
