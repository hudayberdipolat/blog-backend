package repositories

import (
	"errors"
	"github.com/hudayberdipolat/blog-backend/internal/user/dto"
	"github.com/hudayberdipolat/blog-backend/internal/user/models"
	"github.com/hudayberdipolat/blog-backend/pkg/database"
)

type userRepositoryImp struct{}

func NewUserRepository() UserRepository {
	return userRepositoryImp{}
}

func (u userRepositoryImp) CreateUser(registerUser dto.RegisterRequest) (*models.User, error) {
	var user models.User
	user.FullName = registerUser.FullName
	user.PhoneNumber = registerUser.PhoneNumber
	user.Password = registerUser.Password

	if result := database.DB.Create(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u userRepositoryImp) CheckPhoneNumber(phoneNumber string) bool {
	var user models.User
	database.DB.Where("phone_number=?", phoneNumber).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}

func (u userRepositoryImp) GetByUser(phoneNumber string) (*models.User, error) {
	var user models.User
	database.DB.Where("phone_number=?", phoneNumber).First(&user)
	if user.ID == 0 {
		return nil, errors.New("Phone number yada password nadogry!!!")
	}
	return &user, nil
}
