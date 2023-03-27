package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-example/m/v2/configs"
	"go-fiber-example/m/v2/models"
)

type UserController struct {
}

func NewUserController() UserController {
	return UserController{}
}

func (controller *UserController) Index(ctx *fiber.Ctx) error {
	var users []models.User
	configs.DBConn.Preload("Products").Find(&users)
	return ctx.Status(fiber.StatusOK).JSON(users)
}
