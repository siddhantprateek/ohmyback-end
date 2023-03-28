package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Go K8s running perfectly")
	})

	log.Fatal(app.Listen(":4000"))
}
