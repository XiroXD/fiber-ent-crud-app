package main

import (
	"github.com/XiroXD/fiber-ent-crud-app/provider/db"
	"github.com/gofiber/fiber/v2"
)

func init() {
	db.Connect()
}
func main() {
	server()

	// Close database connection before exit
	defer db.Disconnect()
}

func server() {
	app := fiber.New()

	app.Listen(":3000")
}
