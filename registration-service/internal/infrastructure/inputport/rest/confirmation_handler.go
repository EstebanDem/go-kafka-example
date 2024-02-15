package rest

import (
	"encoding/json"
	"net/http"
	"registration-service/internal/application/usecase"
)

func ConfirmEmailHandler(uc usecase.ConfirmEmailUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload usecase.ValidateEmailRequest
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = uc.Confirm(r.Context(), payload)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
