package controllers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/XiroXD/fiber-ent-crud-app/api/services"
	"github.com/XiroXD/fiber-ent-crud-app/schema"
)

type UsersController struct{}

func NewUsersController() *UsersController {
	return &UsersController{}
}

func (c *UsersController) Index(ctx *fiber.Ctx) error {
	return services.Hello(ctx)
}

func (c *UsersController) Create(ctx *fiber.Ctx) error {
	b := new(schema.CreateUser)

	if err := ctx.BodyParser(b); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	return services.Create(ctx, b.Name, b.Email, b.Password)
}
