package handlers

import (
	"time"

	"github.com/Jerry2204/go-employee-rest-api/database"
	"github.com/Jerry2204/go-employee-rest-api/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func GetAllEmployees(c *fiber.Ctx) error {
	employees := []models.Employee{}

	database.DB.Db.Find(&employees)

	if len(employees) > 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "OK",
			"message": "Get Data Sucessfully!",
			"data":    employees,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": "Employees is Empty"})
}

func CreateEmployee(c *fiber.Ctx) error {
	employee := new(models.Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "Internal Server Error",
			"error":  err.Error(),
		})
	}

	if err := validate.Struct(employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "Bad Request",
			"error":  err.Error(),
		})
	}

	employee.HireDate = time.Now()

	database.DB.Db.Create(&employee)

	return c.Status(200).JSON(fiber.Map{
		"status":  "OK",
		"message": "Employee Created Successfully!",
		"data":    employee,
	})
}

func GetEmployeeById(c *fiber.Ctx) error {
	id := c.Params("id")
	employee := models.Employee{}

	result := database.DB.Db.First(&employee, id)

	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"status": "Data Not Found!", "error": "Employee not found"})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "OK",
		"message": "Get Data Sucessfully!",
		"data":    employee,
	})
}

func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	var employee models.Employee
	if err := database.DB.Db.First(&employee, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	if err := database.DB.Db.Save(&employee).Error; err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "OK",
		"message": "Data Updated Successfully!",
		"data":    employee,
	})
}

func DeleteEmpoyee(c *fiber.Ctx) error {
	id := c.Params("id")

	var employee models.Employee

	if err := database.DB.Db.First(&employee, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{ 
			"status": "error",
			"error":  "Data Not Found!",
		})
	}

	if err := database.DB.Db.Delete(&employee).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "OK",
		"message": "Data Deleted Succesfully",
		"data":    employee,
	})
}
