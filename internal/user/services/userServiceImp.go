package services

import (
	"errors"
	"github.com/hudayberdipolat/blog-backend/internal/user/dto"
	"github.com/hudayberdipolat/blog-backend/internal/user/helpers"
	"github.com/hudayberdipolat/blog-backend/internal/user/repositories"
	"github.com/hudayberdipolat/blog-backend/pkg/generateToken"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImp struct {
	repo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return userServiceImp{
		repo: userRepo,
	}
}

func (u userServiceImp) RegisterUser(authRequest dto.RegisterRequest) (*dto.UserResponse, error) {
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

	accessToken, errToken := generateToken.GenerateToken(user.PhoneNumber, int(user.ID))
	if errToken != nil {
		return nil, errToken
	}
	var userResponse dto.UserResponse
	userResponse.ID = user.ID
	userResponse.FullName = user.FullName
	userResponse.PhoneNumber = user.PhoneNumber
	userResponse.AccessToken = accessToken
	return &userResponse, nil
}

func (u userServiceImp) LoginUser(request dto.LoginRequest) (*dto.UserResponse, error) {
	getUser, err := u.repo.GetByUser(request.PhoneNumber)
	if err != nil {
		return nil, err
	}
	errCheckPassword := helpers.CheckPassword(getUser.Password, request.Password)
	if errCheckPassword != nil {
		return nil, errCheckPassword
	}
	accessToken, errToken := generateToken.GenerateToken(getUser.PhoneNumber, int(getUser.ID))
	if errToken != nil {
		return nil, errToken
	}
	var userResponse dto.UserResponse
	userResponse.ID = getUser.ID
	userResponse.FullName = getUser.FullName
	userResponse.PhoneNumber = getUser.PhoneNumber
	userResponse.AccessToken = accessToken
	return &userResponse, nil
}

func (u userServiceImp) GetUser(phoneNumber string) (*dto.GetUserResponse, error) {
	getUser, err := u.repo.GetByUser(phoneNumber)
	if err != nil {
		return nil, err
	}
	var userResponse dto.GetUserResponse
	userResponse.ID = getUser.ID
	userResponse.FullName = getUser.FullName
	userResponse.PhoneNumber = getUser.PhoneNumber
	return &userResponse, nil
}

func (u userServiceImp) UpdateUser(userID int, request dto.UpdateUserRequest) (*dto.UserResponse, error) {
	// userin phone numberin on barlygyny yada yoklugyny barlamaly servicede yerine yetirmeli
	checkUser := u.repo.UpdatePhoneNumber(userID, request.PhoneNumber)
	if checkUser != false {
		return nil, errors.New("Bu telefon belgisi eyyam ulanylyar!!!")
	}
	// update user
	updateUser, errUpdate := u.repo.UserUpdate(userID, request)
	if errUpdate != nil {
		return nil, errUpdate
	}

	// generate token
	accessToken, errToken := generateToken.GenerateToken(updateUser.PhoneNumber, userID)
	if errToken != nil {
		return nil, errToken
	}
	// return user response
	var userResponse dto.UserResponse
	userResponse.FullName = updateUser.FullName
	userResponse.PhoneNumber = updateUser.PhoneNumber
	userResponse.AccessToken = accessToken
	return &userResponse, nil
}

func (u userServiceImp) PasswordChange(userID int, updatePasswordRequest dto.ChangeUserPasswordRequest) error {
	/*
		onki bar bolan password bilen old password gelyan data denmi sol ikisini denesdirmeli get user
		bilen user datany aldyrmaly --> useri id bilen cekdirmeli we sol userin passwordyny requestden
		gelen password bilen denesdirmeli
	*/

	// get user
	getUser, err := u.repo.GetUserByID(userID)
	if err != nil {
		return err
	}
	// old password barlag
	errHashPasswordBarlag := bcrypt.CompareHashAndPassword([]byte(getUser.Password), []byte(updatePasswordRequest.OldPassword))
	if errHashPasswordBarlag != nil {
		return errors.New("Kone passwordynyz nadogry!!!")
	}
	// update user password
	newPassword := helpers.GeneratePassword(updatePasswordRequest.Password)
	errUpdatePassword := u.repo.ChangePassword(userID, newPassword)
	if errUpdatePassword != nil {
		return errUpdatePassword
	}
	// return data
	return nil
}
