package routes

import (
	"github.com/ayushdiitr/url-shortener/db"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func ResolveURL(c *fiber.Ctx) error {

	url := c.Params("url")

	r := db.CreateClient(0)
	defer r.Close()

	val, err := r.Get(db.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found in the db",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to db",
		})
	}

	rInr := db.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(db.Ctx, "counter")

	return c.Redirect(val, 301)
}
