package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Print("hello world!")
	app := fiber.New();
	app.Get("/abc",func(c*fiber.Ctx)error{
		return c.SendString("Hello Ajay!\n hello")
	})
		app.Listen(":3000")
	}


