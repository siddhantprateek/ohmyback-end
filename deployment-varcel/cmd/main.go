package handler

import (
	"net/http"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
)

// Handler is the main entry point of the application.
func Handler(w http.ResponseWriter, r *http.Request) {

	r.RequestURI = r.URL.String()

	handler().ServeHTTP(w, r)

}

func handler() http.HandlerFunc {

	app := fiber.New()

	app.Get("/v1", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"version": "v1",
		})
	})

	app.Get("/v2", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"version": "v2",
		})
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"uri":  ctx.Request().URI().String(),
			"path": ctx.Path(),
		})
	})

	return adaptor.FiberApp(app)
}
