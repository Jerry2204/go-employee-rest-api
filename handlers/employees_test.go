package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetAllEmployees(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "get HTTP status code ",
			route:        "/employees",
			expectedCode: 200,
		},
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/employees/not-found",
			expectedCode: 404,
		},
	}

	app := fiber.New()

	app.Get("/employees", GetAllEmployees)

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)

		resp, _ := app.Test(req, 1)

		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
