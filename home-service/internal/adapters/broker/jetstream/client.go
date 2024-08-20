package jetstream

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"

	"home-service/pkg/config"
)

type JetStreamClient interface {
	Publish(subject string, message []byte) error
	Subscribe(subject string, handler nats.MsgHandler) (*nats.Subscription, error)
}

type Client struct {
	Js nats.JetStreamContext
	cfg config.JetConfig
}

func NewClient(cfg config.JetConfig) (*Client, error) {
	nc, err := nats.Connect(cfg.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, fmt.Errorf("failed to get JetStream context: %w", err)
	}
	createStream(context.Background(), js, cfg)
	return &Client{Js: js}, nil
}

func (c *Client) Publish(subject string, message []byte) error {
	_, err := c.Js.Publish(subject, message)
	return err
}

func (c *Client) Subscribe(subject string, handler nats.MsgHandler) (*nats.Subscription, error) {
	return c.Js.Subscribe(subject, handler)
}

func createStream(ctx context.Context, jsCtx nats.JetStreamContext, cfg config.JetConfig) (*nats.StreamInfo, error) {
	subjects := []string{cfg.Subject}
	stream, err := jsCtx.AddStream(&nats.StreamConfig{
		Name:              cfg.Name,
		Subjects:          subjects,
		Retention:         nats.InterestPolicy, 
		Discard:           nats.DiscardOld,     
		MaxAge:            7 * 24 * time.Hour,  
		Storage:           nats.FileStorage,    
		MaxMsgsPerSubject: 100_000_000,         
		MaxMsgSize:        4 << 20,             
		NoAck:             false,               
	}, nats.Context(ctx))
	if err != nil {
		return nil, fmt.Errorf("add stream: %w", err)
	}

	return stream, nil
}
