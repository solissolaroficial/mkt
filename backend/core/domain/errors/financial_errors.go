package errors

import (
	"errors"
)

var (
	ErrAccountPayableNotFound      = errors.New("account payable not found")
	ErrAccountPayableAlreadyExists = errors.New("account payable with this supplier and due date already exists")
	ErrCannotDeleteSentAccount     = errors.New("cannot delete account that was already sent to finance")
)
