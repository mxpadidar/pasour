package types

// ErrType is the type of domain errors
type ErrType int

const (
	ValidationErr ErrType = iota
	NotFoundErr
	ConflictErr
	UnAuthorizedErr
	InternalErr
)

// UserCtxKey is the key used to store the user in the context
type CtxKey string

const UserCtxKey CtxKey = "user"
