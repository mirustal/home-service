package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"auth-service/internal/adapters/db/postgres"
	"auth-service/internal/app"
	"auth-service/pkg/config"
	"auth-service/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfigYAML("config", "yaml")
	if err != nil {
		// log.Fatalf("fail load config: %v", err)
	}

	fmt.Println(os.Getenv("secret_key"))

	err = config.LoadENV("local", "env")
	if err != nil {
		// log.Fatalf("fail load config: %v", err)
	}

	logger := logger.LogInit(cfg.ModeLog)

	db, err := postgres.New(context.Background(), cfg.DB, logger.Log)

	app := app.New(logger.Log, cfg.GRPC, db)

	go app.GRPCServer.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	logger.Log.Info("recieved signal", <-c)
	app.GRPCServer.Stop()
	logger.Log.Info("auth serviec stop")

}
