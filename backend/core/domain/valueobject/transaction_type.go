package valueobject

import (
	"errors"
)

var (
	ErrInvalidTransactionType = errors.New("invalid transaction type")
)

type TransactionType string

const (
	TransactionTypeIn  TransactionType = "in"  // Entrada de estoque
	TransactionTypeOut TransactionType = "out" // Saída de estoque
)

func NewTransactionType(transactionType string) (TransactionType, error) {
	switch TransactionType(transactionType) {
	case TransactionTypeIn, TransactionTypeOut:
		return TransactionType(transactionType), nil
	default:
		return "", ErrInvalidTransactionType
	}
}

// ReconstructTransactionType reconstrói o TransactionType de dados do banco
// Assume que os dados do banco são válidos
func ReconstructTransactionType(transactionType string) TransactionType {
	return TransactionType(transactionType)
}

func (t TransactionType) String() string {
	return string(t)
}

func (t TransactionType) IsValid() bool {
	switch t {
	case TransactionTypeIn, TransactionTypeOut:
		return true
	default:
		return false
	}
}

// IsEntry verifica se é uma transação de entrada
func (t TransactionType) IsEntry() bool {
	return t == TransactionTypeIn
}

// IsExit verifica se é uma transação de saída
func (t TransactionType) IsExit() bool {
	return t == TransactionTypeOut
}
