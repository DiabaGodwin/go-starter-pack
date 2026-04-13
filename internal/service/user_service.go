package service

import (
	"context"
	"errors"
	"real-estate-app/internal/db"
	"real-estate-app/internal/repository"
	"real-estate-app/internal/service/auth"
	"strings"
	"time"

	"github.com/google/uuid"
)

type UserRepository interface {
	RegisterUser(c context.Context, user repository.RegisterUserInput) (db.CreateUserRow, error)
	List(c context.Context, params db.GetUsersParams) ([]db.GetUsersRow, error)
	GetByID(c context.Context, id uuid.UUID) (db.GetUserByIDRow, error)
	GetUserByEmail(c context.Context, email string) (db.GetUserByEmailRow, error)
	GetUserWithProfile(c context.Context, id uuid.UUID) (db.GetUserWithProfileRow, error)
	CheckUniqueEmail(c context.Context, email string) (bool, error)
}

type ListUserInput struct {
	Limit  int32
	Offset int32
}

type ListUserOutput struct {
	ID        uuid.UUID
	Email     string
	CreatedAt time.Time
}

type LoginInput struct {
	Email    string
	Password string
}
type LoginOutput struct {
	AccessToken string
	UserID      uuid.UUID
	Email       string
	Role        string
}

type RegisterUserOutput struct {
	AccessToken string
	UserID      uuid.UUID
	Email       string
	Role        string
}

type UserService struct {
	repo       UserRepository
	tokenMaker *auth.TokenMaker
}

type GetUserByIdOutput struct {
	ID              uuid.UUID
	Email           string
	Role            string
	Status          string
	PasswordHash    string
	EmailVerifiedAt time.Time
	LastLoginAt     time.Time
	CreatedAt       time.Time
}

type GetUserByEmailOutput struct {
	ID              uuid.UUID
	Email           string
	Role            string
	Status          string
	PasswordHash    string
	EmailVerifiedAt time.Time
	LastLoginAt     time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type GetUserWithProfileOutput struct {
	ID          uuid.UUID
	Email       string
	Role        string
	Status      string
	CreatedAt   time.Time
	FirstName   string
	LastName    string
	DisplayName string
	AvatarUrl   string
}

func NewUserService(repo UserRepository, tokenMaker *auth.TokenMaker) *UserService {
	return &UserService{
		repo:       repo,
		tokenMaker: tokenMaker,
	}
}

func (s *UserService) Login(ctx context.Context, input LoginInput) (LoginOutput, error) {
	input.Email = strings.TrimSpace(input.Email)
	input.Password = strings.TrimSpace(input.Password)

	if input.Email == "" || input.Password == "" {
		return LoginOutput{}, ErrValidation
	}

	user, err := s.repo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return LoginOutput{}, err
	}

	_, err = auth.CheckPassword(input.Password, user.PasswordHash)
	if err != nil {
		return LoginOutput{}, ErrInvalidCredentials
	}

	token, err := s.tokenMaker.CreateToken(user.ID.String(), user.Email, user.Role)
	if err != nil {
		return LoginOutput{}, err
	}

	return LoginOutput{
		AccessToken: token,
		UserID:      user.ID,
		Email:       user.Email,
		Role:        user.Role,
	}, nil

}

func (s *UserService) RegisterUser(ctx context.Context, input repository.RegisterUserInput) (RegisterUserOutput, error) {
	input.Email = strings.TrimSpace(input.Email)
	input.Password = strings.TrimSpace(input.Password)

	if input.Email == "" || input.Password == "" {
		return RegisterUserOutput{}, ErrValidation
	}
	//check unique email
	result, _ := s.repo.CheckUniqueEmail(ctx, input.Email)
	if result {
		return RegisterUserOutput{}, errors.New("the email you entered exist for another record")
	}

	hashPass, err := auth.HashPassword(input.Password)
	if err != nil {
		return RegisterUserOutput{}, err
	}

	res, err := s.repo.RegisterUser(ctx, repository.RegisterUserInput{
		Email:     input.Email,
		Password:  hashPass,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Role:      input.Role,
	})

	if err != nil {
		return RegisterUserOutput{}, err
	}

	token, err := s.tokenMaker.CreateToken(res.ID.String(), res.Email, res.Role)

	return RegisterUserOutput{
		AccessToken: token,
		Email:       res.Email,
		Role:        res.Role,
		UserID:      res.ID,
	}, nil
}

func (s *UserService) List(c context.Context, params ListUserInput) ([]ListUserOutput, error) {
	res, err := s.repo.List(c, db.GetUsersParams{
		Limit:  params.Limit,
		Offset: params.Offset,
	})
	if err != nil {
		return nil, err
	}

	out := make([]ListUserOutput, 0, len(res))

	for _, u := range res {
		out = append(out, ListUserOutput{
			ID:        u.ID,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
		})
	}

	return out, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (GetUserByIdOutput, error) {
	res, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return GetUserByIdOutput{}, err
	}
	return GetUserByIdOutput{
		ID:              res.ID,
		Email:           res.Email,
		Role:            res.Role,
		Status:          res.Status,
		PasswordHash:    res.PasswordHash,
		EmailVerifiedAt: res.EmailVerifiedAt.Time,
		LastLoginAt:     res.LastLoginAt.Time,
		CreatedAt:       res.CreatedAt,
	}, nil

}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailOutput, error) {
	res, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return GetUserByEmailOutput{}, err
	}
	return GetUserByEmailOutput{
		ID:              res.ID,
		Email:           res.Email,
		Role:            res.Role,
		Status:          res.Status,
		PasswordHash:    res.PasswordHash,
		EmailVerifiedAt: res.EmailVerifiedAt.Time,
		LastLoginAt:     res.LastLoginAt.Time,
		CreatedAt:       res.CreatedAt,
	}, nil

}

func (s *UserService) GetUserWithProfile(ctx context.Context, id uuid.UUID) (GetUserWithProfileOutput, error) {
	res, err := s.repo.GetUserWithProfile(ctx, id)
	if err != nil {
		return GetUserWithProfileOutput{}, err
	}

	return GetUserWithProfileOutput{
		ID:          res.ID,
		Email:       res.Email,
		Role:        res.Role,
		Status:      res.Status,
		CreatedAt:   res.CreatedAt,
		FirstName:   res.FirstName.String,
		LastName:    res.LastName.String,
		DisplayName: res.DisplayName.String,
		AvatarUrl:   res.AvatarUrl,
	}, nil
}
