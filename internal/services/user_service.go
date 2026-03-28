package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"github.com/JasperRosales/todo-api/internal/dto"
	"github.com/JasperRosales/todo-api/internal/mapper"
	"github.com/JasperRosales/todo-api/internal/models"
	"github.com/JasperRosales/todo-api/internal/utils"
)

type UserService struct {
	db     *gorm.DB
	hasher *utils.PasswordHasher
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db:     db,
		hasher: utils.NewPasswordHasher(),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.UserResponse, error) {
	hashedPassword, err := s.hasher.Hash(req.Password)
	if err != nil {
		return dto.UserResponse{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user := mapper.ToUserModel(req)
	if user == nil {
		return dto.UserResponse{}, errors.New("failed to build user")
	}
	user.Password = hashedPassword

	if err := s.db.WithContext(ctx).Create(user).Error; err != nil {
		return dto.UserResponse{}, fmt.Errorf("failed to create user: %w", err)
	}

	return mapper.ToUserResponse(*user), nil
}

func (s *UserService) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	var user models.User
	if err := s.db.WithContext(ctx).Where("email = ?", req.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("invalid credentials")
		}
		return "", fmt.Errorf("failed to get user: %w", err)
	}

	if ok, err := s.hasher.Check(user.Password, req.Password); err != nil || !ok {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("2c3f51302b8a5170cde1bd9165a5ca4f07e675dbf0f634b60d14018f2485bd5e"))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

func (s *UserService) GetUser(ctx context.Context, id uint) (dto.UserResponse, error) {
	var user models.User
	if err := s.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UserResponse{}, fmt.Errorf("user not found")
		}
		return dto.UserResponse{}, fmt.Errorf("failed to get user: %w", err)
	}
	return mapper.ToUserResponse(user), nil
}

func (s *UserService) UpdateUser(ctx context.Context, id uint, req dto.UpdateUserRequest) (dto.UserResponse, error) {
	var user models.User
	if err := s.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UserResponse{}, fmt.Errorf("user not found")
		}
		return dto.UserResponse{}, fmt.Errorf("failed to get user: %w", err)
	}

	mapper.ApplyUserUpdates(&user, req)

	if req.Password != nil {
		hashedPassword, err := s.hasher.Hash(*req.Password)
		if err != nil {
			return dto.UserResponse{}, fmt.Errorf("failed to hash password: %w", err)
		}
		user.Password = hashedPassword
	}

	if err := s.db.WithContext(ctx).Save(&user).Error; err != nil {
		return dto.UserResponse{}, fmt.Errorf("failed to update user: %w", err)
	}

	return mapper.ToUserResponse(user), nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uint) error {
	if err := s.db.WithContext(ctx).Delete(&models.User{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (s *UserService) ListUsers(ctx context.Context, req dto.PaginationRequest) (dto.ListResponse[dto.UserResponse], error) {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Limit < 1 || req.Limit > 100 {
		req.Limit = 10
	}
	offset := uint((req.Page - 1) * req.Limit)

	var users []models.User
	var total int64

	tx := s.db.WithContext(ctx).Offset(int(offset)).Limit(req.Limit).Find(&users)
	if err := tx.Count(&total).Error; err != nil {
		return dto.ListResponse[dto.UserResponse]{}, fmt.Errorf("failed to count users: %w", err)
	}
	if err := tx.Error; err != nil {
		return dto.ListResponse[dto.UserResponse]{}, fmt.Errorf("failed to list users: %w", err)
	}

	data := make([]dto.UserResponse, len(users))
	for i, u := range users {
		data[i] = mapper.ToUserResponse(u)
	}

	return dto.ListResponse[dto.UserResponse]{
		Data:  data,
		Total: total,
	}, nil
}
