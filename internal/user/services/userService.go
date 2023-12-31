package services

import (
	"github.com/hudayberdipolat/blog-backend/internal/user/dto"
)

type UserService interface {
	RegisterUser(request dto.RegisterRequest) (*dto.UserResponse, error)
	LoginUser(request dto.LoginRequest) (*dto.UserResponse, error)
	GetUser(phoneNumber string) (*dto.GetUserResponse, error)
	UpdateUser(userID int, request dto.UpdateUserRequest) (*dto.UserResponse, error)
	PasswordChange(userID int, updatePasswordRequest dto.ChangeUserPasswordRequest) error
	UserDelete(userID int, phoneNumber string) error
}
