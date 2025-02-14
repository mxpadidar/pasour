package utils

import (
	"os"
	"pasour/internal/domain/errors"
	"strconv"
)

type EnvType interface {
	string | int | bool
}

func GetEnv[T EnvType](key string, defaultVal T) (value T, err *errors.DomainErr) {
	env, exists := os.LookupEnv(key)
	if !exists {
		return defaultVal, nil
	}

	switch any(defaultVal).(type) {
	case string:
		return any(value).(T), nil
	case int:
		val, err := strconv.Atoi(env)
		if err != nil {
			return defaultVal, errors.NewValidationErr(err)
		}
		return any(val).(T), nil
	case bool:
		val, err := strconv.ParseBool(env)
		if err != nil {
			return defaultVal, errors.NewValidationErr(err)
		}
		return any(val).(T), nil
	default:
		return defaultVal, errors.NewValidationErr("Invalid type")
	}
}
