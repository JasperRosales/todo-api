package mapper

import (
	"github.com/JasperRosales/todo-api/internal/dto"
	"github.com/JasperRosales/todo-api/internal/models"
)

// ToTodoModel converts CreateTodoRequest to Todo model using builder
func ToTodoModel(req dto.CreateTodoRequest, userID uint) *models.Todo {
	return models.NewTodoBuilder().
		WithTitle(req.Title).
		WithDescription(req.Description).
		WithUserID(userID).
		Build()
}

// ToTodoResponse converts Todo model to TodoResponse
func ToTodoResponse(todo models.Todo) dto.TodoResponse {
	return dto.TodoResponse{
		TodoID:      todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		UserID:      todo.UserID,
	}
}

// ApplyTodoUpdates applies partial updates from UpdateTodoRequest to Todo model
func ApplyTodoUpdates(todo *models.Todo, req dto.UpdateTodoRequest) {
	if req.Title != nil {
		todo.Title = *req.Title
	}
	if req.Description != nil {
		todo.Description = *req.Description
	}
}
