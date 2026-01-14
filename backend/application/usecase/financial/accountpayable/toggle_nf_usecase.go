package accountpayable

import (
	"context"

	"github.com/google/uuid"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// ToggleNFInput define os dados de entrada para alternar NF
type ToggleNFInput struct {
	ID uuid.UUID
}

// ToggleNFOutput define os dados de saída após alternar NF
type ToggleNFOutput struct {
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

// ToggleNFUseCase define o use case para alternar NF
type ToggleNFUseCase struct {
	gateway gateway.AccountPayableGateway
}

// NewToggleNFUseCase cria uma nova instância de ToggleNFUseCase
func NewToggleNFUseCase(gateway gateway.AccountPayableGateway) *ToggleNFUseCase {
	return &ToggleNFUseCase{gateway: gateway}
}

// Execute executa o use case para alternar NF
func (uc *ToggleNFUseCase) Execute(ctx context.Context, input ToggleNFInput) (*ToggleNFOutput, error) {
	// Buscar conta
	account, err := uc.gateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, domainerrors.ErrAccountPayableNotFound
	}

	// Alternar status da NF
	account.ToggleNF()

	// Salvar via gateway
	if err := uc.gateway.Update(account); err != nil {
		return nil, err
	}

	return &ToggleNFOutput{
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
