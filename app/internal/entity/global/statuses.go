package global

import (
	"errors"
	"net/http"
)

var (
	// Err
	ErrNoData                 = errors.New("data not found")
	ErrInvalidParam           = errors.New("invalid request params")
	ErrInternalError          = errors.New("internal server error")
	ErrInvalidLoginOrPassword = errors.New("invalid login or password")
	ErrUserAllreadyExists     = errors.New("user allready exist")
	ErrNotAllowedToUse        = errors.New("you are not allowev to use this method")
	ErrExpiredSesstion        = errors.New("your session expired")

	// Success
	SuccessLogedOut = "successfully logged out"
)

var ErrStatusCodes = map[error]int{
	ErrNoData:                 http.StatusNotFound,
	ErrInternalError:          http.StatusInternalServerError,
	ErrInvalidParam:           http.StatusBadRequest,
	ErrInvalidLoginOrPassword: http.StatusUnauthorized,
	ErrUserAllreadyExists:     http.StatusConflict,
	ErrNotAllowedToUse:        http.StatusUnauthorized,
	ErrExpiredSesstion:        http.StatusUnauthorized,
}

var SuccessStatuses = map[string]int{
	SuccessLogedOut: http.StatusOK,
}
