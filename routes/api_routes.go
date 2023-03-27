package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-example/m/v2/controllers"
)

func RegisterApiRoutes(route fiber.Router) {

	demoController := controllers.NewDemoController()
	userController := controllers.NewUserController()
	productController := controllers.NewProductController()

	api := route.Group("/api")

	api.Get("/demo", demoController.Index)
	api.Get("/demo/i18n/hello", demoController.I18nHelloWorld)
	api.Get("/demo/i18n/name/:name", demoController.I18nHelloName)

	api.Get("/users", userController.Index)

	api.Get("/products", productController.Index)
	api.Get("/products/:id", productController.Show)
	api.Post("/products", productController.Create)
	api.Put("/products/:id", productController.Update)
	api.Delete("/products/:id", productController.Delete)
}
