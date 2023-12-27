package services

import (
	"github.com/gofiber/fiber/v2"

	"github.com/XiroXD/fiber-ent-crud-app/hash"
	"github.com/XiroXD/fiber-ent-crud-app/provider/db"
)

func Hello(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "Hello, World!"})
}

func Create(ctx *fiber.Ctx, name string, email string, password string) error {
	c := ctx.Context()

	hash := hash.CreateHash(password)

	user, err := db.GetClient().User.Create().
		SetName(name).
		SetEmail(email).
		SetPassword(string(hash)).
		Save(c)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.JSON(user)
}
