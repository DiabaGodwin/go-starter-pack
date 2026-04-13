package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"real-estate-app/internal/database/migrations"
	"real-estate-app/internal/database/postgres"
	"real-estate-app/internal/repository"
	"real-estate-app/internal/service/auth"
	"syscall"
	"time"

	"real-estate-app/internal/app"
	"real-estate-app/internal/config"
)

func main() {
	cfg := config.MustLoad()

	migrations.Run(cfg.DatabaseURL)

	ctx := context.Background()

	db := postgres.MustConnect(cfg.DatabaseURL)
	defer db.Close()

	store := repository.NewStore(db)
	token := auth.NewTokenMaker(cfg.Issuer, cfg.Secret)

	application, err := app.New(ctx, cfg, token, store)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:         ":" + cfg.HTTPPort,
		Handler:      application.Router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("server running on :%s", cfg.HTTPPort)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}
}
