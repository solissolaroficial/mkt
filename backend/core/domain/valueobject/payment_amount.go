package valueobject

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidPaymentAmount  = errors.New("invalid payment amount")
	ErrPaymentAmountNegative = errors.New("payment amount cannot be negative")
)

type PaymentAmount struct {
	amount float64
}

func NewPaymentAmount(amount float64) (*PaymentAmount, error) {
	if amount < 0 {
		return nil, ErrPaymentAmountNegative
	}

	return &PaymentAmount{amount: amount}, nil
}

func ReconstructPaymentAmount(amount float64) *PaymentAmount {
	return &PaymentAmount{amount: amount}
}

func (p *PaymentAmount) Value() float64 {
	return p.amount
}

func (p *PaymentAmount) String() string {
	return fmt.Sprintf("R$ %.2f", p.amount)
}

// FormatBrazilian retorna o valor no formato brasileiro (R$ 1.234,56)
func (p *PaymentAmount) FormatBrazilian() string {
	return fmt.Sprintf("R$ %.2f", p.amount)
}
