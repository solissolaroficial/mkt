package errors

import "errors"

// Authentication related error variables
var (
	// ErrInvalidCredentials is returned when email/password combination is invalid
	ErrInvalidCredentials = errors.New("invalid credentials")

	// ErrUserNotFound is returned when user is not found in the system
	ErrUserNotFound = errors.New("user not found")

	// ErrUserInactive is returned when user account is not active
	ErrUserInactive = errors.New("user is inactive")

	// ErrInvalidToken is returned when JWT token is invalid or expired
	ErrInvalidToken = errors.New("invalid token")
)
