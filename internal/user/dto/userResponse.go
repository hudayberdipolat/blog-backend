package dto

type UserResponse struct {
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	AccessToken string `json:"access_token"`
}
