package services

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/JasperRosales/todo-api/internal/database"
	"github.com/JasperRosales/todo-api/internal/dto"
	"github.com/JasperRosales/todo-api/internal/mapper"
	"github.com/JasperRosales/todo-api/internal/models"
)

type TodoService struct {
	db *gorm.DB
}

func NewTodoService() *TodoService {
	return &TodoService{
		db: database.GetDB(),
	}
}

func (s *TodoService) Create(ctx context.Context, req dto.CreateTodoRequest, userID uint) (dto.TodoResponse, error) {
	todo := mapper.ToTodoModel(req, userID)
	if err := s.db.WithContext(ctx).Create(todo).Error; err != nil {
		return dto.TodoResponse{}, fmt.Errorf("failed to create todo: %w", err)
	}
	return mapper.ToTodoResponse(*todo), nil
}

func (s *TodoService) Get(ctx context.Context, id uint, userID uint) (dto.TodoResponse, error) {
	var todo models.Todo
	if err := s.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		First(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.TodoResponse{}, fmt.Errorf("todo not found")
		}
		return dto.TodoResponse{}, fmt.Errorf("failed to get todo: %w", err)
	}
	return mapper.ToTodoResponse(todo), nil
}

func (s *TodoService) Update(ctx context.Context, id uint, req dto.UpdateTodoRequest, userID uint) (dto.TodoResponse, error) {
	var todo models.Todo
	if err := s.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		First(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.TodoResponse{}, fmt.Errorf("todo not found")
		}
		return dto.TodoResponse{}, fmt.Errorf("failed to get todo: %w", err)
	}

	mapper.ApplyTodoUpdates(&todo, req)

	if err := s.db.WithContext(ctx).Save(&todo).Error; err != nil {
		return dto.TodoResponse{}, fmt.Errorf("failed to update todo: %w", err)
	}

	return mapper.ToTodoResponse(todo), nil
}

func (s *TodoService) Delete(ctx context.Context, id uint, userID uint) error {
	if err := s.db.WithContext(ctx).
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&models.Todo{}).Error; err != nil {
		return fmt.Errorf("failed to delete todo: %w", err)
	}
	return nil
}

func (s *TodoService) List(ctx context.Context, req dto.PaginationRequest, userID uint) (dto.ListResponse[dto.TodoResponse], error) {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit < 1 || req.Limit > 100 {
		req.Limit = 10
	}
	offset := uint((req.Page - 1) * req.Limit)

	var todos []models.Todo
	var total int64

	tx := s.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Offset(int(offset)).
		Limit(req.Limit).
		Find(&todos)
	if err := tx.Count(&total).Error; err != nil {
		return dto.ListResponse[dto.TodoResponse]{}, fmt.Errorf("failed to count todos: %w", err)
	}
	if err := tx.Error; err != nil {
		return dto.ListResponse[dto.TodoResponse]{}, fmt.Errorf("failed to list todos: %w", err)
	}

	data := make([]dto.TodoResponse, len(todos))
	for i, t := range todos {
		data[i] = mapper.ToTodoResponse(t)
	}

	return dto.ListResponse[dto.TodoResponse]{
		Data:  data,
		Total: total,
	}, nil
}
