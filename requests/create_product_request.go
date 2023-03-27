package requests

import (
	"fmt"
	u "go-fiber-example/m/v2/utils"
)

type CreateProductRequest struct {
	Code  string
	Price string
}

func ValidateCreateProductRequest(request CreateProductRequest, locale string) u.ValidationErrors {

	ve := u.Check("Code", request.Code, []string{u.NotNull}, locale)
	ve = u.Check("Price", request.Price, []string{fmt.Sprintf("%s:%d", u.MinLength, 6)}, locale)

	if len(ve) > 0 {
		return ve
	}
	return nil
}
