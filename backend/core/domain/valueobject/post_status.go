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
func (s PostStatus) CanTransitionTo(newStatus PostStatus) bool {
	// Regras de transição:
	// in_progress -> review -> adjust -> approved -> published
	// approved -> published
	// adjust -> review

	transitions := map[PostStatus][]PostStatus{
		StatusInProgress: {StatusReview},
		StatusReview:     {StatusAdjust, StatusApproved},
		StatusAdjust:     {StatusReview},
		StatusApproved:   {StatusPublished},
		StatusPublished:  {}, // Status final
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
