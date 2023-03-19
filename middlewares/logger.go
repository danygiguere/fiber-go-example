package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func LoadLoggerMiddleware(app *fiber.App) {
	app.Use(logger.New())
}
