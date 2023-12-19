package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/blog-backend/internal/user/handlers"
	"github.com/hudayberdipolat/blog-backend/internal/user/repositories"
	"github.com/hudayberdipolat/blog-backend/internal/user/services"
	"github.com/hudayberdipolat/blog-backend/pkg/middleware"
)

// user-in oz edip bilmeli zatlary :

func UserRoutes(router *fiber.App) {

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	handler := handlers.NewHandler(userService)

	// user auth route begin
	router.Post("/register", handler.Register)
	router.Post("/login", handler.Login)
	// user auth route end
	userGroup := router.Group("/user")
	userGroup.Use(middleware.AuthMiddleware)
	userGroup.Get("/", handler.GetUser)
	userGroup.Put("/update", handler.UpdateUser)
	userGroup.Put("/changePassword", handler.ChangeUserPassword)
	userGroup.Delete("/delete", handler.DeleteUser)

}
