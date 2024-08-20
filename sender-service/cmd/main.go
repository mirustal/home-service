package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"

	"sender-service/internal/adapters/broker"
	"sender-service/internal/modules/sender"
	"sender-service/pkg/config"
)

func main() {
	time.Sleep(3 * time.Second)

	cfg, err := config.LoadConfig("config", "yaml")
	if err != nil {
		log.Fatalf("fail load config: %v", err)
	}

	nc, err := broker.New(cfg.Jet)
	if err != nil {
		log.Fatalf("fail connect nats: %v", err)
	}

	send := sender.New()

	defer nc.Close()

	if err != nil {
		log.Fatalf("Error init sender: %v", err)
	}

	subject := cfg.Jet.Subject
	_, err = nc.Con.Subscribe(subject, func(msg *nats.Msg) {
		err := send.SendEmail(context.Background(), string(msg.Data), "subscribe update")
		if err != nil {
			log.Printf("Failed to send email: %v", err)
		} else {
			log.Printf("Successfully sent email to '%s'", string(msg.Data))
		}
	})
	

	if err != nil {
		log.Fatalf("Error subscribing to subject %s: %v", subject, err)
	}

	log.Println("sender-service start")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)


	<-c

	log.Println("sender-service shutdown")
}
