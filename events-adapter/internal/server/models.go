package server

import (
	"events-adapter/internal/models"
	"events-adapter/pkg/errs"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	cfg   Config
	logs  errs.LogChan
	a     *fiber.App
	cache Setter
}

type Config struct {
	models.BaseConfig
	// special fields for server ...
}

type Setter interface {
	// AddValue добавляет элемент к указанной таблице |ключ(string) : значение([]string)|
	AddValue(dbID int, key string, val string) *errs.Error
}
