package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/JasperRosales/todo-api/internal/database"
	"github.com/JasperRosales/todo-api/internal/routes"
)


func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Goods nmn ang server idol",
		})
	})

	if err := database.InitDB(); err != nil {
		panic(err)
	}

	routes.RegisterUserRoutes(r)
	routes.RegisterTodoRoutes(r)

	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

}
