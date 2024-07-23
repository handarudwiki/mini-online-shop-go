package response

import (
	"errors"
	"net/http"
)

// error general
var (
	ErrNotFound        = errors.New("not found")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbiddenAccess = errors.New("forbidden access")
)

var (
	// auth
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password must have minimum 6 characters")
	ErrUserNotExist          = errors.New("user does not exist")
	ErrEmailAlreadyUsed      = errors.New("email already used")
	ErrPasswordNotMatch      = errors.New("password does not match")
)

type Error struct {
	Message  string
	Code     string
	HttpCode int
}

func NewError(msg string, code string, httpCode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral      = NewError("general error", "99999", http.StatusInternalServerError)
	ErrorBadRequest   = NewError("bad request", "40000", http.StatusBadRequest)
	ErrorNotFound     = NewError(ErrNotFound.Error(), "40400", http.StatusNotFound)
	ErrorUnauthorized = NewError(ErrUnauthorized.Error(), "40100", http.StatusUnauthorized)
)

var (
	// err bad request
	ErrorEmailRequired         = NewError(ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorEmailInvalid          = NewError(ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
	ErrorPasswordRequired      = NewError(ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)
	ErrorEmailAlreadyUsed      = NewError(ErrEmailAlreadyUsed.Error(), "40901", http.StatusConflict)
	ErrorPasswordNotMatch      = NewError(ErrPasswordNotMatch.Error(), "40301", http.StatusUnauthorized)
)

var ()