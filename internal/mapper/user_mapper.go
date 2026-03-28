package mapper

import (
	"log"

	"github.com/JasperRosales/todo-api/internal/dto"
	"github.com/JasperRosales/todo-api/internal/models"
	"github.com/JasperRosales/todo-api/internal/utils"
)

// ToUserModel converts a CreateUserRequest DTO to a User model
func ToUserModel(req dto.CreateUserRequest) *models.User {
	user, err := models.NewUserBuilder().
		WithName(req.Name).
		WithEmail(req.Email).
		WithPassword(req.Password).
		Build()
	if err != nil {
		return nil
	}
	return user
}

// ApplyUserUpdates updates the User model with the fields provided in the UpdateUserRequest
func ApplyUserUpdates(user *models.User, req dto.UpdateUserRequest) {
	hasher := utils.NewPasswordHasher()

	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Password != nil {
		hashedPassword, err := hasher.Hash(*req.Password)
		if err != nil {
			log.Fatalf("failed to hash password: %v", err)
		}
		user.Password = hashedPassword
	}
}

// Return a UserResponse struct based on the User model
func ToUserResponse(user models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
