package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAllEmployees(t *testing.T) {
	tests := struct {
		description  string
		route        string
		expectedCode int
	}{
		description:  "get HTTP status code ",
		route:        "/employees",
		expectedCode: 200,
	}

	app := fiber.New()

	app.Get("/employees", GetAllEmployees)

	req := httptest.NewRequest("GET", tests.route, nil)

	resp, _ := app.Test(req, 1)

	assert.Equal(t, tests.expectedCode, resp.StatusCode, tests.description)
}

func TestDeleteEmplooyee(t *testing.T) {
	app := fiber.New()

	app.Delete("/employees/:id", DeleteEmpoyee)

	req := httptest.NewRequest("DELETE", "/employees/1", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}
