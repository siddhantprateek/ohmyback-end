package handler

import (
	"github.com/gofiber/fiber/v2"
	auth "github.com/siddhantprateek/ohmyback-end/go_0auth2/auth"
)

// Auth fiber handler
func Auth(c *fiber.Ctx) error {
	path := auth.ConfigGoogle()
	url := path.AuthCodeURL("state")
	return c.Redirect(url)

}

// Callback to receive google's response
func Callback(c *fiber.Ctx) error {
	token, error := auth.ConfigGoogle().Exchange(c.Context(), c.FormValue("code"))
	if error != nil {
		panic(error)
	}
	email := auth.GetEmail(token.AccessToken)
	return c.Status(200).JSON(fiber.Map{"email": email, "login": true})

}
