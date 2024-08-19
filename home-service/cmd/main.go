package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	jetctl "home-service/internal/adapters/broker/jetstream"
	"home-service/internal/adapters/cache/redis"
	"home-service/internal/adapters/db/postgres"
	"home-service/internal/app"
	"home-service/pkg/config"
	"home-service/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig("config", "yaml")
	if err != nil {
		log.Fatal("fail load config: %w", err)
	}

	logger := logger.LogInit(cfg.ModeLog)

	db, err := postgres.New(context.Background(), cfg.PostgresDB, logger.Log)
	if err != nil {
		log.Fatal("db not init: %w", err)
	}

	redisA, err := redis.NewRedisCache(*cfg.RedisDB, *logger.Log)
	if err != nil {
		log.Fatal("redis not init: %w", err)
	}

	jsClient, err := jetctl.NewClient(*cfg.Jet)
	if err != nil {
		log.Fatalf("jetstream not init: %w", err)
	}

	app := app.New(logger.Log, cfg.GRPC, db, jsClient, redisA)

	go app.GRPCServer.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	logger.Log.Info("recieved signal", <-c)
	app.GRPCServer.Stop()
	logger.Log.Info("auth serviec stop")

}
