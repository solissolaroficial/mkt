package valueobject

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidMonetaryValue  = errors.New("invalid monetary value")
	ErrMonetaryValueNegative = errors.New("monetary value cannot be negative")
)

type MonetaryValue struct {
	value float64
}

func NewMonetaryValue(value float64) (*MonetaryValue, error) {
	if value < 0 {
		return nil, ErrMonetaryValueNegative
	}

	return &MonetaryValue{value: value}, nil
}

func ReconstructMonetaryValue(value float64) *MonetaryValue {
	return &MonetaryValue{value: value}
}

func (m *MonetaryValue) Value() float64 {
	return m.value
}

func (m *MonetaryValue) String() string {
	return fmt.Sprintf("R$ %.2f", m.value)
}

// FormatBrazilian retorna o valor no formato brasileiro (R$ 1.234,56)
func (m *MonetaryValue) FormatBrazilian() string {
	return fmt.Sprintf("R$ %.2f", m.value)
}
