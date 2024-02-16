package rest

import (
	"encoding/json"
	"net/http"
	"registration-service/internal/application/usecase"
	"time"
)

type UserResponseJson struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	PhoneNumber    string    `json:"phone_number"`
	EmailConfirmed bool      `json:"email_confirmed"`
	PhoneConfirmed bool      `json:"phone_confirmed"`
	CreatedAd      time.Time `json:"created_ad"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func GetAllUsersHandler(uc usecase.GetAllUsersUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := uc.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var payload []UserResponseJson
		for _, u := range users {
			payload = append(payload, mapToJsonResponse(u))
		}

		out, err := json.Marshal(payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(out)

	}
}

func mapToJsonResponse(ur usecase.UserResponse) UserResponseJson {
	return UserResponseJson{
		ID:             ur.ID,
		Name:           ur.Name,
		Email:          ur.Email,
		PhoneNumber:    ur.PhoneNumber,
		EmailConfirmed: ur.EmailConfirmed,
		PhoneConfirmed: ur.PhoneConfirmed,
		CreatedAd:      ur.CreatedAd,
		UpdatedAt:      ur.UpdatedAt,
	}
}
