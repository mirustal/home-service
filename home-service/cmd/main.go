package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"home-service/internal/adapters/cache/redis"
	"home-service/internal/adapters/db/postgres"
	"home-service/internal/app"
	"home-service/pkg/config"
	"home-service/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig("config", "yaml")
	if err != nil {
		log.Fatal("fail load config: %v", err)
	}

	logger := logger.LogInit(cfg.ModeLog)

	db, err := postgres.New(context.Background(), cfg.PostgresDB, logger.Log)
	if err != nil {
		logger.Log.Warn ("fail load config: %v", err)
	}
	_, err = redis.NewRedisCache(*cfg.RedisDB, *logger.Log)
	

	app := app.New(logger.Log, cfg.GRPC, db)

	go app.GRPCServer.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	logger.Log.Info("recieved signal", <-c)
	app.GRPCServer.Stop()
	logger.Log.Info("auth serviec stop")

}
