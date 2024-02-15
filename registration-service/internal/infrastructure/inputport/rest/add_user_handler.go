package rest

import (
	"encoding/json"
	"net/http"
	"registration-service/internal/application/usecase"
)

type AddUserRequestJson struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func AddUserHandler(uc usecase.AddUserUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload AddUserRequestJson
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = uc.Add(r.Context(), payload.mapRequest())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}

func (jsonRequest AddUserRequestJson) mapRequest() usecase.AddUserRequest {
	return usecase.AddUserRequest{
		Name:        jsonRequest.Name,
		Email:       jsonRequest.Email,
		PhoneNumber: jsonRequest.PhoneNumber,
	}
}
