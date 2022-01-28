package routes

import (
	"github.com/ashwani-zomato/shorten-url-fiber-redis/database"
	"github.com/gofiber/fiber/v2"
	"github.com/go-redis/redis/v8"
)


func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()

	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Short URL not found in the database!"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot connect to database!"})
	}

	rInr := database.CreateClient(1)

	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301)
}