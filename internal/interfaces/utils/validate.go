package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidateReqBody[T any](r *http.Request, payload *T) error {
	err := json.NewDecoder(r.Body).Decode(payload)
	if err == io.EOF {
		return fmt.Errorf("request body is empty")
	} else if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	if err := Validate.Struct(payload); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return fmt.Errorf("validation error: %w", err)
		}

		// Handle validation errors
		var errs []string
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, fmt.Sprintf("%s %s", err.Field(), err.Tag()))
		}
		errStr := strings.Join(errs, "; ")
		return fmt.Errorf("validation error: %s", errStr)
	}

	return nil
}
