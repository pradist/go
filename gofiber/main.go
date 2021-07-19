package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		//Format: "{\"pID\":\"${pid}\",\"reqID\":\"${locals:requestid}\", ${status} - ${method} ${path}}​\n​",
	}))

	app.Use(requestid.New(requestid.Config{
		Header: "req-id",
	}))

	//Meddleware
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("name", "pooh")
		fmt.Println("before")
		err := c.Next()
		fmt.Println("after")
		return err
	})

	//Get
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Locals("name")
		fmt.Println("hello")
		return c.SendString(fmt.Sprintf("GET: Hello %v\n", name))
	})

	//Post
	app.Post("/hello", func(c *fiber.Ctx) error {
		return c.SendString("POST: Hello")
	})

	//Parameter Optional
	app.Get("/hello/:name/:surname?", func(c *fiber.Ctx) error {
		name := c.Params("name")
		surname := c.Params("surname")
		return c.SendString("name: " + name + " surname: " + surname)
	})

	//Parameter Int
	app.Get("/hello/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.ErrBadRequest
		}
		return c.SendString(fmt.Sprintf("ID: %v", id))
	})

	//Query
	app.Get("/query", func(c *fiber.Ctx) error {
		name := c.Query("name")
		surname := c.Query("surname")
		return c.SendString("name: " + name + " surname: " + surname)
	})

	//Query struct
	app.Get("/query2", func(c *fiber.Ctx) error {
		person := Person{}
		c.QueryParser(&person)
		return c.JSON(person)
	})

	//Wildcards
	app.Get("/wildcards/*", func(c *fiber.Ctx) error {
		wildcards := c.Params("*")
		return c.SendString(wildcards)
	})

	//Static file
	app.Static("/", "./wwwroot", fiber.Static{
		Index:         "index.html",
		CacheDuration: time.Second * 10,
	})

	//New Error
	app.Get("/error", func(c *fiber.Ctx) error {
		fmt.Println("error")
		return fiber.NewError(fiber.StatusNotFound, "content not found")
	})

	app.Listen(":8000")
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
