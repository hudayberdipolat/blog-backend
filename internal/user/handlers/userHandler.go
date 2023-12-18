package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/blog-backend/internal/user/dto"
	"github.com/hudayberdipolat/blog-backend/internal/user/services"
)

type handler struct {
	service services.UserService
}

func NewHandler(userService services.UserService) handler {
	return handler{
		service: userService,
	}
}

// user Auth begin
func (h handler) Register(ctx *fiber.Ctx) error {
	var registerRequest dto.RegisterRequest
	// body parser
	if err := ctx.BodyParser(&registerRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
			"message": "Body parser error",
		})
	}
	// validate
	validate := validator.New()
	if validateError := validate.Struct(&registerRequest); validateError != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   validateError.Error(),
			"message": "Validation error",
		})
	}

	// register user

	userResponse, registerError := h.service.RegisterUser(registerRequest)
	if registerError != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"error":   registerError.Error(),
			"message": "user not created...",
		})
	}
	// return response
	// generate access token
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"status":       http.StatusCreated,
		"userResponse": userResponse,
		"message":      "user  created successfully...",
	})
}

func (h handler) Login(ctx *fiber.Ctx) error {
	var loginRequest dto.LoginRequest
	// body parser
	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
			"message": "Body parser error",
		})
	}
	// validate
	validate := validator.New()
	if validateError := validate.Struct(&loginRequest); validateError != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   validateError.Error(),
			"message": "Validation error",
		})
	}

	// check user -->> user barlag
	userResponse, err := h.service.LoginUser(loginRequest)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
			"message": "giris maglumatlary nadogry...",
		})
	}
	// generate access token

	//return response
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"status":       http.StatusCreated,
		"userResponse": userResponse,
		"message":      "user  login successfully...",
	})
}

// user Auth end

func (h handler) GetUser(ctx *fiber.Ctx) error {

	return nil
}

func (h handler) GetAllUsers(ctx *fiber.Ctx) error {

	return nil
}

func (h handler) CreateUser(ctx *fiber.Ctx) error {

	return nil
}

func (h handler) UpdateUser(ctx *fiber.Ctx) error {

	return nil
}

func (h handler) DeleteUser(ctx *fiber.Ctx) error {

	return nil
}

func (h handler) ChangeUserPassword(ctx *fiber.Ctx) error {

	return nil
}
