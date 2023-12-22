package main

import (
	"github.com/XiroXD/fiber-ent-crud-app/provider/db"
	"github.com/gofiber/fiber/v2"
)

func init() {
	db.Connect()
}
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
