package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"

	"subscriber/internal/adapters/broker"
	"subscriber/internal/modules/sender"
	"subscriber/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig("config", "yaml")
	if err != nil {
		log.Fatalf("fail load config: %w", err)
	}

	nc, err := broker.New(cfg.Jet)
	if err != nil {
		log.Fatalf("fail connect nats: %w", err)
	}
	defer nc.Close()
	
	send := sender.New()
	if err != nil {
		log.Fatalf("Error init sender: %s", err)
	}

	subject := "house.1.new"
	_, err = nc.Con.Subscribe(subject, func(msg *nats.Msg) {

		send.SendEmail(context.Background(), string(msg.Data), "subscribe update")
	})

	if err != nil {
		log.Fatalf("Error subscribing to subject %s: %w", subject, err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
}
