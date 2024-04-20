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
