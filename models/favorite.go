package models

import (
  "github.com/jinzhu/gorm"
)

type Favorite struct {
	gorm.Model
	TogglerId uint `json:"toggler_id"`
	UserId uint `json:"toggler_id"`
}