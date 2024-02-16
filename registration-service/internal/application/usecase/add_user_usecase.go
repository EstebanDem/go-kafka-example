package usecase

import (
	"context"
	"log"
	"registration-service/internal/application/service"
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
	Add(ctx context.Context, payload AddUserRequest) error
}

type addUserUseCase struct {
	users    user.UserRepository
	producer service.Publisher
}

func NewAddUserUseCase(repo user.UserRepository, producer service.Publisher) AddUserUseCase {
	return &addUserUseCase{
		users:    repo,
		producer: producer,
	}
}

func (uc *addUserUseCase) Add(ctx context.Context, payload AddUserRequest) error {
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
	log.Printf("user with id: %s saved", newUser.ID)

	event := AddUserEvent{
		ID:          newUser.ID,
		Name:        newUser.Name,
		Email:       newUser.Email,
		PhoneNumber: newUser.PhoneNumber,
	}

	err = uc.producer.Publish(ctx, event)
	if err != nil {
		return err
	}

	return nil
}

type AddUserEvent struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}
