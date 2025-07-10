package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	Name     string    `gorm:"column:name; NOT NULL; unique"`
	Age      int       `gorm:"column:age"`
	Birthday time.Time `gorm:"column:birthday"`
}
