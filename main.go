package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("HELLO WORLD") // console.log
	app := fiber.New()         // new server
	todos := []Todo{}          // INTIALIZE data storage (array) of Todo nodes

	app.Get("/api/todos", func(c *fiber.Ctx) error { // GET todo
		return c.Status(200).JSON(todos)
	})

	app.Post("api/todos", func(c *fiber.Ctx) error { // CREATE todo
		todo := &Todo{}

		if error := c.BodyParser(todo); error != nil {
			return error
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		todo.ID = len(todos) + 1 // INCRREMENT TODO ID BY 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})

	// UPDATE/Patch
	app.Patch("api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(fiber.Map{"error": "Todo not found"})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "TODO NOT FOUND"})
	})

	// DELETE A TODO
	app.Delete("api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": "delete deleted"})
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "todo not fooound to delete"})

	})
	log.Fatal(app.Listen(":4000")) // error conosole
}
