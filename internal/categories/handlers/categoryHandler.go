package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/blog-backend/internal/categories/services"
)

type categoryHandler struct {
	categoryService services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) categoryHandler {
	return categoryHandler{
		categoryService: service,
	}
}

func (h categoryHandler) GetAllCategories(ctx *fiber.Ctx) error {
	return nil
}

func (h categoryHandler) GetCategory(ctx *fiber.Ctx) error {
	return nil
}

func (h categoryHandler) CreateCategory(ctx *fiber.Ctx) error {
	return nil
}

func (h categoryHandler) UpdateCategory(ctx *fiber.Ctx) error {
	return nil
}

func (h categoryHandler) DeleteCategory(ctx *fiber.Ctx) error {
	return nil
}
