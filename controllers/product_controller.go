package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-example/m/v2/configs"
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
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(products)
}

func (controller *ProductController) Show(ctx *fiber.Ctx) error {
	product, err := services.GetProductById(ctx.Params("id"))
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(product)
}

//func (controller *ProductController) Create(ctx *fiber.Ctx) error {
//	request := new(requests.CreateProductRequest)
//	if err := ctx.BodyParser(request); err != nil {
//		return err
//	}
//	if err := requests.ValidateCreateProductRequest(*request, ctx.Get("Accept-Language")); err != nil {
//		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err)
//	}
//	product, err := services.CreateProduct()
//	if err != nil {
//		return err
//	}
//	return ctx.Status(fiber.StatusCreated).JSON(product)
//}

func (controller *ProductController) Create(ctx *fiber.Ctx) error {
	request := new(requests.CreateProductRequest)
	if err := ctx.BodyParser(request); err != nil {
		return err
	}
	if err := requests.ValidateCreateProductRequest(*request, ctx.Get("Accept-Language")); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}
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
	err := services.DeleteProduct(ctx.Params("id"))
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
