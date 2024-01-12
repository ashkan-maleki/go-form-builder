package main

import (
	"fmt"
	"github.com/ashkan-maleki/go-form-builder/server/gateway/internal/services/forms"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	port := "3000"
	log.Println(fmt.Sprintf("The API Gateway service is running on port %v", port))
	log.Println("http://127.0.0.1:3000")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	forms.ApiSetup(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
