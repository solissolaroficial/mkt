package valueobject

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidGiftName = errors.New("invalid gift name")
	ErrGiftNameTooLong = errors.New("gift name must be at most 200 characters")
)

type GiftName struct {
	name string
}

func NewGiftName(name string) (*GiftName, error) {
	// Validar tamanho
	if len(name) > 200 {
		return nil, ErrGiftNameTooLong
	}

	// Validar que não está vazio após trim
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return nil, ErrInvalidGiftName
	}

	// Validar que não contém apenas caracteres especiais (aceita letras Unicode com acento)
	matched, err := regexp.MatchString(`^[\p{L}0-9\s\-_&\.]+$`, trimmed)
	if err != nil {
		return nil, err
	}
	if !matched {
		return nil, ErrInvalidGiftName
	}

	return &GiftName{name: trimmed}, nil
}

func ReconstructGiftName(name string) *GiftName {
	// Assume que dados do banco são válidos
	return &GiftName{name: name}
}

func (g *GiftName) Value() string {
	return g.name
}

func (g *GiftName) String() string {
	return g.name
}
