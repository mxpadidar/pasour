package errors

type ErrType int

const (
	ValidationErr ErrType = iota
	NotFoundErr
	ConflictErr
	InternalErr
)

type DomainErr struct {
	Type    ErrType
	Message string
}

func (e *DomainErr) Error() string {
	return e.Message
}

func NewNotFoundErr(errMsg interface{}) *DomainErr {
	return &DomainErr{
		Type:    NotFoundErr,
		Message: errMsgToString(errMsg),
	}
}

func NewConflictErr(errMsg interface{}) *DomainErr {
	return &DomainErr{
		Type:    ConflictErr,
		Message: errMsgToString(errMsg),
	}
}

func NewValidationErr(errMsg interface{}) *DomainErr {
	return &DomainErr{
		Type:    ValidationErr,
		Message: errMsgToString(errMsg),
	}
}

func NewInternalErr(errMsg interface{}) *DomainErr {
	return &DomainErr{
		Type:    InternalErr,
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
