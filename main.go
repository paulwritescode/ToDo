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
	fmt.Println("Hellow Moon, I love World")
	app := fiber.New()

	// Slice to hold todos globally
	todos := []Todo{}

	// GET Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hellow World, This is Moon",
		})
	})

	// POST Route for creating todos
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		// Parse body into a new Todo struct
		todo := &Todo{}
		// todo := new(Todo)
		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body meli"})

		}

		// Validate body field
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body is required"})
		}

		// Assign ID and add to todos slice
		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		// Return the newly created todo
		return c.Status(200).JSON(todos)
	})

	// Start the server
	log.Fatal(app.Listen(":4000"))
}
