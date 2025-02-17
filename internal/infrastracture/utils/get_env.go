package utils

import (
	"fmt"
	"os"
	"strconv"
)

type EnvType interface {
	string | int | bool
}

func GetEnv[T EnvType](key string, defaultVal T) (value T, err error) {
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
			return defaultVal, fmt.Errorf("invalid integer: %w", err)
		}
		return any(val).(T), nil
	case bool:
		val, err := strconv.ParseBool(env)
		if err != nil {
			return defaultVal, fmt.Errorf("invalid boolean: %w", err)
		}
		return any(val).(T), nil
	default:
		return defaultVal, fmt.Errorf("invalid type")
	}
}
