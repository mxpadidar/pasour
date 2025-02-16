package utils

import (
	"context"
	"pasour/internal/domain/errors"
	"pasour/internal/domain/types"
)

// ExtractFromContext extracts a value from a context
// and returns it as a pointer to the value.
// Example:
// user, err := utils.ExtractFromContext[UserDTO](r.Context(), types.UserCtxKey)
func ExtractFromContext[T any](ctx context.Context, key types.CtxKey) (*T, error) {
	value, ok := ctx.Value(key).(*T)
	if !ok {
		return nil, errors.NewNotFoundErr("value not found in context")
	}
	return value, nil
}
