package main

import (
	"context"

	"github.com/XiroXD/fiber-ent-crud-app/provider/db"
	"github.com/gofiber/fiber/v2"
)

func init() {
	db.Connect()

	ctx := context.Background()

	cat, _ := db.GetClient().Pet.Create().SetName("Garfield").SetBreed("Maine Coon").SetAge(5).Save(ctx)
	dog, _ := db.GetClient().Pet.Create().SetName("Odie").SetBreed("Beagle").SetAge(3).Save(ctx)

	db.GetClient().User.Create().SetName("Jon Arbuckle").SetEmail("JonArbuckle@example.com").SetPassword("password").AddPets(cat, dog).Save(ctx)
}
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		u, _ := db.GetClient().User.Query().WithPets().Only(c.Context())

		return c.JSON(u)
	})

	app.Listen(":3000")
}
