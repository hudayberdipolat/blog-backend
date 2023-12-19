package dto

type UserResponse struct {
	ID          uint   `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	AccessToken string `json:"access_token"`
}

type GetUserResponse struct {
	ID          uint   `json:"id"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
}
