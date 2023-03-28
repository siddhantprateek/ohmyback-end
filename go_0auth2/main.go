package main

import (
	router "github.com/siddhantprateek/ohmyback-end/go_0auth2/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New()

	router.Routes(app)

	app.Listen("3030")

}
