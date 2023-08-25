package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	model "github.com/siddhantprateek/ohmyback-end/go_fiber_redis/model"
)

func toJson(val []byte) model.User {
	user := model.User{}
	err := json.Unmarshal(val, &user)
	if err != nil {
		panic(err)
	}
	return user
}

func verifyCache(c *fiber.Ctx) error {
	id := c.Params("id")
	val, err := cache.Get(ctx, id).Bytes()
	if err != nil {
		return c.Next()
	}

	data := toJson(val)
	return c.JSON(fiber.Map{"Cached": data})
}

var cache = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

var ctx = context.Background()

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("It is working 👊")
	})

	app.Get("/:id", verifyCache, func(c *fiber.Ctx) error {
		id := c.Params("id")
		res, err := http.Get("https://jsonplaceholder.typicode.com/users/" + id)
		if err != nil {
			return err
		}

		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cacheErr := cache.Set(ctx, id, body, 10*time.Second).Err()
		if cacheErr != nil {
			return cacheErr
		}

		data := toJson(body)
		return c.JSON(fiber.Map{"Data": data})
	})

	app.Listen(":3000")
}
