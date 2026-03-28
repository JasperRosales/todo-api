package dto

type BaseUserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PaginationRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type ListResponse[T any] struct {
	Data  []T   `json:"data"`
	Total int64 `json:"total"`
}