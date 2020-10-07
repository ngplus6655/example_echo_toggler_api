package models

import (
  "github.com/jinzhu/gorm"
)

type BoolToggler struct {
	// type gorm.Model struct {
			// ID        uint           `gorm:"primaryKey"`
			// CreatedAt time.Time
			// UpdatedAt time.Time
			// DeletedAt gorm.DeletedAt `gorm:"index"`
	// }
	gorm.Model
  Toggler bool `json:"toggler"`
}