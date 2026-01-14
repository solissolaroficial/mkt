package accountpayable

import (
	"context"

	"github.com/google/uuid"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// ToggleBoletoInput define os dados de entrada para alternar Boleto
type ToggleBoletoInput struct {
	ID uuid.UUID
}

// ToggleBoletoOutput define os dados de saída após alternar Boleto
type ToggleBoletoOutput struct {
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

// ToggleBoletoUseCase define o use case para alternar Boleto
type ToggleBoletoUseCase struct {
	gateway gateway.AccountPayableGateway
}

// NewToggleBoletoUseCase cria uma nova instância de ToggleBoletoUseCase
func NewToggleBoletoUseCase(gateway gateway.AccountPayableGateway) *ToggleBoletoUseCase {
	return &ToggleBoletoUseCase{gateway: gateway}
}

// Execute executa o use case para alternar Boleto
func (uc *ToggleBoletoUseCase) Execute(ctx context.Context, input ToggleBoletoInput) (*ToggleBoletoOutput, error) {
	// Buscar conta
	account, err := uc.gateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, domainerrors.ErrAccountPayableNotFound
	}

	// Alternar status do Boleto
	account.ToggleBoleto()

	// Salvar via gateway
	if err := uc.gateway.Update(account); err != nil {
		return nil, err
	}

	return &ToggleBoletoOutput{
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
