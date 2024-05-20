package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/sahildhargave/ShortenUrl/database"
)

// Resolve URL ...

func ResolveResolveURL(c *fiber.Ctx) error{
	// get the Short from Url
	url := c.Params("url")

	// query the db to find the original URL, if the a match is found 
	// increment the direct the redirect counter and redirect to the original URL
	// else return an error message
	r := database.CreateClient(0)
	defer r.Close()
	value, err := r.Get(database.Ctx, url).Result()

	if err == redis.Nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Short Not Found On Database",
		})

	}else if err != nil{
       return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Internal Server Error",
	   })

	}

	// increment the counter
	rInr := database.CreateClient(1)
	defer rInr.Close()
	_ = rInr.Incr(database.Ctx, "counter")
	// redirect to oritinal  URL
	return c.Redirect(value, 301)
}