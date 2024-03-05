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

func CreateEmployee(c *fiber.Ctx) error {
	employee := new(models.Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&employee)

	return c.Status(200).JSON(employee)
}

func GetEmployeeById(c *fiber.Ctx) error {
	id := c.Params("id")
	employee := models.Employee{}

	result := database.DB.Db.First(&employee, id)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Employee not found"})
	}

	return c.Status(200).JSON(employee)
}
