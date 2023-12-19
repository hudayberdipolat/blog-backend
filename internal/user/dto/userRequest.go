package dto

type RegisterRequest struct {
	FullName        string `json:"full_name" validate:"required,min=4"`
	PhoneNumber     string `json:"phone_number" validate:"required,min=6"`
	Password        string `json:"password" validate:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	FullName    string `json:"full_name" validate:"required,min=4"`
	PhoneNumber string `json:"phone_number" validate:"required,min=6"`
}

type ChangeUserPasswordRequest struct {
	OldPassword     string `json:"old_password" validate:"required"`
	Password        string `json:"password" validate:"required,min=5"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
