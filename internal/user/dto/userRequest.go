package dto

type Register struct {
	FullName        string `json:"full_name" validate:"required"`
	PhoneNumber     string `json:"phone_number" validate:"required,gt=6"`
	Password        string `json:"password" validate:"required,gt=5"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type Login struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type UpdateUser struct {
	FullName    string `json:"full_name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type ChangeUserPassword struct {
	OldPassword     string `json:"old_password" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}
