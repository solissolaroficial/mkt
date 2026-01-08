package valueobject

import (
	"errors"
	"strings"
)

// PdvStatus representa o status de um post PDV
// Value Object imutável para encapsular validação de status
type PdvStatus struct {
	value string
}

// Constantes de status válidos
const (
	PdvStatusPending   string = "pending"
	PdvStatusApproved  string = "approved"
	PdvStatusRejected  string = "rejected"
	PdvStatusPublished string = "published"
	PdvStatusCancelled string = "cancelled"
)

// NewPdvStatus cria um novo PdvStatus com validação
func NewPdvStatus(status string) (*PdvStatus, error) {
	status = strings.ToLower(strings.TrimSpace(status))

	if status == "" {
		return nil, errors.New("status não pode ser vazio")
	}

	if !isValidPdvStatus(status) {
		return nil, errors.New("status inválido: " + status)
	}

	return &PdvStatus{value: status}, nil
}

// isValidPdvStatus verifica se o status é válido
func isValidPdvStatus(status string) bool {
	switch status {
	case PdvStatusPending, PdvStatusApproved, PdvStatusRejected, PdvStatusPublished, PdvStatusCancelled:
		return true
	default:
		return false
	}
}

// Value retorna o status encapsulado
func (s *PdvStatus) Value() string {
	return s.value
}

// String retorna o status como string
func (s *PdvStatus) String() string {
	return s.value
}

// Equals verifica se dois status são iguais
func (s *PdvStatus) Equals(other *PdvStatus) bool {
	if other == nil {
		return false
	}
	return s.value == other.value
}

// IsPending verifica se o status é pending
func (s *PdvStatus) IsPending() bool {
	return s.value == PdvStatusPending
}

// IsApproved verifica se o status é approved
func (s *PdvStatus) IsApproved() bool {
	return s.value == PdvStatusApproved
}

// IsRejected verifica se o status é rejected
func (s *PdvStatus) IsRejected() bool {
	return s.value == PdvStatusRejected
}

// IsPublished verifica se o status é published
func (s *PdvStatus) IsPublished() bool {
	return s.value == PdvStatusPublished
}

// IsCancelled verifica se o status é cancelled
func (s *PdvStatus) IsCancelled() bool {
	return s.value == PdvStatusCancelled
}

// CanTransitionTo verifica se é possível fazer transição para outro status
func (s *PdvStatus) CanTransitionTo(newStatus *PdvStatus) bool {
	if newStatus == nil {
		return false
	}

	// Transições permitidas
	switch s.value {
	case PdvStatusPending:
		return newStatus.value == PdvStatusApproved ||
			newStatus.value == PdvStatusRejected ||
			newStatus.value == PdvStatusCancelled
	case PdvStatusApproved:
		return newStatus.value == PdvStatusPublished ||
			newStatus.value == PdvStatusCancelled
	case PdvStatusRejected:
		return newStatus.value == PdvStatusPending ||
			newStatus.value == PdvStatusCancelled
	case PdvStatusPublished:
		return false // Uma vez publicado, não pode mudar
	case PdvStatusCancelled:
		return false // Uma vez cancelado, não pode mudar
	default:
		return false
	}
}
