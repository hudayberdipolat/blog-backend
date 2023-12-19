package dto

type CategoryRequest struct {
	CategoryName string `json:"category_name" validate:"required,min=2"`
}
