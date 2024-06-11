package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Catlordx/CampusTrade/internal/utils"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Miss of Header Authorization"})
			c.Abort()
			return
		}
		token := authHeader[len("Bearer "):]
		claims, err := utils.VerifyToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的Token"})
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
