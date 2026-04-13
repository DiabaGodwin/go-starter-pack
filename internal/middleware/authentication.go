package middleware

import (
	"net/http"
	"real-estate-app/internal/service/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

const CurrentUserKey = "currentUser"

func AuthMiddleware(tokenMaker *auth.TokenMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "missing authorization header",
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid authorization header",
			})
			return
		}

		claims, err := tokenMaker.VerifyToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid or expired token",
			})
			return
		}

		c.Set(CurrentUserKey, claims)
		c.Next()
	}
}
