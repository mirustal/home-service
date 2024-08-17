package app

import (
	"log/slog"

	"auth-service/internal/adapters/db/postgres"
	"auth-service/internal/app/grpc/auth"
	"auth-service/internal/services/auth"
	"auth-service/pkg/config"
)

type App struct {
	GRPCServer *grpc.App
}

func New(log *slog.Logger, cfg *config.GRPCConfig, db *postgres.DbPostgres) *App {
	authService := auth.New(log, db, db)
	grpcApp := grpc.New(log, authService, cfg)

	return &App{
		GRPCServer: grpcApp,
	}
}
