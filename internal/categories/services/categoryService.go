package services

import "github.com/hudayberdipolat/blog-backend/internal/categories/dto"

type CategoryService interface {
	GetCategories() ([]dto.CategoryResponse, error)
	GetCategoryByID(categoryID int) (*dto.CategoryResponse, error)
	GetCategoryBySlug(categorySlug string) (*dto.CategoryResponse, error)
	Create(categoryRequest dto.CategoryRequest) (*dto.CategoryResponse, error)
	Update(categoryID int, categoryRequest dto.CategoryRequest) (*dto.CategoryResponse, error)
	Delete(categoryID int) error
}
