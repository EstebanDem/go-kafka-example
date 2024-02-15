package usecase

import (
	"context"
	"registration-service/internal/domain/user"
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
	u, err := c.users.GetByID(request.ID)
	if err != nil {
		return err
	}
	u.EmailConfirmed = true
	err = c.users.Save(*u)
	if err != nil {
		return err
	}
	return nil
}
