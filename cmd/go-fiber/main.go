package main

import (
	"local/go-benchmarks/internal/data"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send(data.Get())
	})

	app.Listen(":8000")
}
