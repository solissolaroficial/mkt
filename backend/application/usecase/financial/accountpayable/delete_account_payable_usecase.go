package accountpayable

import (
	"context"

	"github.com/google/uuid"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// DeleteAccountPayableInput define os dados de entrada para deletar um AccountPayable
type DeleteAccountPayableInput struct {
	ID uuid.UUID
}

// DeleteAccountPayableUseCase define o use case para deletar um AccountPayable
type DeleteAccountPayableUseCase struct {
	gateway gateway.AccountPayableGateway
}

// NewDeleteAccountPayable cria uma nova instância de DeleteAccountPayableUseCase
func NewDeleteAccountPayable(gateway gateway.AccountPayableGateway) *DeleteAccountPayableUseCase {
	return &DeleteAccountPayableUseCase{
		gateway: gateway,
	}
}

// Execute executa o use case para deletar um AccountPayable
func (uc *DeleteAccountPayableUseCase) Execute(ctx context.Context, input DeleteAccountPayableInput) error {
	// Buscar conta para verificar se foi enviada ao financeiro
	account, err := uc.gateway.FindByID(input.ID)
	if err != nil {
		return err
	}
	if account == nil {
		return domainerrors.ErrAccountPayableNotFound
	}

	// Verificar se já foi enviada ao financeiro
	if account.IsSentToFinance() {
		return domainerrors.ErrCannotDeleteSentAccount
	}

	// Deletar (soft delete)
	return uc.gateway.Delete(input.ID)
}
