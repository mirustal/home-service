package app

import (
	"log/slog"

	"home-service/internal/adapters/db/postgres"
	"home-service/internal/app/grpc"
	"home-service/internal/services/home"
	"home-service/pkg/config"
)

type App struct {
	GRPCServer *grpc.App
}

func New(log *slog.Logger, cfg *config.GRPCConfig, db *postgres.DbPostgres) *App {
	homeService := home.New(log)
	grpcApp := grpc.New(log, homeService, cfg)

	return &App{
		GRPCServer: grpcApp,
	}
}
