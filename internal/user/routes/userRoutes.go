package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/blog-backend/internal/user/handler"
)

// user-in oz edip bilmeli zatlary :

func UserRoutes(router *fiber.App) {

	handler := handler.NewHandler()

	// user auth route begin
	router.Post("/register", handler.GetUser)
	router.Post("/login", handler.GetUser)
	// user auth route end

	router.Get("/user", handler.GetUser)
	router.Get("/user/update", handler.UpdateUser)
	router.Get("/user/changePassword", handler.ChangeUserPassword)
	router.Get("/user/delete", handler.DeleteUser)

}
