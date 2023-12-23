package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"

	"github.com/XiroXD/fiber-ent-crud-app/provider/db"
)

func init() {
	db.Connect()
}
func main() {
	app := fiber.New()

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	// Create channel to signify a signal being sent
	c := make(chan os.Signal, 1)
	// When an interrupt or termination signal is sent, notify the channel
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c //lint:ignore S1005 Blocks main thread until an interrupt is received
	fmt.Println("\nGracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	db.Disconnect()
	fmt.Println("Fiber was successful shutdown.")
}
