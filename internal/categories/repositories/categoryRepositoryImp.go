package repositories

import (
	"github.com/gosimple/slug"
	"github.com/hudayberdipolat/blog-backend/internal/categories/dto"
	"github.com/hudayberdipolat/blog-backend/internal/categories/models"
	"github.com/hudayberdipolat/blog-backend/pkg/database"
)

type categoryRepositoryImp struct {
}

func NewCategoryRepository() CategoryRepository {
	return categoryRepositoryImp{}
}

func (c categoryRepositoryImp) AllCategories() ([]models.Category, error) {
	var categories []models.Category
	result := database.DB.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (c categoryRepositoryImp) GetCategoryByID(categoryID int) (*models.Category, error) {
	var category models.Category
	result := database.DB.Where("id = ?", categoryID).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (c categoryRepositoryImp) GetCategoryBySlug(categorySlug string) (*models.Category, error) {
	var category models.Category
	result := database.DB.Where("category_slug = ?", categorySlug).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (c categoryRepositoryImp) CreateCategory(categoryRequest dto.CategoryRequest) (*models.Category, error) {
	var category models.Category
	slugCategory := slug.Make(categoryRequest.CategoryName)
	category.CategorySlug = slugCategory
	category.CategoryName = categoryRequest.CategoryName
	result := database.DB.Create(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (c categoryRepositoryImp) UpdateCategory(categoryID int, categoryRequest dto.CategoryRequest) (*models.Category, error) {
	var category models.Category
	slugCategory := slug.Make(categoryRequest.CategoryName)
	// create category slug end
	result := database.DB.Model(&category).Where("id=?", categoryID).Updates(models.Category{
		CategoryName: categoryRequest.CategoryName,
		CategorySlug: slugCategory,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (c categoryRepositoryImp) DeleteCategory(categoryID int) error {
	var category models.Category

	result := database.DB.Unscoped().Where("id = ?", categoryID).Delete(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (c categoryRepositoryImp) CheckCategoryName(categoryName string) bool {
	var category models.Category
	database.DB.Where("category_name=?", categoryName).First(&category)
	if category.ID == 0 {
		return false
	}
	return true
}
