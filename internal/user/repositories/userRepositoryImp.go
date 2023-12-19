package repositories

import (
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

	if result := database.DB.Where("phone_number=?", phoneNumber).First(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (u userRepositoryImp) UpdatePhoneNumber(userID int, phoneNumber string) bool {
	var user models.User
	database.DB.Where("id != ?", userID).Where("phone_number=?", phoneNumber).First(&user)
	if user.ID == 0 {
		return false
	}
	return true
}

func (u userRepositoryImp) UserUpdate(userID int, updateRequest dto.UpdateUserRequest) (*models.User, error) {
	// update user
	var user models.User
	result := database.DB.Model(&user).Where("id = ?", userID).Updates(models.User{
		FullName:    updateRequest.FullName,
		PhoneNumber: updateRequest.PhoneNumber,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u userRepositoryImp) GetUserByID(userID int) (*models.User, error) {
	var user models.User
	result := database.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u userRepositoryImp) ChangePassword(userID int, newPassword string) error {
	//update password
	var user models.User
	result := database.DB.Model(&user).Where("id=?", userID).Updates(models.User{Password: newPassword})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u userRepositoryImp) DeleteUser(userID int, phoneNumber string) error {
	var user models.User
	result := database.DB.Unscoped().Where("id = ?", userID).Where("phone_number =?", phoneNumber).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
