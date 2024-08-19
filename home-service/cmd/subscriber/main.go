package main

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"

	"home-service/internal/modules/sender"
)

const (
	natsURL = "nats://localhost:4222"
)

func connectToNATS() (*nats.Conn, error) {
	return nats.Connect(natsURL)
}

func main() {
	nc, err := connectToNATS()
	if err != nil {
		log.Fatalf("Error connecting to NATS: %w", err)
	}
	defer nc.Close()
	send := sender.New()


	if err != nil {
		log.Fatalf("Error init sender: %s", err)
	}

	subject := "house.1.new"
	_, err = nc.Subscribe(subject, func(msg *nats.Msg) {

		send.SendEmail(context.Background(), string(msg.Data), "subscribe update")
	})


	if err != nil {
		log.Fatalf("Error subscribing to subject %s: %w", subject, err)
	}

	select {}
}
