package requests

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	u "go-fiber-example/m/v2/utils"
)

type CreateProductRequest struct {
	Code  string
	Price string
}

func ValidateCreateProductRequest(ctx *fiber.Ctx) CreateProductRequest {
	request := new(CreateProductRequest)
	ctx.BodyParser(request)

	locale := ctx.Get("Accept-Language")

	ve := make(u.ValidationErrors)
	ve = u.Check("Code", request.Code, []string{u.NotNull}, locale, ve)
	ve = u.Check("Price", request.Price, []string{fmt.Sprintf("%s:%d", u.MinLength, 6)}, locale, ve)

	if len(ve) > 0 {
		panic(ctx.Status(fiber.StatusUnprocessableEntity).JSON(ve))
	}
	return *request
}
