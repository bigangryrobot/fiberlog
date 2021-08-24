package main

import (
	"github.com/bigangryrobot/fiberlog"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Default
	app.Use(fiberlog.New())

	app.Get("/ok", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Get("/warn", func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Get("/err", func(c *fiber.Ctx) error {
		return c.SendStatus(500)
	})

	app.Listen(":3000")
}
