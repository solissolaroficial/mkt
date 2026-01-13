package gifts

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// ListGiftItemsInput define os dados de entrada para listar GiftItems
type ListGiftItemsInput struct {
	Name     *string
	MinStock *int
	MaxStock *int
	MinPrice *float64
	MaxPrice *float64
	Page     *int
	Limit    *int
}

// ListGiftItemsOutput define os dados de saída após listar GiftItems
type ListGiftItemsOutput struct {
	ID        uuid.UUID
	Name      string
	Stock     int
	Price     float64
	CreatedAt string
	UpdatedAt string
}

// ListGiftItemsUseCase define o use case para listar GiftItems
type ListGiftItemsUseCase struct {
	giftItemGateway gateway.GiftItemGateway
}

// NewListGiftItemsUseCase cria uma nova instância de ListGiftItemsUseCase
func NewListGiftItemsUseCase(giftItemGateway gateway.GiftItemGateway) *ListGiftItemsUseCase {
	return &ListGiftItemsUseCase{
		giftItemGateway: giftItemGateway,
	}
}

// Execute executa o use case para listar GiftItems
func (uc *ListGiftItemsUseCase) Execute(input ListGiftItemsInput) ([]*ListGiftItemsOutput, error) {
	// Criar critérios de busca
	criteria := domain.NewGiftItemCriteria()

	if input.Name != nil {
		name, err := valueobject.NewGiftName(*input.Name)
		if err == nil {
			criteria.WithName(name)
		}
	}
	if input.MinStock != nil {
		criteria.WithMinStock(input.MinStock)
	}
	if input.MaxStock != nil {
		criteria.WithMaxStock(input.MaxStock)
	}
	if input.MinPrice != nil {
		criteria.WithMinPrice(input.MinPrice)
	}
	if input.MaxPrice != nil {
		criteria.WithMaxPrice(input.MaxPrice)
	}
	if input.Page != nil {
		criteria.WithPage(input.Page)
	}
	if input.Limit != nil {
		criteria.WithLimit(input.Limit)
	}

	// Buscar GiftItems
	giftItems, err := uc.giftItemGateway.List(criteria)
	if err != nil {
		return nil, err
	}

	// Converter para output
	output := make([]*ListGiftItemsOutput, len(giftItems))
	for i, item := range giftItems {
		output[i] = &ListGiftItemsOutput{
			ID:        item.ID(),
			Name:      item.Name().Value(),
			Stock:     item.Stock(),
			Price:     item.Price().Value(),
			CreatedAt: item.CreatedAt().Format("2006-01-02 15:04:05"),
			UpdatedAt: item.UpdatedAt().Format("2006-01-02 15:04:05"),
		}
	}

	return output, nil
}
