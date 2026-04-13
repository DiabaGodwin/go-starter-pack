package repository

import (
	"context"
	"real-estate-app/internal/db"

	"github.com/google/uuid"
)

type UserProfileRepository struct {
	querier db.Querier
}

func NewUserProfileRepository(db db.Querier) *UserProfileRepository {
	return &UserProfileRepository{querier: db}
}

func (r *UserProfileRepository) Create(c context.Context, userPrf db.CreateUserProfileParams) (uuid.UUID, error) {

	user, err := r.querier.CreateUserProfile(c, userPrf)
	if err != nil {
		return uuid.Nil, err
	}
	return user.ID, nil
}

func (r *UserProfileRepository) GetAllUserProfiles(c context.Context) ([]db.UserProfile, error) {
	user, err := r.querier.GetAllUserProfiles(c)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserProfileRepository) GetUserProfileByID(c context.Context, id uuid.UUID) (db.UserProfile, error) {

	data, err := r.querier.GetUserProfileByID(c, id)
	if err != nil {
		return db.UserProfile{}, err
	}
	return data, nil
}
