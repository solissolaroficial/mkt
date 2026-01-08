package errors

import (
	"errors"
)

var (
	// ErrPdvPostNotFound é retornado quando um post de PDV não é encontrado
	ErrPdvPostNotFound = errors.New("pdv post not found")

	// ErrRecurrentPdvNotFound é retornado quando um PDV recorrente não é encontrado
	ErrRecurrentPdvNotFound = errors.New("recurrent pdv not found")

	// ErrInvalidStatusTransition é retornado quando uma transição de status é inválida
	ErrInvalidStatusTransition = errors.New("invalid status transition")

	// ErrInvalidURL é retornado quando uma URL é inválida
	ErrInvalidURL = errors.New("invalid URL format")

	// ErrInvalidDateRange é retornado quando uma data está fora do range permitido
	ErrInvalidDateRange = errors.New("invalid date range")

	// ErrInvalidPlatform é retornado quando uma plataforma é inválida
	ErrInvalidPlatform = errors.New("invalid platform")
)
