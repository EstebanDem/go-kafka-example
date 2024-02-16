package rest

import (
	"email-service/internal/application/usecase"
	"encoding/json"
	"net/http"
)

type AddUserEventRequestJson struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func ValidateEmailHandler(uc usecase.ValidateEmailUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload AddUserEventRequestJson
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = uc.Validate(r.Context(), payload.mapToUseCaseRequest())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)

	}
}

func (p AddUserEventRequestJson) mapToUseCaseRequest() usecase.AddUserEventRequest {
	return usecase.AddUserEventRequest{
		ID:          p.ID,
		Name:        p.Name,
		Email:       p.Email,
		PhoneNumber: p.PhoneNumber,
	}
}
