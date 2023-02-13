package pkg

import "github.com/gofiber/fiber/v2"

func Source() {
	router := fiber.New()

	router.Get("", func(c *fiber.Ctx) error {

		return nil
	})
	router.Listen(":8080")
}
