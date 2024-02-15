package service

import (
	"context"
)

type ConsumerEmail interface {
	Read(ctx context.Context, chMsg chan EmailMessage, chErr chan error)
}

type EmailMessage struct {
	ID    string `json:"id"`
	Valid bool   `json:"valid"`
}
