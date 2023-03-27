package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-example/m/v2/resources/translations"
)

type DemoController struct{}

func NewDemoController() DemoController {
	return DemoController{}
}

func (controller *DemoController) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    "Hello World",
	})
}

func (controller *DemoController) I18nHelloWorld(ctx *fiber.Ctx) error {
	return ctx.SendString(translations.TranslateHelloWorld(ctx))
}

func (controller *DemoController) I18nHelloName(ctx *fiber.Ctx) error {
	return ctx.SendString(translations.TranslateHelloName(ctx))
}
