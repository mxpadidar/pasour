package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pasour/internal/domain/errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidateReqBody[T any](r *http.Request, payload *T) *errors.DomainErr {
	err := json.NewDecoder(r.Body).Decode(payload)
	if err == io.EOF {
		return errors.NewValidationErr("request body is empty")
	} else if err != nil {
		return errors.NewValidationErr(err)
	}

	if err := Validate.Struct(payload); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errors.NewValidationErr(err)
		}

		// Handle validation errors
		var errs []string
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, fmt.Sprintf("%s %s", err.Field(), err.Tag()))
		}
		errStr := strings.Join(errs, "; ")
		return errors.NewValidationErr(errStr)
	}

	return nil
}
