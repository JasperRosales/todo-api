package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/JasperRosales/todo-api/internal/controllers"
	"github.com/JasperRosales/todo-api/internal/middleware"
)

func RegisterUserRoutes(r *gin.Engine) {
	userController := controllers.NewUserController()
	v1 := r.Group("/api/v1")
	users := v1.Group("/users")

	// Public routes
	users.POST("/register", userController.Register)
	users.POST("/login", userController.LoginUser)

	// Protected routes
	protected := users.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("", userController.ListUsers)
		protected.GET("/:id", userController.GetUser)
		protected.PATCH("/:id", userController.UpdateUser)
		protected.DELETE("/:id", userController.DeleteUser)
	}
}
