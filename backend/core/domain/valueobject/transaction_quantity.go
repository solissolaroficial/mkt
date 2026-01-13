package valueobject

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidTransactionQuantity = errors.New("transaction quantity must be greater than 0")
)

type TransactionQuantity struct {
	quantity int
}

func NewTransactionQuantity(quantity int) (*TransactionQuantity, error) {
	if quantity <= 0 {
		return nil, ErrInvalidTransactionQuantity
	}

	return &TransactionQuantity{quantity: quantity}, nil
}

func ReconstructTransactionQuantity(quantity int) *TransactionQuantity {
	return &TransactionQuantity{quantity: quantity}
}

func (t *TransactionQuantity) Value() int {
	return t.quantity
}

func (t *TransactionQuantity) String() string {
	return fmt.Sprintf("%d", t.quantity)
}
