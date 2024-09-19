package api

import (
	"encoding/json"
	"log"

	"github.com/pkrayzel/health-sync-api/domain"
	"github.com/pkrayzel/health-sync-api/metrics"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// UploadHandler handles the upload and stores the payload in Redis
func UploadHandler(c *fiber.Ctx) error {
	rawBody := c.Body()
	err := rdb.Set(ctx, "latest_payload", rawBody, 0).Err()
	if err != nil {
		log.Printf("Error storing payload in Redis: %v", err)
		return c.Status(500).SendString("Failed to store payload in Redis")
	}

	return c.SendString("Payload was stored")
}

// GetPayloadHandler retrieves the stored payload from Redis
func GetPayloadHandler(c *fiber.Ctx) error {
	payload, err := rdb.Get(ctx, "latest_payload").Result()
	if err == redis.Nil {
		return c.Status(404).SendString("No payload found")
	} else if err != nil {
		log.Printf("Error retrieving payload from Redis: %v", err)
		return c.Status(500).SendString("Error retrieving payload from Redis")
	}

	return c.SendString(payload)
}

// GetAverageCaloriesHandler retrieves the payload from Redis, parses it, and calculates the average calories.
func GetAverageCaloriesHandler(c *fiber.Ctx) error {
	// Get the latest payload from Redis
	payload, err := rdb.Get(ctx, "latest_payload").Result()
	if err == redis.Nil {
		return c.Status(404).SendString("No payload found")
	} else if err != nil {
		log.Printf("Error retrieving payload from Redis: %v", err)
		return c.Status(500).SendString("Error retrieving payload from Redis")
	}

	// Unmarshal the payload into a map
	var parsedPayload map[string]interface{}
	if err := json.Unmarshal([]byte(payload), &parsedPayload); err != nil {
		log.Printf("Error parsing payload from Redis: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Failed to parse payload from Redis")
	}

	// Parse the payload into domain models
	metricsData, err := metrics.ParsePayload(parsedPayload)
	if err != nil {
		log.Printf("Error parsing metrics: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid data")
	}

	// Call domain service for processing/aggregation
	avgCalories := domain.CalculateAverageCalories(metricsData)
	if avgCalories == 0 {
		log.Printf("No calorie data available to calculate average")
		return c.Status(200).SendString("No calorie data available")
	}

	log.Printf("Average daily calories: %f", avgCalories)

	// Return the calculated average in JSON format
	return c.JSON(fiber.Map{
		"averageCalories": avgCalories,
	})
}
