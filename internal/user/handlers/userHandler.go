package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hudayberdipolat/blog-backend/internal/user/dto"
	"github.com/hudayberdipolat/blog-backend/internal/user/services"
	"net/http"
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
			"message": "phone number yada password nadogry",
		})
	}
	//return response
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":       http.StatusOK,
		"userResponse": userResponse,
		"message":      "user  login successfully...",
	})
}

// user Auth end

// user information

func (h handler) GetUser(ctx *fiber.Ctx) error {
	phone := ctx.Locals("phone_number")
	phoneNumber := phone.(string)
	getUser, err := h.service.GetUser(phoneNumber)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"error":   err.Error(),
			"message": "get user error...",
		})
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"user":    getUser,
		"message": "get user successfully ...",
	})
}

func (h handler) UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Locals("user_id")
	userID := id.(int)
	var updateUserRequest dto.UpdateUserRequest

	// body parser
	if err := ctx.BodyParser(&updateUserRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
			"message": "body parser error",
		})
	}
	// validation
	validate := validator.New()
	if errValidate := validate.Struct(&updateUserRequest); errValidate != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   errValidate.Error(),
			"message": "validator error",
		})
	}
	// update user
	userResponse, errUserUpdate := h.service.UpdateUser(userID, updateUserRequest)
	if errUserUpdate != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   errUserUpdate.Error(),
			"message": "update user  error",
		})
	}
	// return response
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":    http.StatusOK,
		"message":   "user update successfully",
		"user_id":   userID,
		"user_data": userResponse,
	})
}

func (h handler) ChangeUserPassword(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int)
	var changePasswordRequest dto.ChangeUserPasswordRequest
	// body parser
	if err := ctx.BodyParser(&changePasswordRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
			"message": "body parser error",
		})
	}
	// validation
	validate := validator.New()
	if validateError := validate.Struct(&changePasswordRequest); validateError != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"error":   validateError.Error(),
			"message": "change password validate error",
		})
	}
	// update password
	errorUserPasswordChange := h.service.PasswordChange(userID, changePasswordRequest)
	if errorUserPasswordChange != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"error":   errorUserPasswordChange.Error(),
			"message": "change password error",
		})
	}
	//return response
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "user Password Changed Successfully",
	})
}

func (h handler) DeleteUser(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int)
	phoneNumber := ctx.Locals("phone_number").(string)
	// delete user
	deleteUserError := h.service.UserDelete(userID, phoneNumber)
	if deleteUserError != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "user not deleted",
			"error":   deleteUserError.Error(),
		})
	}
	ctx.Set("Authorization", "")
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "user deleted successfully",
	})
}
