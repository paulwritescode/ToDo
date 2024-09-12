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
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// POST Route for creating todos
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		// Parse body into a new Todo struct
		// todo := &Todo{}
		todo := new(Todo)
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
	// Update path(patch) route
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	// Delete path(delete) route
	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			// convert the todo.ID to string since it is an int from the interface
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": "Todo deleted"})
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	// Start the server
	log.Fatal(app.Listen(":4000"))
}
