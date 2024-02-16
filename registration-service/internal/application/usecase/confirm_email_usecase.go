package usecase

import (
	"context"
	"log"
	"registration-service/internal/domain/user"
	"time"
)

type ValidateEmailRequest struct {
	ID    string
	Valid bool
}

type ConfirmEmailUseCase interface {
	Confirm(ctx context.Context, request ValidateEmailRequest) error
}

func NewConfirmEmailUseCase(repo user.UserRepository) ConfirmEmailUseCase {
	return &confirmEmailUseCase{users: repo}
}

type confirmEmailUseCase struct {
	users user.UserRepository
}

func (c *confirmEmailUseCase) Confirm(_ context.Context, request ValidateEmailRequest) error {
	// dummy sleep to delay action in demo
	time.Sleep(2 * time.Second)

	u, err := c.users.GetByID(request.ID)
	if err != nil {
		return err
	}
	u.EmailConfirmed = true
	err = c.users.Save(*u)
	if err != nil {
		return err
	}

	log.Printf("email confirmed for user: %s", u.ID)
	return nil
}
