package valueobject

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidBrandName = errors.New("invalid brand name")
	ErrBrandNameTooLong = errors.New("brand name must be at most 200 characters")
)

type BrandName struct {
	name string
}

func NewBrandName(name string) (*BrandName, error) {
	// Validar tamanho
	if len(name) > 200 {
		return nil, ErrBrandNameTooLong
	}

	// Validar que não está vazio após trim
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return nil, ErrInvalidBrandName
	}

	// Validar que não contém apenas caracteres especiais
	matched, err := regexp.MatchString(`^[a-zA-Z0-9\s\-_&\.]+$`, trimmed)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, ErrInvalidBrandName
	}

	return &BrandName{name: trimmed}, nil
}

func (b *BrandName) Value() string {
	return b.name
}

func (b *BrandName) String() string {
	return b.name
}

// ReconstructBrandName creates a BrandName from a string without validation
// This is used when reconstructing entities from database
func ReconstructBrandName(name string) *BrandName {
	return &BrandName{name: name}
}
