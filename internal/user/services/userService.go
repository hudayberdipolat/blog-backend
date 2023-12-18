package services

import "github.com/hudayberdipolat/blog-backend/internal/user/dto"

type UserService interface {
	RegisterUser(request dto.RegisterRequest) (*dto.UserResponse, error)
	LoginUser(request dto.LoginRequest) (*dto.UserResponse, error)
}
