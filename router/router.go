package router

import (
	"github.com/XiroXD/fiber-ent-crud-app/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	userController := controllers.NewUsersController()

	user := app.Group("/user")

	user.Get("/", userController.Index)
}
