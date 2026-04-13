package routes

import (
	"real-estate-app/internal/service/auth"
	handlers "real-estate-app/internal/transport/http/handlers"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	User        *handlers.UserHandler
	UserProfile *handlers.UserProfileHandler
}

func NewRouter(h Handlers, tokenMaker *auth.TokenMaker) *gin.Engine {
	r := gin.Default()

	RegisterUserRoutes(r, h.User, tokenMaker)
	RegisterUserProfileRoute(r, h.UserProfile)

	return r
}
