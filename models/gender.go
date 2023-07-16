package models

import (
	"errors"

	"gorm.io/gorm"
)

type Gender struct {
	gorm.Model
	Name string `gorm:"serializer:json;not null;type:varchar(15)"`
}

func (g *Gender) Validate() error {
	if g.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
