package global

import (
	"errors"
	"net/http"
)

var (
	ErrNoData                 = errors.New("data not found")
	ErrInvalidParam           = errors.New("invalid request params")
	ErrInternalError          = errors.New("internal server error")
	ErrInvalidLoginOrPassword = errors.New("invalid login or password")
)

var StatusCodes = map[error]int{
	ErrNoData:                 http.StatusNotFound,
	ErrInternalError:          http.StatusInternalServerError,
	ErrInvalidLoginOrPassword: http.StatusUnauthorized,
}

type ErrorResponse struct {
	Message string `json:"message"`
}
