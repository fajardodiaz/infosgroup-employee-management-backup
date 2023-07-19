package models

import (
	"errors"

	"gorm.io/gorm"
)

type Position struct {
	gorm.Model
	Name string `gorm:"not null;serializer:json"`
}

func (p *Position) Validate() error {
	if p.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
