package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go/v5"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	pusherClient := pusher.Client{
		AppID:   "1733954",
		Key:     "fedda67831cf87c65d4d",
		Secret:  "f94972b1f9fbe92f53d3",
		Cluster: "ap2",
		Secure:  true,
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		data := make(map[string]string)

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		pusherClient.Trigger("chat", "message", data)
		return c.JSON(data)
	})

	app.Listen(":8080")
}
