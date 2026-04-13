package routes

import (
	handler "real-estate-app/internal/transport/http/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterUserProfileRoute(router *gin.Engine, handler *handler.UserProfileHandler) {

	profile := router.Group("api/v1/profile")
	{
		profile.POST("/create", handler.Create)
		//profile.GET("/get-all-profiles", handler.GetById)
		profile.GET("/all", handler.ListUserProfiles)
	}
}
