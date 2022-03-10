package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// Respond with a string back for GET calls to the endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("the API is UP!")
		return err
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}

}
