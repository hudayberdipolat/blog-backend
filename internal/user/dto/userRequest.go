package dto

type Register struct {
	FullName        string `json:"full_name"`
	PhoneNumber     string `json:"phone_number"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Login struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UpdateUser struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}

type ChangeUserPassword struct {
	OldPassword     string `json:"old_password"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
