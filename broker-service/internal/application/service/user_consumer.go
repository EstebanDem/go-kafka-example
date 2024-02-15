package service

import (
	"context"
)

type Consumer interface {
	Read(ctx context.Context, chMsg chan NewUserMessage, chErr chan error)
}

type ConsumerMessage interface {
	EventName() string
}

type NewUserMessage struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
