package gifts

import (
	"github.com/google/uuid"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// UpdateGiftItemInput define os dados de entrada para atualizar um GiftItem
type UpdateGiftItemInput struct {
	ID    uuid.UUID
	Name  *string
	Stock *int
	Price *float64
}

// UpdateGiftItemOutput define os dados de saída após atualizar um GiftItem
type UpdateGiftItemOutput struct {
	ID        uuid.UUID
	Name      string
	Stock     int
	Price     float64
	CreatedAt string
	UpdatedAt string
}

// UpdateGiftItemUseCase define o use case para atualizar um GiftItem
type UpdateGiftItemUseCase struct {
	giftItemGateway gateway.GiftItemGateway
}

// NewUpdateGiftItemUseCase cria uma nova instância de UpdateGiftItemUseCase
func NewUpdateGiftItemUseCase(giftItemGateway gateway.GiftItemGateway) *UpdateGiftItemUseCase {
	return &UpdateGiftItemUseCase{
		giftItemGateway: giftItemGateway,
	}
}

// Execute executa o use case para atualizar um GiftItem
func (uc *UpdateGiftItemUseCase) Execute(input UpdateGiftItemInput) (*UpdateGiftItemOutput, error) {
	// Buscar o item existente
	giftItem, err := uc.giftItemGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	if giftItem == nil {
		return nil, domainerrors.ErrGiftItemNotFound
	}

	// Atualizar campos se fornecidos
	if input.Name != nil {
		// Verificar se o novo nome já está em uso por outro item
		if *input.Name != giftItem.Name().Value() {
			exists, err := uc.giftItemGateway.ExistsByName(*input.Name)
			if err != nil {
				return nil, err
			}
			if exists {
				return nil, domainerrors.ErrGiftItemNameAlreadyUsed
			}
		}
		if err := giftItem.UpdateName(*input.Name); err != nil {
			return nil, err
		}
	}
	if input.Stock != nil {
		if err := giftItem.UpdateStock(*input.Stock); err != nil {
			return nil, err
		}
	}
	if input.Price != nil {
		if err := giftItem.UpdatePrice(*input.Price); err != nil {
			return nil, err
		}
	}

	// Validar a entidade
	if err := giftItem.Validate(); err != nil {
		return nil, err
	}

	// Salvar as alterações
	if err := uc.giftItemGateway.Update(giftItem); err != nil {
		return nil, err
	}

	return &UpdateGiftItemOutput{
		ID:        giftItem.ID(),
		Name:      giftItem.Name().Value(),
		Stock:     giftItem.Stock(),
		Price:     giftItem.Price().Value(),
		CreatedAt: giftItem.CreatedAt().Format("2006-01-02 15:04:05"),
		UpdatedAt: giftItem.UpdatedAt().Format("2006-01-02 15:04:05"),
	}, nil
}
