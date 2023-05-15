package server

import (
	"api/internal/models"
	"api/pkg/errs"
	"api/pkg/tools/den"
	"github.com/caarlos0/env/v6"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

func GetConfig() (Config, *errs.Error) {
	const source = "getConfig"

	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, errs.NewError(logrus.ErrorLevel, err.Error()).Wrap(source)
	}

	return cfg, nil
}

func NewFiberApp(lg errs.Logger4Fiber) *fiber.App {
	a := fiber.New(fiber.Config{
		JSONEncoder: func(v any) ([]byte, error) {
			b, e := den.EncodeJson(v)
			if e != nil {
				return nil, e.GetErr()
			}
			return b.Bytes(), nil
		},
	})

	a.Use(logger.New(logger.Config{
		Format: "{\"status\": ${status}, \"duration\": \"${latency}\", \"method\": \"${method}\", \"path\": \"${path}\", \"resp\": \"${resBody}\"}\n",
		Output: lg,
	}))
	a.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	a.Use(recover.New())

	return a
}

func NewFiberAppWithJWT(lg errs.Logger4Fiber) *fiber.App {
	a := NewFiberApp(lg)
	a.Use(jwtware.New(jwtware.Config{
		SigningKey:    []byte(models.JWTSalt),
		SigningMethod: jwt.SigningMethodHS256.Name,
	}))

	return a
}

func SetupSentry(serverName, dsn string) *errs.Error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
		//Debug:            true,
		EnableTracing:    true,
		AttachStacktrace: true,
		ServerName:       serverName,
		MaxErrorDepth:    3, // test
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			event.User.Email = "" // Don't send user's email address

			return event
		},
		MaxBreadcrumbs: 15,
	})
	if err != nil {
		return errs.NewError(logrus.FatalLevel, err.Error()).Wrap("SetupSentry")
	}

	return nil
}

func getQueryParams(c *fiber.Ctx) map[string]string {
	v := make(map[string]string)

	c.Request().URI().QueryArgs().VisitAll(func(key, value []byte) {
		v[string(key)] = string(value)
	})

	return v
}
