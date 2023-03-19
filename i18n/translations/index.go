package translations

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-example/m/v2/i18n"
)

func TranslateHelloWorld(ctx *fiber.Ctx) string {
	return i18n.Localize(ctx, "helloWorld", nil)
}

func TranslateHelloName(ctx *fiber.Ctx) string {
	templateData := map[string]string{"Name": ctx.Params("name")}
	return i18n.Localize(ctx, "helloName", templateData)
}
