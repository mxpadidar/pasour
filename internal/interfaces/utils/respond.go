package utils

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}

func RespondError(w http.ResponseWriter, status int, err error) {
	RespondJSON(w, status, map[string]string{"error": err.Error()})
}
