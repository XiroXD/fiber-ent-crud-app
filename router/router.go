package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/XiroXD/fiber-ent-crud-app/api/controllers"
)

func Setup(app *fiber.App) {
	userController := controllers.NewUsersController()

	user := app.Group("/user")

	user.Get("/", userController.Index)
	user.Post("/", userController.Create)
}
