package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string `json:"category_name"`
	CategorySlug string `json:"category_slug"`
}

func (*Category) TableName() string {
	return "categories"
}
