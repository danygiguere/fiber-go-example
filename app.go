package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-fiber-example/m/v2/configs"
	"go-fiber-example/m/v2/middlewares"
	"go-fiber-example/m/v2/routes"
	"go-fiber-example/m/v2/utils"
	"log"
	"os"
)

func main() {
	configs.LoadEnvVariables()
	configs.ConnectToDB()
	utils.Load()
	app := fiber.New()
	middlewares.LoadRecoverMiddleware(app)
	middlewares.LoadLoggerMiddleware(app)
	middlewares.LoadCorsMiddleware(app)

	routes.RegisterApiRoutes(app)
	err := app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatal("Failed to init app. \n", err)
		return
	}
}
