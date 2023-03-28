package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-example/m/v2/models"
	"go-fiber-example/m/v2/requests"
	"go-fiber-example/m/v2/services"
)

type ProductController struct {
}

func NewProductController() ProductController {
	return ProductController{}
}

func (controller *ProductController) Index(ctx *fiber.Ctx) error {
	products, err := services.GetAllProducts()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(products)
}

func (controller *ProductController) Show(ctx *fiber.Ctx) error {
	product, err := services.GetProductById(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(product)
}

func (controller *ProductController) Create(ctx *fiber.Ctx) error {
	/* validation */
	request := new(requests.CreateProductRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	if err := requests.ValidateCreateProductRequest(*request, ctx.Get("Accept-Language")); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}
	/* validation end */
	product := new(models.Product)
	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	newProduct, err := services.CreateProduct(*product)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return ctx.Status(fiber.StatusCreated).JSON(newProduct)
}

func (controller *ProductController) Update(ctx *fiber.Ctx) error {
	product := new(models.Product)
	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	err := services.UpdateProduct(*product, ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (controller *ProductController) Delete(ctx *fiber.Ctx) error {
	err := services.DeleteProduct(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
