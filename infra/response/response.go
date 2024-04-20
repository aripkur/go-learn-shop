package response

import "errors"

// error general
var (
	ErrNotFound = errors.New("error not found")
)

var (
	ErrEmailRequired    = errors.New("email is required")
	ErrEmailInvalid     = errors.New("email is invalid")
	ErrPasswordRequired = errors.New("password is required")
	ErrPasswordInvalid  = errors.New("password must have minimum 6 char")
	ErrEmailAlReadyUsed = errors.New("email already used")
	ErrPasswordNotMatch = errors.New("password not match")
)

type Error struct {
	Message string
	Code    string
}

func (e Error) Error() string {
	return e.Message
}

func NewError(message string, code string) Error {
	return Error{
		Message: message,
		Code:    code,
	}
}

var (
	ErrorGeneral = NewError("general error", "99999")
)

var (
	ErrorEmailRequired    = NewError(ErrEmailRequired.Error(), "40001")
	ErrorEmailInvalid     = NewError(ErrEmailInvalid.Error(), "40002")
	ErrorPasswordRequired = NewError(ErrPasswordRequired.Error(), "40003")
	ErrorPasswordInvalid  = NewError(ErrPasswordInvalid.Error(), "40004")
	ErrorEmailAlReadyUsed = NewError(ErrEmailAlReadyUsed.Error(), "40901")
	ErrorPasswordNotMatch = NewError(ErrPasswordNotMatch.Error(), "40101")
)
