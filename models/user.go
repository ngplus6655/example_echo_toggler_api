package models

import (
  "github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	// FavoritesでboolTogglerのスライスを保持
	Favorites []BoolToggler `gorm:"many2many:user_favorite_togglers;"`
}