package valueobject

import (
	"errors"
)

var (
	ErrInvalidRecurrenceType = errors.New("invalid recurrence type")
)

type RecurrenceType string

const (
	RecurrenceTypeNone    RecurrenceType = "none"
	RecurrenceTypeMonthly RecurrenceType = "monthly"
	RecurrenceTypeYearly  RecurrenceType = "yearly"
)

func NewRecurrenceType(recurrenceType string) (RecurrenceType, error) {
	switch RecurrenceType(recurrenceType) {
	case RecurrenceTypeNone, RecurrenceTypeMonthly, RecurrenceTypeYearly:
		return RecurrenceType(recurrenceType), nil
	default:
		return "", ErrInvalidRecurrenceType
	}
}

// ReconstructRecurrenceType reconstrói o RecurrenceType de dados do banco
func ReconstructRecurrenceType(recurrenceType string) RecurrenceType {
	return RecurrenceType(recurrenceType)
}

func (r RecurrenceType) String() string {
	return string(r)
}

// IsNone verifica se não tem recorrência
func (r RecurrenceType) IsNone() bool {
	return r == RecurrenceTypeNone
}

// IsMonthly verifica se é mensal
func (r RecurrenceType) IsMonthly() bool {
	return r == RecurrenceTypeMonthly
}

// IsYearly verifica se é anual
func (r RecurrenceType) IsYearly() bool {
	return r == RecurrenceTypeYearly
}
