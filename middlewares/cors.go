package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func LoadCorsMiddleware(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
}
