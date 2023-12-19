package services

import "github.com/hudayberdipolat/blog-backend/internal/categories/repositories"

type categoryServiceImp struct {
	categoryRepo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return categoryServiceImp{
		categoryRepo: repo,
	}
}
