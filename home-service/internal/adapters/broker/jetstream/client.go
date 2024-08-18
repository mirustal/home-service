package jetstream

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type JetStreamClient interface {
	Publish(subject string, message []byte) error
	Subscribe(subject string, handler nats.MsgHandler) (*nats.Subscription, error)
}

type Client struct {
	js nats.JetStreamContext
}

func NewClient(natsURL string) (*Client, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS: %w", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, fmt.Errorf("failed to get JetStream context: %w", err)
	}

	return &Client{js: js}, nil
}


func (c *Client) Publish(subject string, message []byte) error {
	_, err := c.js.Publish(subject, message)
	return err
}

func (c *Client) Subscribe(subject string, handler nats.MsgHandler) (*nats.Subscription, error) {
	return c.js.Subscribe(subject, handler)
}
