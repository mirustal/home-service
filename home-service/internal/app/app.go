package app

import (
	"log/slog"
	logs "log"

	"home-service/internal/adapters/db/postgres"
	"home-service/internal/app/grpc"
	authclient "home-service/internal/client/auth-service"
	"home-service/internal/modules/home"
	"home-service/pkg/config"
)

type App struct {
	GRPCServer *grpc.App
}

func New(log *slog.Logger, cfg *config.GRPCConfig, db *postgres.DbPostgres) *App {
	homeService := home.New(log, db, db, db, db, db, db, db)

	authClient, err := authclient.NewAuthClient(cfg.AuthAddress)
	if err != nil {
		logs.Fatalf("failed to create auth client: %v", err)
	}
	
	grpcApp := grpc.New(log, authClient, homeService, cfg)

	return &App{
		GRPCServer: grpcApp,
	}
}
