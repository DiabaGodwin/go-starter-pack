package routes

import (
	"real-estate-app/internal/middleware"
	"real-estate-app/internal/service/auth"
	handlers "real-estate-app/internal/transport/http/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, h *handlers.UserHandler, tokenMaker *auth.TokenMaker) {
	users := r.Group("/api/v1/auth")
	{
		users.POST("/register", h.CreateUser)
		users.GET("/list-users", h.ListUsers)
		users.POST("/login", h.Login)
		users.GET("/get-user-by-id", h.GetUserByID)
		users.GET("/get-user-by-email", h.GetUserByEmail)
		protected := users.Group("/")
		protected.Use(middleware.AuthMiddleware(tokenMaker))
		{
			protected.GET("/me", h.Me)
		}
	}

}
