package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New()

	// Middleware to check API key
	app.Use(func(c *fiber.Ctx) error {
		apiKey := c.Get("X-API-Key")
		if apiKey != os.Getenv("API_KEY") {
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}
		return c.Next()
	})

	app.Post("/upload", func(c *fiber.Ctx) error {
		body := c.Body() // Get the body
		// Use it here, for example, print/log it
		fmt.Println(string(body)) // Print the body for now (remove in production)

		return c.SendString("Payload received")
	})

	// Start the server
	app.Listen(":8888")
}
