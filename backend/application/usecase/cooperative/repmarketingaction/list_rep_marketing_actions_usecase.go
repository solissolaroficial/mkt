package repmarketingaction

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type ListRepMarketingActionsUseCase struct {
	gateway gateway.RepMarketingActionGateway
}

type ListRepMarketingActionsInput struct {
	RepName   *string
	Month     *string
	Page      int
	Limit     int
	SortBy    *string
	SortOrder *string
}

func NewListRepMarketingActions(gateway gateway.RepMarketingActionGateway) *ListRepMarketingActionsUseCase {
	return &ListRepMarketingActionsUseCase{gateway: gateway}
}

func (uc *ListRepMarketingActionsUseCase) Execute(ctx context.Context, input ListRepMarketingActionsInput) ([]*entity.RepMarketingAction, int64, error) {
	// Criar criteria
	crit := domain.NewRepMarketingActionCriteria()

	if input.RepName != nil {
		crit = crit.WithRepName(input.RepName)
	}

	if input.Month != nil {
		crit = crit.WithMonth(input.Month)
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

	// Buscar ações
	actions, err := uc.gateway.FindByCriteria(ctx, crit, &pagination, sortOrder)
	if err != nil {
		return nil, 0, err
	}

	// Contar total
	total, err := uc.gateway.CountByCriteria(ctx, crit)
	if err != nil {
		return nil, 0, err
	}

	return actions, total, nil
}
