package gifts

import (
	"github.com/google/uuid"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetGiftItemInput define os dados de entrada para buscar um GiftItem
type GetGiftItemInput struct {
	ID uuid.UUID
}

// GetGiftItemOutput define os dados de saída após buscar um GiftItem
type GetGiftItemOutput struct {
	ID        uuid.UUID
	Name      string
	Stock     int
	Price     float64
	CreatedAt string
	UpdatedAt string
}

// GetGiftItemUseCase define o use case para buscar um GiftItem
type GetGiftItemUseCase struct {
	giftItemGateway gateway.GiftItemGateway
}

// NewGetGiftItemUseCase cria uma nova instância de GetGiftItemUseCase
func NewGetGiftItemUseCase(giftItemGateway gateway.GiftItemGateway) *GetGiftItemUseCase {
	return &GetGiftItemUseCase{
		giftItemGateway: giftItemGateway,
	}
}

// Execute executa o use case para buscar um GiftItem
func (uc *GetGiftItemUseCase) Execute(input GetGiftItemInput) (*GetGiftItemOutput, error) {
	giftItem, err := uc.giftItemGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	if giftItem == nil {
		return nil, domainerrors.ErrGiftItemNotFound
	}

	return &GetGiftItemOutput{
		ID:        giftItem.ID(),
		Name:      giftItem.Name().Value(),
		Stock:     giftItem.Stock(),
		Price:     giftItem.Price().Value(),
		CreatedAt: giftItem.CreatedAt().Format("2006-01-02 15:04:05"),
		UpdatedAt: giftItem.UpdatedAt().Format("2006-01-02 15:04:05"),
	}, nil
}
