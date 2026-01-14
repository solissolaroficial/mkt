package accountpayable

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// CreateAccountPayableInput define os dados de entrada para criar um AccountPayable
type CreateAccountPayableInput struct {
	Supplier    string
	Description string
	Amount      float64
	DueDate     string
	Recurrence  string
}

// CreateAccountPayableOutput define os dados de saída após criar um AccountPayable
type CreateAccountPayableOutput struct {
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

// CreateAccountPayableUseCase define o use case para criar um AccountPayable
type CreateAccountPayableUseCase struct {
	gateway gateway.AccountPayableGateway
}

// NewCreateAccountPayable cria uma nova instância de CreateAccountPayableUseCase
func NewCreateAccountPayable(gateway gateway.AccountPayableGateway) *CreateAccountPayableUseCase {
	return &CreateAccountPayableUseCase{gateway: gateway}
}

// Execute executa o use case para criar um AccountPayable
func (uc *CreateAccountPayableUseCase) Execute(ctx context.Context, input CreateAccountPayableInput) (*CreateAccountPayableOutput, error) {
	// Criar entidade (validação interna)
	account, err := entity.NewAccountPayable(
		input.Supplier,
		input.Description,
		input.Amount,
		input.DueDate,
		input.Recurrence,
	)
	if err != nil {
		return nil, err
	}

	// Validar unicidade (fornecedor + data de vencimento)
	exists, err := uc.gateway.ExistsBySupplierAndDueDate(input.Supplier, account.DueDate().Value())
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, domainerrors.ErrAccountPayableAlreadyExists
	}

	// Salvar via gateway
	if err := uc.gateway.Create(account); err != nil {
		return nil, err
	}

	return &CreateAccountPayableOutput{
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
