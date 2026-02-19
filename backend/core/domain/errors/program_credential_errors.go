package errors

import (
	"errors"
)

var (
	ErrProgramCredentialNotFound   = errors.New("program credential not found")
	ErrProgramCredentialNameExists = errors.New("program credential with this name already exists")
	ErrProgramCredentialInvalid    = errors.New("invalid program credential data")
)
