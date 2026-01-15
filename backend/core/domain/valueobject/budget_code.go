package valueobject

import (
	"errors"
	"regexp"
)

type BudgetCode struct {
	value string
}

// NewBudgetCode cria um novo Value Object para código de orçamento
// Validações:
// - Não pode ser vazio
// - Deve ser alfanumérico (apenas letras e números)
// - Máximo de 20 caracteres
func NewBudgetCode(code string) (*BudgetCode, error) {
	if code == "" {
		return nil, errors.New("budget code cannot be empty")
	}

	// Validar formato: apenas números e letras, max 20 caracteres
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9]{1,20}$`, code)
	if !matched {
		return nil, errors.New("budget code must be alphanumeric and max 20 characters")
	}

	return &BudgetCode{value: code}, nil
}

// ReconstructBudgetCode reconstrói o Value Object sem validação
// Usado ao carregar do banco de dados
func ReconstructBudgetCode(code string) *BudgetCode {
	return &BudgetCode{value: code}
}

// String retorna o valor do código
func (c *BudgetCode) String() string {
	return c.value
}

// Equals verifica se dois códigos são iguais
func (c *BudgetCode) Equals(other *BudgetCode) bool {
	if c == nil || other == nil {
		return c == other
	}
	return c.value == other.value
}

// IsEmpty verifica se o código está vazio
func (c *BudgetCode) IsEmpty() bool {
	return c == nil || c.value == ""
}
