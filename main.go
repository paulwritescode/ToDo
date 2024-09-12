package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hellow Moon i love marion")
	app := fiber.New()
	log.Fatal(app.Listen(":4000"))
}
