package app

import (
	"log/slog"

	"auth-service/internal/adapters/db/postgres"
	grpcapp "auth-service/internal/app/grpc"
	"auth-service/internal/modules/auth"
	"auth-service/pkg/config"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, cfg *config.GRPCConfig, db *postgres.DbPostgres) *App {
	authService := auth.New(log, db, db, db)
	grpcApp := grpcapp.New(log, authService, cfg)

	return &App{
		GRPCServer: grpcApp,
	}
}
