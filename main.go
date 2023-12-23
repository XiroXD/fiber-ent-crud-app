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

	// Close database connection before exit
	defer db.Disconnect()

	app.Listen(":3000")
}
