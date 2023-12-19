package dto

type CategoryResponse struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	CategorySlug string `json:"category_slug"`
}
