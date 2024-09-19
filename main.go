package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

	// Import from the "api" folder in the local project structure
	"github.com/pkrayzel/health-sync-api/api"
)

// Middleware to check API key
func apiKeyMiddleware(c *fiber.Ctx) error {
	apiKey := c.Get("X-API-Key")
	if apiKey != os.Getenv("API_KEY") {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}
	return c.Next()
}

func main() {
	// Initialize Redis connection
	api.InitRedis()

	// Create a new Fiber instance
	app := fiber.New()

	// Use middleware to check API key
	app.Use(apiKeyMiddleware)

	// Define API routes
	app.Post("/upload", api.UploadHandler)     // Upload route
	app.Get("/payload", api.GetPayloadHandler) // Retrieve payload route

	// Start the server on port 8888
	log.Fatal(app.Listen(":8888"))
}
