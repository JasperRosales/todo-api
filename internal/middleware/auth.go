package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/JasperRosales/todo-api/internal/services"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token required"})
			return
		}

		claims, err := services.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		userID := claims.UserID

		c.Set("userID", uint(userID))
		c.Next()
	}
}
