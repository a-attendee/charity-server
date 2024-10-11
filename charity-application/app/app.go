package app

import (
	"log/slog"
	"strconv"

	"github.com/a-attendee/charity-server/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	slogfiber "github.com/samber/slog-fiber"
)

type AppInstance struct {
    Port int
    App *fiber.App
}

func (a *AppInstance) Listen() error {
    return a.App.Listen(":" + strconv.Itoa(a.Port))
}

type Options func(*AppInstance)

func WithPort(port int) Options {
    return func(a *AppInstance) {
        a.Port = port
    }
}

func WithHandler(handler fiber.Handler) Options {
    return func(a *AppInstance) {
        a.App.Use(handler)
    }
}

func WithConfig(config config.Config) Options {
    return func(a *AppInstance) {
        a.Port = config.HttpServer.Port
        a.App.Use(limiter.New(
            limiter.Config{
                Max: config.HttpServer.MaxRequests,
                Expiration: config.HttpServer.MaxRequestsExperation,
            }))
    }
}

func WithLogger(logger *slog.Logger) Options {
    return func(a *AppInstance) {
        a.App.Use(slogfiber.New(logger))
    }
}

func New(opts ...Options) *AppInstance {
    return &AppInstance{
        Port: 3000,
        App: fiber.New(),
    }
}
