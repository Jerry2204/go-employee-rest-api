package handlers

import (
	"github.com/Jerry2204/go-employee-rest-api/database"
	"github.com/Jerry2204/go-employee-rest-api/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllEmployees(c *fiber.Ctx) error {
	employees := []models.Employee{}
	database.DB.Db.Find(&employees)

	if len(employees) > 0 {
		return c.Status(200).JSON(employees)
	}

	return c.Status(200).JSON(fiber.Map{"data": "Employees is Empty"})
}