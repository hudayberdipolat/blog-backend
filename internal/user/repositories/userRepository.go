package repositories

import (
	"github.com/hudayberdipolat/blog-backend/internal/user/dto"
	"github.com/hudayberdipolat/blog-backend/internal/user/models"
)

type UserRepository interface {
	CreateUser(registerUser dto.RegisterRequest) (*models.User, error)
	CheckPhoneNumber(phoneNumber string) bool
	GetByUser(phoneNumber string) (*models.User, error)
}
