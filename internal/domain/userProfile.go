package domain

import "time"

type UserProfile struct {
	ID        int64
	Firstname string
	Lastname  string
	Email     string
	AvatarUrl string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
