package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-fiber-example/m/v2/configs"
	"go-fiber-example/m/v2/i18n"
	"go-fiber-example/m/v2/middlewares"
	"go-fiber-example/m/v2/routes"
	"log"
	"os"
)

func main() {
	configs.LoadEnvVariables()
	configs.ConnectToDB()
	i18n.Load()
	app := fiber.New()
	middlewares.LoadLoggerMiddleware(app)
	middlewares.LoadCorsMiddleware(app)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.RegisterApiRoutes(app)
	err := app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatal("Failed to init app. \n", err)
		return
	}
}
