package main

import (
	"github.com/Jerry2204/go-employee-rest-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":8080")
}
