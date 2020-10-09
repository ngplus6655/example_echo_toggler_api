package models

import (
  "github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	// UserはBoolTogglerを所有します
	BoolTogglers []*BoolToggler `gorm:"many2many:user_boolTogglers;"`
}