package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name  string `gorm:"column:name; NOT NULL; unique"`
	Age   int    `gorm:"column:age"`
	Email string `gorm:"column:email; NOT NULL; unique"`
}

func (User) TableName() string {
	return "users"
}
