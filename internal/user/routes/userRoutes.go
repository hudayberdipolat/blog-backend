package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/blog-backend/internal/user/handler"
)

func UserRoutes(router *fiber.App) {

	handler := handler.NewHandler()

	router.Get("/user", handler.GetUser)

}
