package accountpayable

import (
	"context"

	"github.com/google/uuid"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetAccountPayableInput define os dados de entrada para buscar um AccountPayable
type GetAccountPayableInput struct {
	ID uuid.UUID
}

// GetAccountPayableOutput define os dados de saída após buscar um AccountPayable
type GetAccountPayableOutput struct {
	ID            uuid.UUID
	Supplier      string
	Description   string
	Amount        float64
	DueDate       string
	NFArrived     bool
	BoletoArrived bool
	Status        string
	Recurrence    string
	CreatedAt     string
	UpdatedAt     string
}

// GetAccountPayableUseCase define o use case para buscar um AccountPayable
type GetAccountPayableUseCase struct {
	gateway gateway.AccountPayableGateway
}

// NewGetAccountPayable cria uma nova instância de GetAccountPayableUseCase
func NewGetAccountPayable(gateway gateway.AccountPayableGateway) *GetAccountPayableUseCase {
	return &GetAccountPayableUseCase{gateway: gateway}
}

// Execute executa o use case para buscar um AccountPayable
func (uc *GetAccountPayableUseCase) Execute(ctx context.Context, input GetAccountPayableInput) (*GetAccountPayableOutput, error) {
	account, err := uc.gateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, domainErrors.ErrAccountPayableNotFound
	}

	return &GetAccountPayableOutput{
		ID:            account.ID(),
		Supplier:      account.Supplier().String(),
		Description:   account.Description(),
		Amount:        account.Amount().Value(),
		DueDate:       account.DueDate().String(),
		NFArrived:     account.NFArrived(),
		BoletoArrived: account.BoletoArrived(),
		Status:        account.Status().String(),
		Recurrence:    account.Recurrence().String(),
		CreatedAt:     account.CreatedAt().Format("2006-01-02 15:04:05"),
		UpdatedAt:     account.UpdatedAt().Format("2006-01-02 15:04:05"),
	}, nil
}
