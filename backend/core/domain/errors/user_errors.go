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

	// ErrStorageNotConfigured é retornado quando o storage S3 não está configurado
	ErrStorageNotConfigured = errors.New("storage is not configured")

	// ErrFileTooLarge é retornado quando o arquivo excede o tamanho máximo
	ErrFileTooLarge = errors.New("file too large (max 5MB)")

	// ErrInvalidFileType é retornado quando o tipo de arquivo não é permitido
	ErrInvalidFileType = errors.New("invalid file type (only jpg, jpeg, png, gif allowed)")
)
