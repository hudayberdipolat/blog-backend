package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/blog-backend/internal/user/handlers"
	"github.com/hudayberdipolat/blog-backend/internal/user/repositories"
	"github.com/hudayberdipolat/blog-backend/internal/user/services"
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

	router.Get("/user", handler.GetUser)
	router.Get("/user/update", handler.UpdateUser)
	router.Get("/user/changePassword", handler.ChangeUserPassword)
	router.Get("/user/delete", handler.DeleteUser)

}
