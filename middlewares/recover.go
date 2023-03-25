package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func LoadRecoverMiddleware(app *fiber.App) {
	app.Use(recover.New())
}
