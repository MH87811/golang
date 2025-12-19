package services

import "errors"

var (
	ErrInvalidCredential = errors.New("invalid credential")
	ErrEmailExists       = errors.New("email already exits")
	ErrUserNotFound      = errors.New("user not found")
)
