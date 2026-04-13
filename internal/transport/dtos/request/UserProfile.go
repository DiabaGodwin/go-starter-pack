package request

import (
	"time"

	"github.com/google/uuid"
)

type UserProfileRequest struct {
	AvatarUrl string `json:"avatar_url"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type UserProfileResponse struct {
	Id        uuid.UUID `json:"id"`
	AvatarUrl string    `json:"avatar_url"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	CreatedAt time.Time `json:"created_at"`
}
