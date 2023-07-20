package models

import (
	"errors"

	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	Name string `gorm:"not null;serializer:json"`
}

func (s *State) Validate() error {
	if s.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
