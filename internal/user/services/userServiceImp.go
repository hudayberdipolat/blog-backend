package services

import (
	"errors"
	"github.com/hudayberdipolat/blog-backend/internal/user/dto"
	"github.com/hudayberdipolat/blog-backend/internal/user/helpers"
	"github.com/hudayberdipolat/blog-backend/internal/user/repositories"
)

type userServiceImp struct {
	repo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return userServiceImp{
		repo: userRepo,
	}
}

func (u userServiceImp) RegisterUser(authRequest dto.Register) (*dto.UserResponse, error) {
	password := helpers.GeneratePassword(authRequest.Password)
	authRequest.Password = password
	checkUser := u.repo.CheckPhoneNumber(authRequest.PhoneNumber)
	if checkUser != false {
		return nil, errors.New("Bu telefon belgisi eyyam ulanylyar!!!")
	}
	user, err := u.repo.CreateUser(authRequest)
	if err != nil {
		return nil, err
	}
	var userResponse dto.UserResponse
	userResponse.FullName = user.FullName
	userResponse.PhoneNumber = user.PhoneNumber
	return &userResponse, nil
}
