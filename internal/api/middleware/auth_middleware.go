package middleware

import (
	"Proyectos_Go/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("auth_token")

		if err != nil {
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" || len(authHeader) < 8 {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "No se encontró token de autorización"})
				c.Abort()
				return
			}
			tokenString = authHeader[7:]
		}
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
