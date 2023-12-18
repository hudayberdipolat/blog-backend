package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName    string `gorm:"not null"`
	PhoneNumber string `gorm:"not null,unique"`
	Password    string `gorm:"not null"`
}

func (*User) TableName() string {
	return "users"
}
