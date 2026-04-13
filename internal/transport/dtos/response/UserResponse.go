package response

import (
	"time"

	"github.com/google/uuid"
)

type ListUserResponse struct {
	Id        uuid.UUID
	Email     string
	CreatedAt time.Time
}
type LoginResponse struct {
	AccessToken string    `json:"accessToken"`
	UserID      uuid.UUID `json:"userId"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
}

type GetUserResponse struct {
	ID              uuid.UUID `json:"id"`
	Email           string    `json:"email"`
	Role            string    `json:"role"`
	Status          string    `json:"status"`
	PasswordHash    string    `json:"passwordHash"`
	EmailVerifiedAt time.Time `json:"emailVerifiedAt"`
	LastLoginAt     time.Time `json:"lastLoginAt"`
	CreatedAt       time.Time `json:"createdAt"`
}
