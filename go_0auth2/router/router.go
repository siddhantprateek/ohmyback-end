package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	handler "github.com/siddhantprateek/ohmyback-end/go_0auth2/handler"
)

func Routes(app *fiber.App) {

	api := app.Group("/api", logger.New())

	api.Get("/", handler.Auth)
	api.Get("/auth/google/callback", handler.Callback)
}
