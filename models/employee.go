package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	CodEmpleado       string     `gorm:"not null;serializer:json"`
	Name              *string    `gorm:"not null;serializer:json"`
	Email             string     `gorm:"serializer:json"`
	Birth             *time.Time `gorm:"serializer:json"`
	IngressDate       *time.Time `gorm:"serializer:json"`
	EndEvaluationDate *time.Time `gorm:"serializer:json"`
	Phone             string     `gorm:"serializer:json"`
	GenderID          int
	Gender            Gender
	PositionID        int
	Position          Position
	StateID           int
	State             State
	ProjectID         int
	Project           Project
	TeamID            int
	Team              Team
}
