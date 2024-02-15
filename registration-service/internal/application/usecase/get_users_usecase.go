package usecase

import (
	"registration-service/internal/domain/user"
	"time"
)

type UserResponse struct {
	ID             string
	Name           string
	Email          string
	PhoneNumber    string
	EmailConfirmed bool
	PhoneConfirmed bool
	CreatedAd      time.Time
	UpdatedAt      time.Time
}

type GetAllUsersUseCase interface {
	GetAll() ([]UserResponse, error)
}

type getUserUseCase struct {
	users user.UserRepository
}

func NewGetAllUsersUseCase(repo user.UserRepository) GetAllUsersUseCase {
	return &getUserUseCase{users: repo}
}

func (uc *getUserUseCase) GetAll() ([]UserResponse, error) {
	allUsers, err := uc.users.GetAll()
	if err != nil {
		return nil, err
	}

	var allUsersResponse []UserResponse

	for _, ud := range allUsers {
		allUsersResponse = append(allUsersResponse, mapDomainToResponse(&ud))
	}

	return allUsersResponse, nil
}

func mapDomainToResponse(user *user.User) UserResponse {
	return UserResponse{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		PhoneNumber:    user.PhoneNumber,
		EmailConfirmed: user.EmailConfirmed,
		PhoneConfirmed: user.PhoneConfirmed,
		CreatedAd:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
}
