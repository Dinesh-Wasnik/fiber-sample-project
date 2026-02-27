package routes

import (
	"github.com/gofiber/fiber/v3"
)

func Api(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/health", func(c fiber.Ctx) error {
		return (c).JSON(fiber.Map{
			"status": "Api is ok",
		})
	})

}
