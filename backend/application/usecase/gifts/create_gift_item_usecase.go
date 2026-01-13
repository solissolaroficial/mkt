package gifts

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// CreateGiftItemInput define os dados de entrada para criar um GiftItem
type CreateGiftItemInput struct {
	Name  string
	Stock int
	Price float64
}

// CreateGiftItemOutput define os dados de saída após criar um GiftItem
type CreateGiftItemOutput struct {
	ID    uuid.UUID
	Name  string
	Stock int
	Price float64
}

// CreateGiftItemUseCase define o use case para criar um GiftItem
type CreateGiftItemUseCase struct {
	giftItemGateway gateway.GiftItemGateway
}

// NewCreateGiftItemUseCase cria uma nova instância de CreateGiftItemUseCase
func NewCreateGiftItemUseCase(giftItemGateway gateway.GiftItemGateway) *CreateGiftItemUseCase {
	return &CreateGiftItemUseCase{
		giftItemGateway: giftItemGateway,
	}
}

// Execute executa o use case para criar um GiftItem
func (uc *CreateGiftItemUseCase) Execute(input CreateGiftItemInput) (*CreateGiftItemOutput, error) {
	// Verificar se já existe um item com o mesmo nome
	exists, err := uc.giftItemGateway.ExistsByName(input.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, domainerrors.ErrGiftItemNameAlreadyUsed
	}

	// Criar a entidade GiftItem
	giftItem, err := entity.NewGiftItem(input.Name, input.Stock, input.Price)
	if err != nil {
		return nil, err
	}

	// Salvar no banco de dados
	if err := uc.giftItemGateway.Create(giftItem); err != nil {
		return nil, err
	}

	return &CreateGiftItemOutput{
		ID:    giftItem.ID(),
		Name:  giftItem.Name().Value(),
		Stock: giftItem.Stock(),
		Price: giftItem.Price().Value(),
	}, nil
}
