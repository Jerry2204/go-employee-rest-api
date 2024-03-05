package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName string `json:"firstName" gorm:"text;not null;default:null"`
	LastName  string `json:"lastName" gorm:"text;not null;default:null"`
	Email     string `json:"email" gorm:"text;not null;default:null"`
	HireDate  time.Time
}
