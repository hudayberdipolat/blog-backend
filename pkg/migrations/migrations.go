package migrations

import (
	categoryModel "github.com/hudayberdipolat/blog-backend/internal/categories/models"
	userModel "github.com/hudayberdipolat/blog-backend/internal/user/models"
	"github.com/hudayberdipolat/blog-backend/pkg/database"
)

func MigrateTable() error {
	err := database.DB.AutoMigrate(&userModel.User{}, &categoryModel.Category{})
	if err != nil {
		return err
	}
	return nil
}
