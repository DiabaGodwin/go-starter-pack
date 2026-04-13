package app

import (
	"context"
	"real-estate-app/internal/db"
	"real-estate-app/internal/repository"
	"real-estate-app/internal/service/auth"
	handlers "real-estate-app/internal/transport/http/handlers"

	"real-estate-app/internal/config"
	"real-estate-app/internal/service"
	httptransport "real-estate-app/internal/transport/http/routes"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Router *gin.Engine
	DB     *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config, tokenMaker *auth.TokenMaker, store *repository.Store) (*App, error) {
	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	queries := db.New(pool)

	userRepo := repository.NewUserRepository(store)
	userService := service.NewUserService(userRepo, tokenMaker)
	userHandler := handlers.NewUserHandler(userService)

	userProfRepo := repository.NewUserProfileRepository(queries)
	userProfService := service.NewUserProfileService(userProfRepo)
	userProfHandler := handlers.NewUserProfileHandler(userProfService)

	router := httptransport.NewRouter(httptransport.Handlers{
		User:        userHandler,
		UserProfile: userProfHandler,
	}, tokenMaker)

	return &App{
		Router: router,
		DB:     pool,
	}, nil
}
