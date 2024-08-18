package sender

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Sender interface {
	SendEmail(ctx context.Context, recipient string, message string) error
}

type EmailSender struct{}

func NewEmailSender() *EmailSender {
	return &EmailSender{}
}

func (s *EmailSender) SendEmail(ctx context.Context, recipient string, message string) error {

	duration := time.Duration(rand.Int63n(3000)) * time.Millisecond
	time.Sleep(duration)

	errorProbability := 0.1
	if rand.Float64() < errorProbability {
		return errors.New("internal error")
	}

	fmt.Printf("send message '%s' to '%s'\n", message, recipient)

	return nil
}
