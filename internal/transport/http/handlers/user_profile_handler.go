package UserHandlers

import (
	"errors"
	"real-estate-app/internal/db"
	"real-estate-app/internal/service"
	"real-estate-app/internal/transport/dtos/request"
	"real-estate-app/internal/transport/dtos/response/Common"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserProfileHandler struct {
	usrProfileService *service.UserProfileService
}

func NewUserProfileHandler(usrProfileService *service.UserProfileService) *UserProfileHandler {
	return &UserProfileHandler{usrProfileService}
}

func (uh *UserProfileHandler) Create(c *gin.Context) {
	var req request.CreateProfileRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		resp := Common.BadRequest[uuid.UUID]("Invalid request body")
		c.JSON(resp.Code, resp)
		return
	}
	result, err := uh.usrProfileService.Create(c, service.CreateProfileInput{
		Lastname:  req.Lastname,
		AvatarUrl: req.AvatarUrl,
	})
	if err != nil {
		if errors.Is(err, service.ErrValidation) {
			resp := Common.BadRequest[uuid.UUID]("Invalid request body")
			c.JSON(resp.Code, resp)
			return
		}
		resp := Common.InternalServerError[uuid.UUID]("internal server error")
		c.JSON(resp.Code, resp)
		return
	}
	resp := Common.Created[uuid.UUID](result, "Profile created successfully")
	c.JSON(resp.Code, resp)
	return
}

func (uh *UserProfileHandler) GetById(c *gin.Context) {
	var req request.GetUserProfileByIdRequest
	if req.Id == uuid.Nil {
		resp := Common.BadRequest[*db.UserProfile]("Id can not be empty or zero")
		c.JSON(resp.Code, resp)
		return
	}
	result, err := uh.usrProfileService.GetById(c, req.Id)
	if err != nil {
		if result == nil {
			resp := Common.BadRequest[*db.UserProfile]("Record not found")
			c.JSON(resp.Code, resp)
			return
		}

		resp := Common.InternalServerError[*db.UserProfile]("internal server error")
		c.JSON(resp.Code, resp)
		return
	}

	resp := Common.Ok[*db.UserProfile](result, "data found")
	c.JSON(resp.Code, resp)
	return
}
func (uh *UserProfileHandler) ListUserProfiles(c *gin.Context) {

	result, err := uh.usrProfileService.List(c)

	if err != nil {
		resp := Common.BadRequest[*db.UserProfile]("Internal server error")
		c.JSON(resp.Code, resp)
		return
	}
	resp := Common.Ok(result)
	c.JSON(resp.Code, resp)
	return
}
