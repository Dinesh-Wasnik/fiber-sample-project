// cmd/api/main.go
package main

import (
	"fiber-sample-project/app/config"
	"fiber-sample-project/app/routes"
	"fiber-sample-project/docs"

	swagger "github.com/gofiber/contrib/v3/swaggo"

	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {

	app := fiber.New(fiber.Config{
		ErrorHandler: config.ErrorHandler,
	})

	config.ErrorLog.Printf(
		"info | %v",
		"First message "+os.Getenv("APP_NAME")+"...",
	)

	//log the panic message
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c fiber.Ctx, e interface{}) {
			config.ErrorLog.Printf(
				"[PANIC] %s %s | %v",
				c.Method(),
				c.OriginalURL(),
				e,
			)
		},
	}))

	app.Use(logger.New(config.AccessLogger()))
	docs.SwaggerInfo.Host = os.Getenv("APP_URL")

	//route for swagger docs
	app.Get("/swagger/*", swagger.New())

	routes.Api(app)
	routes.Web(app)

	if err := app.Listen(":3000"); err != nil {
		log.Printf("❌ Server error: %v\n", err)
	}
}
