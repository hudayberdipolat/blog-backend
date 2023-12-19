package repositories

import (
	"github.com/hudayberdipolat/blog-backend/internal/categories/dto"
	"github.com/hudayberdipolat/blog-backend/internal/categories/models"
)

type CategoryRepository interface {
	AllCategories() ([]models.Category, error)
	GetCategoryByID(categoryID int) (*models.Category, error)
	GetCategoryBySlug(categorySlug string) (*models.Category, error)
	CreateCategory(categoryRequest dto.CategoryRequest) (*models.Category, error)
	UpdateCategory(categoryID int, categoryRequest dto.CategoryRequest) (*models.Category, error)
	DeleteCategory(categoryID int) error
	CheckCategoryName(categoryName string) bool
}
