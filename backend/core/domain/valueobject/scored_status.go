package valueobject

import (
	"errors"
)

var (
	ErrInvalidScoredStatus = errors.New("invalid scored status")
)

type ScoredStatus string

const (
	ScoredYes    ScoredStatus = "SIM"
	ScoredNo     ScoredStatus = "NÃO"
	ScoredNotYet ScoredStatus = "AINDA NÃO"
)

func NewScoredStatus(status string) (ScoredStatus, error) {
	switch ScoredStatus(status) {
	case ScoredYes, ScoredNo, ScoredNotYet:
		return ScoredStatus(status), nil
	default:
		return "", ErrInvalidScoredStatus
	}
}

func (s ScoredStatus) String() string {
	return string(s)
}

func (s ScoredStatus) IsValid() bool {
	switch s {
	case ScoredYes, ScoredNo, ScoredNotYet:
		return true
	default:
		return false
	}
}
