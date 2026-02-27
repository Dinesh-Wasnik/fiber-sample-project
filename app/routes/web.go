package routes

import (
	"fiber-sample-project/app/config"
	"os"

	"github.com/gofiber/fiber/v3"
)

func Web(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("welcome to " + os.Getenv("APP_NAME"))
	})
	app.Get("/health", func(c fiber.Ctx) error {

		config.ErrorLog.Printf(
			"info | %v",
			"Starting "+os.Getenv("APP_NAME")+"...",
		)
		return c.SendString("I am okay")
	})

}
