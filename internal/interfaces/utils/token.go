package utils

import (
	"errors"
	"net/http"
	"strings"
)

func ExtractBearerToken(r *http.Request) (error, string) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return errors.New("missing Authorization header"), ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		return errors.New("invalid Authorization header"), ""
	}

	if parts[0] != "Bearer" {
		return errors.New("invalid Authorization header"), ""
	}

	return nil, parts[1]

}
