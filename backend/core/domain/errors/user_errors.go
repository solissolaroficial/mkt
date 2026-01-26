package errors

import "errors"

var (
	// ErrCurrentPasswordMismatch é retornado quando a senha atual está incorreta
	ErrCurrentPasswordMismatch = errors.New("current password is incorrect")

	// ErrPasswordTooWeak é retornado quando a nova senha não atende aos requisitos mínimos
	ErrPasswordTooWeak = errors.New("password is too weak")

	// ErrPasswordSameAsCurrent é retornado quando a nova senha é igual à atual
	ErrPasswordSameAsCurrent = errors.New("new password must be different from current password")

	// ErrUserEmailExists é retornado quando o e-mail já está em uso por outro usuário
	ErrUserEmailExists = errors.New("user email already exists")
)
