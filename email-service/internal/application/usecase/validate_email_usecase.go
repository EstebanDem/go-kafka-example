package usecase

import (
	"context"
	"email-service/internal/application/service"
	"time"
)

type AddUserEventRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type ValidateEmailResponse struct {
	ID    string `json:"id"`
	Valid bool   `json:"valid"`
}

type ValidateEmailUseCase interface {
	Validate(ctx context.Context, eventRequest AddUserEventRequest) error
}

type validateEmailUseCase struct {
	producer service.Publisher
}

func NewEmailValidator(producer service.Publisher) ValidateEmailUseCase {
	return &validateEmailUseCase{producer: producer}
}

func (uc *validateEmailUseCase) Validate(ctx context.Context, eventRequest AddUserEventRequest) error {
	// there could be a lot more logic, but let's keep this simple for the example
	time.Sleep(10 * time.Second)
	err := uc.producer.Publish(ctx, ValidateEmailResponse{ID: eventRequest.ID, Valid: true})
	if err != nil {
		return err
	}
	return nil
}
