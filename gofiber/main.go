package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func main() {
	app := fiber.New()
	var err error

	db, err = sqlx.Open("mysql", "root:P@ssw0rd@tcp(localhost:3306)/mysql1")
	if err != nil {
		panic(err)
	}

	app.Post("/signup", Signup)
	app.Post("/login", Login)
	app.Get("/hello", Hello)

	app.Listen(":8000")
}

func Signup(c *fiber.Ctx) error {
	request := SignupRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	if request.Username == "" || request.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	query := "insert user (username, password) values (?, ?)"
	result, err := db.Exec(query, request.Username, request.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	user := User{
		Id:       int(id),
		Username: request.Username,
		Password: request.Password,
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	return nil
}

func Hello(c *fiber.Ctx) error {
	return nil
}

type User struct {
	Id       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Fiber() {
	app := fiber.New(fiber.Config{
		Prefork: true,
	})

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
	// app.Get("/hello/:name/:surname?", func(c *fiber.Ctx) error {
	// 	name := c.Params("name")
	// 	surname := c.Params("surname")
	// 	return c.SendString("name: " + name + " surname: " + surname)
	// })

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

	v1 := app.Group("v1", func(c *fiber.Ctx) error {
		c.Set("version", "1")
		return c.Next()
	})
	v1.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("hello v1")
	})

	v2 := app.Group("v2", func(c *fiber.Ctx) error {
		c.Set("version", "2")
		return c.Next()
	})
	v2.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("hello v2")
	})

	//Mount
	userApp := fiber.New()
	userApp.Get("/login", func(c *fiber.Ctx) error {
		return c.SendString("Login")
	})

	app.Mount("/user", userApp)

	//Server
	app.Server().MaxConnsPerIP = 1
	app.Get("server", func(c *fiber.Ctx) error {
		time.Sleep(time.Second * 30)
		return c.SendString("Server")
	})

	app.Listen(":8000")
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
