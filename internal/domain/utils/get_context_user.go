package utils

import (
	"context"
	"errors"
	"pasour/internal/domain/types"
)

func ExtractFromContext[T any](ctx context.Context, key types.CtxKey) (*T, error) {
	value, ok := ctx.Value(key).(*T)
	if !ok {
		return nil, errors.New("user is not exists")
	}
	return value, nil
}
