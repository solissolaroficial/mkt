package showroomitem

import (
	"context"

	"solis/backend/core/domain"
	"solis/backend/core/domain/entity"
	"solis/backend/core/domain/gateway"
	"solis/backend/core/domain/valueobject"
)

type ListShowroomItemsUseCase struct {
	gateway gateway.ShowroomItemGateway
}

type ListShowroomItemsInput struct {
	RepName   *string
	City      *string
	Delivered *bool
	Page      int
	Limit     int
	SortBy    *string
	SortOrder *string
}

func NewListShowroomItems(gateway gateway.ShowroomItemGateway) *ListShowroomItemsUseCase {
	return &ListShowroomItemsUseCase{gateway: gateway}
}

func (uc *ListShowroomItemsUseCase) Execute(ctx context.Context, input ListShowroomItemsInput) ([]*entity.ShowroomItem, int64, error) {
	// Criar criteria
	crit := domain.NewShowroomItemCriteria()

	if input.RepName != nil {
		crit = crit.WithRepName(input.RepName)
	}

	if input.City != nil {
		crit = crit.WithCity(input.City)
	}

	if input.Delivered != nil {
		crit = crit.WithDelivered(input.Delivered)
	}

	// Validar critérios
	if err := crit.Validate(); err != nil {
		return nil, 0, err
	}

	// Criar paginação
	pagination := valueobject.NewPagination(input.Page, input.Limit)

	// Criar ordenação
	var sortOrder *valueobject.SortOrder
	if input.SortBy != nil && input.SortOrder != nil {
		var err error
		sortOrder, err = valueobject.NewSortOrder(*input.SortBy, valueobject.SortDirection(*input.SortOrder))
		if err != nil {
			return nil, 0, err
		}
	}

	// Buscar itens
	items, err := uc.gateway.FindByCriteria(ctx, crit, &pagination, sortOrder)
	if err != nil {
		return nil, 0, err
	}

	// Contar total
	total, err := uc.gateway.CountByCriteria(ctx, crit)
	if err != nil {
		return nil, 0, err
	}

	return items, total, nil
}
