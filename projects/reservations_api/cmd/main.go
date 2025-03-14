package main

import (
	"log"
	"reservations_api/api/availability"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	router := app.Group("/availability")

	availability.RegisterRoutes(router)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
