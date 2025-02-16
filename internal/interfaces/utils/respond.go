package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pasour/internal/domain/errors"
	"pasour/internal/domain/types"
)

func RespondJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}

func WriteErrResponse(err *errors.DomainErr, w http.ResponseWriter) {
	var ErrTypeStatus = map[types.ErrType]int{
		types.ValidationErr: http.StatusBadRequest,
		types.NotFoundErr:   http.StatusNotFound,
		types.ConflictErr:   http.StatusConflict,
		types.InternalErr:   http.StatusInternalServerError,
	}

	if statusCode, ok := ErrTypeStatus[err.Type]; ok {
		RespondJSON(w, statusCode, map[string]string{"error": err.Message})
	} else {
		RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": fmt.Sprintf("unknown error: %s", err.Message)})
	}

}

func RespondError(w http.ResponseWriter, status int, err error) {
	RespondJSON(w, status, map[string]string{"error": err.Error()})
}
