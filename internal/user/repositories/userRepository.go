package repositories

import (
	"github.com/hudayberdipolat/blog-backend/internal/user/dto"
	"github.com/hudayberdipolat/blog-backend/internal/user/models"
)

type UserRepository interface {
	CreateUser(registerUser dto.Register) (*models.User, error)
	CheckPhoneNumber(phoneNumber string) bool
}
