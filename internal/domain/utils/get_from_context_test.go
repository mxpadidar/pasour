package utils

import (
	"context"
	"pasour/internal/domain/types"
	"reflect"
	"testing"
)

type TestStruct struct {
	Name string
}

// ExtractFromContext extracts a value from a context
func TestExtractFromContext(t *testing.T) {
	ctx := context.Background()
	key := types.CtxKey("test")
	value := &TestStruct{Name: "test"}
	ctx = context.WithValue(ctx, key, value)

	result, err := ExtractFromContext[TestStruct](ctx, key)
	if err != nil {
		t.Errorf("got error: %v", err)
	}

	// Check type using reflect
	if reflect.TypeOf(result) != reflect.TypeOf(value) {
		t.Errorf("expected type *TestStruct, got %T", result)
	}

	if result.Name != value.Name {
		t.Errorf("expected %s, got %s", value.Name, result.Name)
	}
}
