package helpers

import (
	"github.com/hudayberdipolat/blog-backend/internal/categories/dto"
	"github.com/hudayberdipolat/blog-backend/internal/categories/models"
)

func NewCategoryResponse(category models.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		ID:           int(category.ID),
		CategoryName: category.CategoryName,
		CategorySlug: category.CategorySlug,
	}
}
