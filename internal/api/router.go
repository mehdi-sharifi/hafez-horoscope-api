package api

import (
	"github.com/gofiber/fiber/v2"
	"hafez-horoscope-api/internal/handlers"
)

func SetupRouter() *fiber.App {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"hello": "world"})
	})
	apiGroup := app.Group("/api")
	hafez := apiGroup.Group("/hafez")
	hafez.Get("/random", handlers.GetRandomPoem)
	return app
}
