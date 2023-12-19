package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string `json:"category_name" gorm:"not null,unique"`
	CategorySlug string `json:"category_slug" gorm:"not null,unique"`
}

func (*Category) TableName() string {
	return "categories"
}
