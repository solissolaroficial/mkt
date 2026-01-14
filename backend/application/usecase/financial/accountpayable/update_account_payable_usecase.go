package accountpayable

import (
	"context"

	"github.com/google/uuid"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// UpdateAccountPayableInput define os dados de entrada para atualizar um AccountPayable
type UpdateAccountPayableInput struct {
	ID          uuid.UUID
	Supplier    *string
	Description *string
	Amount      *float64
	DueDate     *string
	Recurrence  *string
}

// UpdateAccountPayableOutput define os dados de saída após atualizar um AccountPayable
type UpdateAccountPayableOutput struct {
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

// UpdateAccountPayableUseCase define o use case para atualizar um AccountPayable
type UpdateAccountPayableUseCase struct {
	gateway gateway.AccountPayableGateway
}

// NewUpdateAccountPayable cria uma nova instância de UpdateAccountPayableUseCase
func NewUpdateAccountPayable(gateway gateway.AccountPayableGateway) *UpdateAccountPayableUseCase {
	return &UpdateAccountPayableUseCase{gateway: gateway}
}

// Execute executa o use case para atualizar um AccountPayable
func (uc *UpdateAccountPayableUseCase) Execute(ctx context.Context, input UpdateAccountPayableInput) (*UpdateAccountPayableOutput, error) {
	// Buscar conta existente
	account, err := uc.gateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, domainerrors.ErrAccountPayableNotFound
	}

	// Atualizar campos se fornecidos
	if input.Supplier != nil {
		if err := account.UpdateSupplier(*input.Supplier); err != nil {
			return nil, err
		}
	}

	if input.Description != nil {
		if err := account.UpdateDescription(*input.Description); err != nil {
			return nil, err
		}
	}

	if input.Amount != nil {
		if err := account.UpdateAmount(*input.Amount); err != nil {
			return nil, err
		}
	}

	if input.DueDate != nil {
		if err := account.UpdateDueDate(*input.DueDate); err != nil {
			return nil, err
		}
	}

	if input.Recurrence != nil {
		if err := account.UpdateRecurrence(*input.Recurrence); err != nil {
			return nil, err
		}
	}

	// Salvar via gateway
	if err := uc.gateway.Update(account); err != nil {
		return nil, err
	}

	return &UpdateAccountPayableOutput{
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
