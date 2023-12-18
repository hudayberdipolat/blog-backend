package services

import "github.com/hudayberdipolat/blog-backend/internal/user/dto"

type UserService interface {
	RegisterUser(request dto.Register) (*dto.UserResponse, error)
}
