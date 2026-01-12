package valueobject

import (
	"errors"
)

var (
	ErrInvalidOfflineStatus = errors.New("invalid offline status")
)

type OfflineStatus string

const (
	OfflineStatusPending   OfflineStatus = "pending"
	OfflineStatusApproved  OfflineStatus = "approved"
	OfflineStatusRejected  OfflineStatus = "rejected"
	OfflineStatusCompleted OfflineStatus = "completed"
)

func NewOfflineStatus(status string) (OfflineStatus, error) {
	switch OfflineStatus(status) {
	case OfflineStatusPending, OfflineStatusApproved, OfflineStatusRejected, OfflineStatusCompleted:
		return OfflineStatus(status), nil
	default:
		return "", ErrInvalidOfflineStatus
	}
}

func (s OfflineStatus) String() string {
	return string(s)
}

func (s OfflineStatus) IsValid() bool {
	switch s {
	case OfflineStatusPending, OfflineStatusApproved, OfflineStatusRejected, OfflineStatusCompleted:
		return true
	default:
		return false
	}
}

// CanTransitionTo valida se é possível fazer a transição de status
func (s OfflineStatus) CanTransitionTo(newStatus OfflineStatus) bool {
	// Regras de transição:
	// pending -> approved
	// pending -> rejected
	// approved -> completed
	// rejected -> pending (pode reenviar)

	transitions := map[OfflineStatus][]OfflineStatus{
		OfflineStatusPending:   {OfflineStatusApproved, OfflineStatusRejected},
		OfflineStatusApproved:  {OfflineStatusCompleted},
		OfflineStatusRejected:  {OfflineStatusPending},
		OfflineStatusCompleted: {}, // Terminal state
	}

	allowed, exists := transitions[s]
	if !exists {
		return false
	}

	for _, status := range allowed {
		if status == newStatus {
			return true
		}
	}

	return false
}
