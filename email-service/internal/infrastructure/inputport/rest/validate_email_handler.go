package rest

import (
	"email-service/internal/application/usecase"
	"encoding/json"
	"net/http"
)

func ValidateEmailHandler(uc usecase.ValidateEmailUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload usecase.AddUserEventRequest
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = uc.Validate(r.Context(), payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)

	}
}
