package service

import (
	"context"
	"errors"
	"real-estate-app/internal/db"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserProfileRepository interface {
	Create(c context.Context, userPrf db.CreateUserProfileParams) (uuid.UUID, error)
	GetAllUserProfiles(c context.Context) ([]db.UserProfile, error)
	GetUserProfileByID(c context.Context, id uuid.UUID) (db.UserProfile, error)
}

type CreateProfileInput struct {
	Firstname   string
	Lastname    string
	AvatarUrl   string
	UserID      uuid.UUID
	FirstName   string
	LastName    string
	DisplayName string
	Phone       string
	Bio         string
	DateOfBirth time.Time
	Gender      string
	Country     string
	City        string
	Address     string
}

type UserProfileService struct {
	UserProfileRepository UserProfileRepository
}

func NewUserProfileService(userProfileRepository UserProfileRepository) *UserProfileService {
	return &UserProfileService{UserProfileRepository: userProfileRepository}
}

func (s *UserProfileService) Create(c *gin.Context, input CreateProfileInput) (uuid.UUID, error) {
	input.Firstname = strings.TrimSpace(input.Firstname)
	input.Lastname = strings.TrimSpace(input.Lastname)
	input.AvatarUrl = strings.TrimSpace(input.AvatarUrl)

	if input.Lastname == "" || input.AvatarUrl == "" {
		return uuid.Nil, errors.New("one or more validation errors occurred, please check and try again")
	}

	profile := db.CreateUserProfileParams{
		FirstName: input.Firstname,
		LastName:  input.Lastname,
	}

	id, err := s.UserProfileRepository.Create(c, profile)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (s *UserProfileService) List(c context.Context) ([]db.UserProfile, error) {
	result, err := s.UserProfileRepository.GetAllUserProfiles(c)
	return result, err
}

func (s *UserProfileService) GetById(c *gin.Context, id uuid.UUID) (*db.UserProfile, error) {
	profile, err := s.UserProfileRepository.GetUserProfileByID(c, id)
	return &profile, err
}
