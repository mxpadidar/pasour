package errors

import "pasour/internal/domain/types"

type DomainErr struct {
	Type    types.ErrType
	Message string
}

func (e *DomainErr) Error() string {
	return e.Message
}

func NewNotFoundErr(errMsg interface{}) *DomainErr {
	return &DomainErr{
		Type:    types.NotFoundErr,
		Message: errMsgToString(errMsg),
	}
}

func NewConflictErr(errMsg interface{}) *DomainErr {
	return &DomainErr{
		Type:    types.ConflictErr,
		Message: errMsgToString(errMsg),
	}
}

func NewUnAuthorizedErr(errMsg interface{}) *DomainErr {
	return &DomainErr{
		Type:    types.UnAuthorizedErr,
		Message: errMsgToString(errMsg),
	}
}

func NewValidationErr(errMsg interface{}) *DomainErr {
	return &DomainErr{
		Type:    types.ValidationErr,
		Message: errMsgToString(errMsg),
	}
}

func NewInternalErr(errMsg interface{}) *DomainErr {
	return &DomainErr{
		Type:    types.InternalErr,
		Message: errMsgToString(errMsg),
	}
}

func errMsgToString(errMsg interface{}) string {
	switch errMsg.(type) {
	case string:
		return errMsg.(string)
	case error:
		return errMsg.(error).Error()
	default:
		return "unknown error"
	}
}
