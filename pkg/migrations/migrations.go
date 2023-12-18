package migrations

import (
	"github.com/hudayberdipolat/blog-backend/internal/user/models"
	"github.com/hudayberdipolat/blog-backend/pkg/database"
)

func MigrateTable() error {
	err := database.DB.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	return nil
}
