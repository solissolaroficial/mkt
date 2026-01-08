package valueobject

import (
	"errors"
)

var ErrInvalidStatus = errors.New("invalid status")

type PostStatus string

const (
	StatusInProgress PostStatus = "in_progress"
	StatusReview     PostStatus = "review"
	StatusAdjust     PostStatus = "adjust"
	StatusApproved   PostStatus = "approved"
	StatusPublished  PostStatus = "published"
)

func NewPostStatus(status string) (PostStatus, error) {
	switch PostStatus(status) {
	case StatusInProgress, StatusReview, StatusAdjust, StatusApproved, StatusPublished:
		return PostStatus(status), nil
	default:
		return "", ErrInvalidStatus
	}
}

func (s PostStatus) String() string {
	return string(s)
}

func (s PostStatus) IsValid() bool {
	switch s {
	case StatusInProgress, StatusReview, StatusAdjust, StatusApproved, StatusPublished:
		return true
	default:
		return false
	}
}

// CanTransitionTo valida se é possível fazer a transição de status
// Permite QUALQUER transição de status, desde que o novo status seja válido
func (s PostStatus) CanTransitionTo(newStatus PostStatus) bool {
	// Apenas valida que o novo status é válido
	// Não há restrições de transição - permite qualquer mudança de status
	return newStatus.IsValid()
}
