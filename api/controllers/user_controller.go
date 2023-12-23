package controllers

import (
	"github.com/XiroXD/fiber-ent-crud-app/api/services"
	"github.com/gofiber/fiber/v2"
)

type UsersController struct{}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (c *UsersController) Index(ctx *fiber.Ctx) error {
	return services.Hello(ctx)
}
