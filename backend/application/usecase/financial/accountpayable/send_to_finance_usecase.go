package accountpayable

import (
	"context"
	"errors"

	"github.com/google/uuid"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// SendToFinanceInput define os dados de entrada para enviar ao financeiro
type SendToFinanceInput struct {
	ID uuid.UUID
}

// SendToFinanceOutput define os dados de saída após enviar ao financeiro
type SendToFinanceOutput struct {
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

// SendToFinanceUseCase define o use case para enviar ao financeiro
type SendToFinanceUseCase struct {
	gateway gateway.AccountPayableGateway
}

// NewSendToFinanceUseCase cria uma nova instância de SendToFinanceUseCase
func NewSendToFinanceUseCase(gateway gateway.AccountPayableGateway) *SendToFinanceUseCase {
	return &SendToFinanceUseCase{gateway: gateway}
}

// Execute executa o use case para enviar ao financeiro
func (uc *SendToFinanceUseCase) Execute(ctx context.Context, input SendToFinanceInput) (*SendToFinanceOutput, error) {
	// Buscar conta
	account, err := uc.gateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, domainerrors.ErrAccountPayableNotFound
	}

	// Verificar se NF e Boleto chegaram
	if !account.NFArrived() || !account.BoletoArrived() {
		return nil, errors.New("cannot send to finance: NF and Boleto must be received first")
	}

	// Marcar como enviado ao financeiro
	account.SendToFinance()

	// Salvar via gateway
	if err := uc.gateway.Update(account); err != nil {
		return nil, err
	}

	return &SendToFinanceOutput{
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
