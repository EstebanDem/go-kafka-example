package service

import (
	"context"
)

type Consumer interface {
	Read(ctx context.Context, chMsg chan ConsumerMessage, chErr chan error, msgTypeConstructor func() ConsumerMessage)
}

type ConsumerMessage interface {
	EventName() string
}
