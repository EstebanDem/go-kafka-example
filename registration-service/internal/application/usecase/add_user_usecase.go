package usecase

import (
	"registration-service/internal/domain/user"
	"registration-service/internal/pkg"
	"time"
)

type AddUserRequest struct {
	Name        string
	Email       string
	PhoneNumber string
}

type AddUserUseCase interface {
	Add(payload AddUserRequest) error
}

type addUserUseCase struct {
	users user.UserRepository
}

func NewAddUserUseCase(repo user.UserRepository) AddUserUseCase {
	return &addUserUseCase{
		users: repo,
	}
}

func (uc *addUserUseCase) Add(payload AddUserRequest) error {
	newUser := user.User{
		ID:             pkg.NewUUID().String(),
		Name:           payload.Name,
		Email:          payload.Email,
		PhoneNumber:    payload.PhoneNumber,
		EmailConfirmed: false,
		PhoneConfirmed: false,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	err := uc.users.Save(newUser)
	if err != nil {
		return err
	}

	return nil
}
