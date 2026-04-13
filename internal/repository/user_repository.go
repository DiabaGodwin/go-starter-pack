package repository

import (
	"context"
	"real-estate-app/internal/db"

	"github.com/google/uuid"
)

type RegisterUserInput struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
	Role      string
}
type UserRepository struct {
	s *Store
}

func NewUserRepository(db *Store) *UserRepository {
	return &UserRepository{s: db}
}

func (r *UserRepository) RegisterUser(ctx context.Context, req RegisterUserInput) (db.CreateUserRow, error) {
	var result db.CreateUserRow

	err := r.s.InTx(ctx, func(q *db.Queries) error {

		user, err := q.CreateUser(ctx, db.CreateUserParams{
			PasswordHash: req.Password,
			Role:         req.Role,
			Email:        req.Email,
			Status:       "active",
		})
		if err != nil {
			return err
		}

		_, err = q.CreateUserProfile(ctx, db.CreateUserProfileParams{
			UserID:      user.ID,
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			DisplayName: req.FirstName,
		})
		if err != nil {
			return err
		}

		result = user
		return nil
	})

	return result, err
}
func (r *UserRepository) List(c context.Context, params db.GetUsersParams) ([]db.GetUsersRow, error) {
	data, err := r.s.queries.GetUsers(c, params)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *UserRepository) GetByID(c context.Context, id uuid.UUID) (db.GetUserByIDRow, error) {
	data, err := r.s.queries.GetUserByID(c, id)
	if err != nil {
		return db.GetUserByIDRow{}, err
	}
	return data, nil
}

func (r *UserRepository) GetUserByEmail(c context.Context, email string) (db.GetUserByEmailRow, error) {
	data, err := r.s.queries.GetUserByEmail(c, email)
	if err != nil {
		return db.GetUserByEmailRow{}, err
	}
	return data, nil
}
func (r *UserRepository) GetUserWithProfile(c context.Context, id uuid.UUID) (db.GetUserWithProfileRow, error) {

	data, err := r.s.queries.GetUserWithProfile(c, id)
	if err != nil {
		return db.GetUserWithProfileRow{}, err
	}
	return data, nil
}

func (r *UserRepository) CheckUniqueEmail(c context.Context, email string) (bool, error) {
	res, err := r.s.queries.EmailExists(c, email)
	if err != nil {
		return false, err
	}
	return res != db.EmailExistsRow{}, nil
}
