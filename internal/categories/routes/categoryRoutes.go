package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/blog-backend/internal/categories/handlers"
	"github.com/hudayberdipolat/blog-backend/internal/categories/repositories"
	"github.com/hudayberdipolat/blog-backend/internal/categories/services"
	"github.com/hudayberdipolat/blog-backend/pkg/middleware"
)

func CategoryRoutes(router *fiber.App) {
	categoryRepo := repositories.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepo)
	handler := handlers.NewCategoryHandler(categoryService)
	categoryRoute := router.Group("/categories")
	categoryRoute.Use(middleware.AuthMiddleware)
	categoryRoute.Get("", handler.GetAllCategories)
	//categoryRoute.Get("/:category_slug", handler.GetCategoryBySlug)
	categoryRoute.Get("/:category_id", handler.GetCategoryByID)

	categoryRoute.Post("/create", handler.CreateCategory)
	categoryRoute.Put("update/:category_id", handler.UpdateCategory)
	categoryRoute.Delete("delete/:category_id", handler.DeleteCategory)
}
