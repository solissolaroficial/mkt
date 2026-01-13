package gifts

import (
	"github.com/google/uuid"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// DeleteGiftItemInput define os dados de entrada para deletar um GiftItem
type DeleteGiftItemInput struct {
	ID uuid.UUID
}

// DeleteGiftItemUseCase define o use case para deletar um GiftItem
type DeleteGiftItemUseCase struct {
	giftItemGateway        gateway.GiftItemGateway
	giftTransactionGateway gateway.GiftTransactionGateway
}

// NewDeleteGiftItemUseCase cria uma nova instância de DeleteGiftItemUseCase
func NewDeleteGiftItemUseCase(
	giftItemGateway gateway.GiftItemGateway,
	giftTransactionGateway gateway.GiftTransactionGateway,
) *DeleteGiftItemUseCase {
	return &DeleteGiftItemUseCase{
		giftItemGateway:        giftItemGateway,
		giftTransactionGateway: giftTransactionGateway,
	}
}

// Execute executa o use case para deletar um GiftItem
func (uc *DeleteGiftItemUseCase) Execute(input DeleteGiftItemInput) error {
	// Buscar o item existente
	giftItem, err := uc.giftItemGateway.FindByID(input.ID)
	if err != nil {
		return err
	}
	if giftItem == nil {
		return domainerrors.ErrGiftItemNotFound
	}

	// Verificar se há transações associadas
	count, err := uc.giftTransactionGateway.CountByItemName(giftItem.Name().Value())
	if err != nil {
		return err
	}
	if count > 0 {
		return domainerrors.ErrGiftItemHasTransactions
	}

	// Deletar o item
	if err := uc.giftItemGateway.Delete(input.ID); err != nil {
		return err
	}

	return nil
}
