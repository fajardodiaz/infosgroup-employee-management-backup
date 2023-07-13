package models

import "gorm.io/gorm"

type Gender struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
}
