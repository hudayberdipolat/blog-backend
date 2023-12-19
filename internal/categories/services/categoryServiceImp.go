package services

import (
	"errors"
	"github.com/hudayberdipolat/blog-backend/internal/categories/dto"
	"github.com/hudayberdipolat/blog-backend/internal/categories/helpers"
	"github.com/hudayberdipolat/blog-backend/internal/categories/repositories"
)

type categoryServiceImp struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return categoryServiceImp{
		categoryRepo: repo,
	}
}

func (c categoryServiceImp) GetCategories() ([]dto.CategoryResponse, error) {
	categories, err := c.categoryRepo.AllCategories()
	if err != nil {
		return nil, err
	}
	var getCategories []dto.CategoryResponse
	for _, category := range categories {
		getCategory := helpers.NewCategoryResponse(category)
		getCategories = append(getCategories, getCategory)
	}
	return getCategories, nil
}

func (c categoryServiceImp) GetCategoryByID(categoryID int) (*dto.CategoryResponse, error) {
	category, err := c.categoryRepo.GetCategoryByID(categoryID)
	if err != nil {
		return nil, err
	}
	getCategory := helpers.NewCategoryResponse(*category)
	return &getCategory, nil
}

func (c categoryServiceImp) GetCategoryBySlug(categorySlug string) (*dto.CategoryResponse, error) {
	category, err := c.categoryRepo.GetCategoryBySlug(categorySlug)
	if err != nil {
		return nil, err
	}
	getCategory := helpers.NewCategoryResponse(*category)
	return &getCategory, nil
}

func (c categoryServiceImp) Create(categoryRequest dto.CategoryRequest) (*dto.CategoryResponse, error) {

	checkCategoryName := c.categoryRepo.CheckCategoryName(categoryRequest.CategoryName)
	if checkCategoryName != false {
		return nil, errors.New("Bu categoryName eyyam bar!!!")
	}
	createCategory, errCategoryCreate := c.categoryRepo.CreateCategory(categoryRequest)
	if errCategoryCreate != nil {
		return nil, errCategoryCreate
	}
	getCategoryResponse := helpers.NewCategoryResponse(*createCategory)
	return &getCategoryResponse, nil
}

func (c categoryServiceImp) Update(categoryID int, categoryRequest dto.CategoryRequest) (*dto.CategoryResponse, error) {
	checkCategoryName := c.categoryRepo.CheckCategoryName(categoryRequest.CategoryName)
	if checkCategoryName != false {
		return nil, errors.New("Bu categoryName eyyam bar!!!")
	}
	updateCategory, updateCategoryError := c.categoryRepo.UpdateCategory(categoryID, categoryRequest)
	if updateCategoryError != nil {
		return nil, updateCategoryError
	}
	updateCategory.ID = uint(categoryID)
	getUpdatedCategory := helpers.NewCategoryResponse(*updateCategory)
	return &getUpdatedCategory, nil
}

func (c categoryServiceImp) Delete(categoryID int) error {
	deleteError := c.categoryRepo.DeleteCategory(categoryID)
	if deleteError != nil {
		return deleteError
	}
	return nil
}
