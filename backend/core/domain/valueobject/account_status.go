package valueobject

import (
	"errors"
)

var (
	ErrInvalidAccountStatus = errors.New("invalid account status")
)

type AccountStatus string

const (
	AccountStatusPending       AccountStatus = "pending"
	AccountStatusSentToFinance AccountStatus = "sent_to_finance"
)

func NewAccountStatus(status string) (AccountStatus, error) {
	switch AccountStatus(status) {
	case AccountStatusPending, AccountStatusSentToFinance:
		return AccountStatus(status), nil
	default:
		return "", ErrInvalidAccountStatus
	}
}

// ReconstructAccountStatus reconstrói o AccountStatus de dados do banco
func ReconstructAccountStatus(status string) AccountStatus {
	return AccountStatus(status)
}

func (s AccountStatus) String() string {
	return string(s)
}

// IsPending verifica se está pendente
func (s AccountStatus) IsPending() bool {
	return s == AccountStatusPending
}

// IsSentToFinance verifica se foi enviado ao financeiro
func (s AccountStatus) IsSentToFinance() bool {
	return s == AccountStatusSentToFinance
}
