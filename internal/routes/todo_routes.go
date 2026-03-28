package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/JasperRosales/todo-api/internal/controllers"
	"github.com/JasperRosales/todo-api/internal/middleware"
)

func RegisterTodoRoutes(r *gin.Engine) {
	tc := controllers.NewTodoController()

	protected := r.Group("/todos")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/", tc.CreateTodo)
		protected.GET("/", tc.ListTodos)
		protected.GET("/:id", tc.GetTodo)
		protected.PATCH("/:id", tc.UpdateTodo)
		protected.DELETE("/:id", tc.DeleteTodo)
	}
}
