package UserHandlers

import (
	"errors"
	"real-estate-app/internal/middleware"
	"real-estate-app/internal/repository"
	"real-estate-app/internal/service"
	"real-estate-app/internal/transport/dtos/request"
	"real-estate-app/internal/transport/dtos/response"
	"real-estate-app/internal/transport/dtos/response/Common"
	"real-estate-app/internal/validator"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req request.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		resp := Common.BadRequest[uuid.UUID]()
		c.JSON(resp.Code, resp)
		return
	}

	// Validate
	if errs, err := validator.ValidateStruct(req); err != nil {
		c.JSON(400, gin.H{
			"message": "validation failed",
			"errors":  errs,
		})
		return
	}

	result, err := h.userService.RegisterUser(c.Request.Context(), repository.RegisterUserInput{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      req.Role,
	})

	if err != nil {
		if errors.Is(err, service.ErrEmailExists) {
			resp := Common.BadRequest[service.RegisterUserOutput]("email already exists")
			c.JSON(resp.Code, resp)
			return
		}

		if errors.Is(err, service.ErrValidation) {
			resp := Common.BadRequest[service.RegisterUserOutput]()
			c.JSON(resp.Code, resp)
			return
		}

		resp := Common.InternalServerError[service.RegisterUserOutput]()
		c.JSON(resp.Code, resp)
		return
	}

	resp := Common.Created[service.RegisterUserOutput](result)
	c.JSON(resp.Code, resp)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	var req request.ListUserRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		resp := Common.BadRequest[[]response.ListUserResponse]()
		c.JSON(resp.Code, resp)
	}

	users, err := h.userService.List(c.Request.Context(), service.ListUserInput{
		Limit:  req.Limit,
		Offset: req.Offset,
	})

	out := make([]response.ListUserResponse, 0, len(users))

	for _, u := range users {
		out = append(out, response.ListUserResponse{
			Id:        u.ID,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
		})
	}

	if err != nil {
		resp := Common.InternalServerError[[]response.ListUserResponse]()
		c.JSON(resp.Code, out)
		return
	}

	resp := Common.Ok(users)
	c.JSON(resp.Code, resp)
	return
}

func (h *UserHandler) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		resp := Common.BadRequest[response.LoginResponse]()
		c.JSON(resp.Code, resp)
		return
	}

	result, err := h.userService.Login(c.Request.Context(), service.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			resp := Common.BadRequest[response.LoginResponse]()
			c.JSON(resp.Code, resp)
			return
		}
		resp := Common.InternalServerError[response.LoginResponse]("login failed")
		c.JSON(resp.Code, resp)
		return
	}
	res := response.LoginResponse{
		AccessToken: result.AccessToken,
		UserID:      result.UserID,
		Email:       result.Email,
		Role:        result.Role,
	}
	resp := Common.Ok[response.LoginResponse](res, "Login successful")
	c.JSON(resp.Code, resp)
	return
}

func (h *UserHandler) Me(c *gin.Context) {
	user, err := middleware.CurrentUser(c)
	if err != nil {
		c.JSON(401, gin.H{"message": "unauthorized"})
		return
	}

	c.JSON(200, gin.H{
		"userId": user.UserID,
		"email":  user.Email,
		"role":   user.Role,
	})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	_, err := middleware.CurrentUser(c)
	if err != nil {
		c.JSON(401, gin.H{"message": "unauthorized"})
		return
	}

	idParam := c.Query("id")
	if idParam == "" {
		resp := Common.BadRequest[response.GetUserResponse]("id query parameter is required")
		c.JSON(resp.Code, resp)
		return
	}

	id, err := uuid.Parse(idParam)
	if err != nil {
		resp := Common.BadRequest[response.GetUserResponse]("Invalid uuid format")
		c.JSON(resp.Code, resp)
		return
	}

	result, err := h.userService.GetUserByID(c.Request.Context(), id)
	if err != nil {
		resp := Common.InternalServerError[response.GetUserResponse]("Failed fetching user")
		c.JSON(resp.Code, resp)
		return
	}

	res := response.GetUserResponse{
		ID:              result.ID,
		Email:           result.Email,
		Role:            result.Role,
		Status:          result.Status,
		PasswordHash:    result.PasswordHash,
		EmailVerifiedAt: result.EmailVerifiedAt,
		LastLoginAt:     result.LastLoginAt,
		CreatedAt:       result.CreatedAt,
	}

	resp := Common.Ok[response.GetUserResponse](res, "fetching user successful")
	c.JSON(resp.Code, resp)
	return
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {

	_, err := middleware.CurrentUser(c)
	if err != nil {
		resp := Common.Unauthorized[response.GetUserResponse]()
		c.JSON(resp.Code, resp)
		return
	}

	emailParam := c.Query("email")
	if emailParam == "" {
		resp := Common.BadRequest[response.GetUserResponse]("email query parameter is required")
		c.JSON(resp.Code, resp)
		return
	}

	result, err := h.userService.GetUserByEmail(c.Request.Context(), emailParam)
	if err != nil {
		resp := Common.InternalServerError[response.GetUserResponse]("Failed fetching user")
		c.JSON(resp.Code, resp)
		return
	}
	res := response.GetUserResponse{
		ID:              result.ID,
		Email:           result.Email,
		Role:            result.Role,
		Status:          result.Status,
		PasswordHash:    result.PasswordHash,
		EmailVerifiedAt: result.EmailVerifiedAt,
		LastLoginAt:     result.LastLoginAt,
		CreatedAt:       result.CreatedAt,
	}

	resp := Common.Ok[response.GetUserResponse](res, "fetching user successful")
	c.JSON(resp.Code, resp)
	return
}
