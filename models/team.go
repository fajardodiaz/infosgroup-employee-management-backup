package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	ID   int    `gorm:"serializer:json"`
	Name string `gorm:"not null;serializer:json"`
}
