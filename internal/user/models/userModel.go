package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username    string
	FullName    string
	Email       string
	PhoneNumber string
	Password    string
}

func (*User) TableName() string {
	return "users"
}

func (user *User) BeforeCreate() error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(password)
	return nil
}
