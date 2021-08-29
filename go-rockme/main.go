package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/hello", Hello)
	app.Listen(":8500")
}

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
