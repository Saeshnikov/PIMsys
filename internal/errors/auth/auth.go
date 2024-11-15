package auth_errors

import "errors"

var (
	ErrUserExists         = errors.New("user already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrAppNotFound        = errors.New("app not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
