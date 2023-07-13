package models

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	ID   int    `gorm:"serializer:json"`
	Name string `gorm:"not null;serializer:json"`
}
