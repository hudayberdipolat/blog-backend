package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/blog-backend/internal/categories/dto"
	"github.com/hudayberdipolat/blog-backend/internal/categories/services"
	"net/http"
	"strconv"
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
	categories, err := h.categoryService.GetCategories()
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"error":   err.Error(),
			"message": "get all  categories error",
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":     http.StatusOK,
		"categories": categories,
		"message":    "get all  categories successfully",
	})
}

func (h categoryHandler) GetCategoryByID(ctx *fiber.Ctx) error {
	id := ctx.Params("category_id")
	categoryID, _ := strconv.Atoi(id)

	getCategory, err := h.categoryService.GetCategoryByID(categoryID)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"error":   err.Error(),
			"message": "get category By id error",
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":   http.StatusOK,
		"category": getCategory,
		"message":  "get category successfully",
	})
}

//
//func (h categoryHandler) GetCategoryBySlug(ctx *fiber.Ctx) error {
//	categorySlug := ctx.Params("category_slug")
//	getCategory, err := h.categoryService.GetCategoryBySlug(categorySlug)
//	if err != nil {
//		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
//			"status":  http.StatusNotFound,
//			"error":   err.Error(),
//			"message": "category not found error",
//		})
//	}
//	return ctx.Status(http.StatusOK).JSON(fiber.Map{
//		"status":   http.StatusOK,
//		"category": getCategory,
//		"message":  "get category successfully with category slug",
//	})
//}

func (h categoryHandler) CreateCategory(ctx *fiber.Ctx) error {

	var createCategoryRequest dto.CategoryRequest

	// body parser
	if err := ctx.BodyParser(&createCategoryRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
			"message": "error body parser",
		})
	}
	// validation
	validate := validator.New()
	if validateError := validate.Struct(&createCategoryRequest); validateError != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   validateError.Error(),
			"message": "validate error",
		})
	}
	// create  category
	categoryResponse, createError := h.categoryService.Create(createCategoryRequest)
	if createError != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"error":   createError.Error(),
			"message": "create category error",
		})
	}
	// return response
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"status":   http.StatusCreated,
		"category": categoryResponse,
		"message":  "category created successfully",
	})
}

func (h categoryHandler) UpdateCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("category_id")
	categoryID, _ := strconv.Atoi(id)
	var updateCategoryRequest dto.CategoryRequest
	// body parser
	if err := ctx.BodyParser(&updateCategoryRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
			"message": "error body parser",
		})
	}
	// validation
	validate := validator.New()
	if validateError := validate.Struct(&updateCategoryRequest); validateError != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   validateError.Error(),
			"message": "validate error",
		})
	}

	// update category
	updateCategoryResponse, updateError := h.categoryService.Update(categoryID, updateCategoryRequest)
	if updateError != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   updateError.Error(),
			"message": "get category By id error",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":   http.StatusOK,
		"category": updateCategoryResponse,
		"message":  "category updated successfully",
	})
}

func (h categoryHandler) DeleteCategory(ctx *fiber.Ctx) error {
	id := ctx.Params("category_id")
	categoryID, _ := strconv.Atoi(id)
	deleteCategoryError := h.categoryService.Delete(categoryID)
	if deleteCategoryError != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   deleteCategoryError.Error(),
			"message": "category not deleted",
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "category deleted successfully",
	})
}
