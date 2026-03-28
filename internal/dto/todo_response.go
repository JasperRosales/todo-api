package dto

// TodoResponse represents the response for a todo item
type TodoResponse struct {
	TodoID      uint   `json:"todo_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
}
