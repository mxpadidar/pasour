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

// UserRole is the role of the user
type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)
