package main

import (
	"github.com/Jerry2204/go-employee-rest-api/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	employeeGroup := app.Group("/employees")
	employeeGroup.Get("/", handlers.GetAllEmployees)
}
