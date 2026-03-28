package dto

// CreateTodoRequest represents the request body for creating a new todo
type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required,min=1"`
	Description string `json:"description" binding:"required,min=1"`
}

// UpdateTodoRequest represents the request body for updating a todo (partial updates)
type UpdateTodoRequest struct {
	Title       *string `json:"title" binding:"omitempty,min=1"`
	Description *string `json:"description" binding:"omitempty,min=1"`
}
