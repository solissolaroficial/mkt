package valueobject

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidSupplierName = errors.New("invalid supplier name")
	ErrSupplierNameTooLong = errors.New("supplier name must be at most 200 characters")
)

type SupplierName struct {
	name string
}

func NewSupplierName(name string) (*SupplierName, error) {
	// Validar tamanho
	if len(name) > 200 {
		return nil, ErrSupplierNameTooLong
	}

	// Validar que não está vazio após trim
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return nil, ErrInvalidSupplierName
	}

	// Validar que não contém apenas caracteres especiais
	matched, err := regexp.MatchString(`^[a-zA-Z0-9\s\-_&\.]+$`, trimmed)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, ErrInvalidSupplierName
	}

	return &SupplierName{name: trimmed}, nil
}

func ReconstructSupplierName(name string) *SupplierName {
	// Assume que dados do banco são válidos
	return &SupplierName{name: name}
}

func (s *SupplierName) Value() string {
	return s.name
}

func (s *SupplierName) String() string {
	return s.name
}
