package translations

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-example/m/v2/utils"
)

func TranslateHelloWorld(ctx *fiber.Ctx) string {
	return utils.Localize(ctx, "helloWorld", nil)
}

func TranslateHelloName(ctx *fiber.Ctx) string {
	templateData := map[string]string{"Name": ctx.Params("name")}
	return utils.Localize(ctx, "helloName", templateData)
}
