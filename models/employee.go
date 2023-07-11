package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	CodEmpleado string     `json:"cod_empleado"`
	Name        *string    `json:"name"`
	Email       string     `json:"email"`
	Birth       *time.Time `json:"time"`
	IngressDate *time.Time `json:"ingress_date"`
	Phone       string     `json:"phone"`
	Gender      string     `json:"gender"`
	// Position
	// State
	// Project
}
