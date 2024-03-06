package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	HireDate  time.Time `json:"hireDate"`
}
