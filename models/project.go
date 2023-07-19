package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name string `gorm:"not null;serializer:json"`
}
