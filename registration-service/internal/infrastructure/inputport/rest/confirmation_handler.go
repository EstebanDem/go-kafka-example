package rest

import (
	"encoding/json"
	"net/http"
	"registration-service/internal/application/usecase"
)

type ValidateEmailRequestJson struct {
	ID    string `json:"id"`
	Valid bool   `json:"valid"`
}

func ConfirmEmailHandler(uc usecase.ConfirmEmailUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload ValidateEmailRequestJson
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = uc.Confirm(r.Context(), payload.mapToRequest())
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (p ValidateEmailRequestJson) mapToRequest() usecase.ValidateEmailRequest {
	return usecase.ValidateEmailRequest{
		ID:    p.ID,
		Valid: p.Valid,
	}
}
