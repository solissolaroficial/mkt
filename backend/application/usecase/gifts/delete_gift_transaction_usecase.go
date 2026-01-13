package gifts

import (
	"github.com/google/uuid"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// DeleteGiftTransactionInput define os dados de entrada para deletar uma GiftTransaction
type DeleteGiftTransactionInput struct {
	ID uuid.UUID
}

// DeleteGiftTransactionUseCase define o use case para deletar uma GiftTransaction
type DeleteGiftTransactionUseCase struct {
	giftTransactionGateway gateway.GiftTransactionGateway
}

// NewDeleteGiftTransactionUseCase cria uma nova instância de DeleteGiftTransactionUseCase
func NewDeleteGiftTransactionUseCase(giftTransactionGateway gateway.GiftTransactionGateway) *DeleteGiftTransactionUseCase {
	return &DeleteGiftTransactionUseCase{
		giftTransactionGateway: giftTransactionGateway,
	}
}

// Execute executa o use case para deletar uma GiftTransaction
func (uc *DeleteGiftTransactionUseCase) Execute(input DeleteGiftTransactionInput) error {
	// Buscar a transação existente
	transaction, err := uc.giftTransactionGateway.FindByID(input.ID)
	if err != nil {
		return err
	}
	if transaction == nil {
		return domainerrors.ErrGiftTransactionNotFound
	}

	// Deletar a transação (soft delete)
	if err := uc.giftTransactionGateway.Delete(input.ID); err != nil {
		return err
	}

	return nil
}
