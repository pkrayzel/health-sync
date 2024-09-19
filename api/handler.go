package api

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// UploadHandler handles the upload and stores the payload in Redis
func UploadHandler(c *fiber.Ctx) error {
	body := c.Body()

	// Store the body in Redis
	err := rdb.Set(ctx, "latest_payload", body, 0).Err()
	if err != nil {
		return c.Status(500).SendString("Failed to store payload in Redis")
	}

	return c.SendString("Payload stored successfully")
}

// GetPayloadHandler retrieves the stored payload from Redis
func GetPayloadHandler(c *fiber.Ctx) error {
	payload, err := rdb.Get(ctx, "latest_payload").Result()
	if err == redis.Nil {
		return c.Status(404).SendString("No payload found")
	} else if err != nil {
		return c.Status(500).SendString("Error retrieving payload from Redis")
	}

	return c.SendString(payload)
}
