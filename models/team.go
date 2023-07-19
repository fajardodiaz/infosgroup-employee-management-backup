package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name string `gorm:"not null;serializer:json"`
}
