package models

import (
  "github.com/jinzhu/gorm"
)

type BoolToggler struct {
	gorm.Model
	Toggler bool `json:"toggler"`
	Users []*User `gorm:"many2many:user_boolTogglers;"`
}