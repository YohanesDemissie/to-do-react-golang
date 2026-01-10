package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("HELLO WORLD") // console.log
	app := fiber.New()         // new server

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello World"})
	})
	log.Fatal(app.Listen(":4000")) // error conosole
}
