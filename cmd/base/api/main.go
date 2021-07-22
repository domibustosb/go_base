package main

import (
	"log"
	"os"

	"github.com/domibustosb/go_base/internal/api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var port string
	app := fiber.New(fiber.Config{
		ReadBufferSize: 9000,
	})

	routes.SetupRoutes(app)

	if port = os.Getenv("HTTP_PORT"); port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
