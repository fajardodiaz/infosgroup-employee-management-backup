package models

import (
	"errors"

	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name string `gorm:"not null;serializer:json"`
}

func (t *Team) Validate() error {
	if t.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
