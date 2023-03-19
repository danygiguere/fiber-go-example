package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-example/m/v2/configs"
	"go-fiber-example/m/v2/models"
)

type ProductController struct {
}

func NewProductController() ProductController {
	return ProductController{}
}

func (controller *ProductController) Index(ctx *fiber.Ctx) error {
	var products []models.Product
	configs.DBConn.Find(&products)
	return ctx.Status(fiber.StatusOK).JSON(products)
}

func (controller *ProductController) Show(ctx *fiber.Ctx) error {
	var product models.Product
	configs.DBConn.First(&product, ctx.Params("id"))
	return ctx.Status(fiber.StatusOK).JSON(product)
}

func (controller *ProductController) Create(ctx *fiber.Ctx) error {
	product := new(models.Product)
	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	configs.DBConn.Create(&product)
	return ctx.Status(fiber.StatusCreated).JSON(product)
}

func (controller *ProductController) Update(ctx *fiber.Ctx) error {
	product := new(models.Product)
	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	updatedProduct := models.Product{Code: product.Code, Price: product.Price}
	configs.DBConn.Model(&product).Where("id = ?", ctx.Params("id")).Updates(updatedProduct)
	return ctx.Status(fiber.StatusNoContent).JSON("updated")
}

func (controller *ProductController) Delete(ctx *fiber.Ctx) error {
	configs.DBConn.Delete(&models.Product{}, ctx.Params("id"))
	return ctx.Status(fiber.StatusNoContent).JSON("deleted")
}
