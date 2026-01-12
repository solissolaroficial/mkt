package offlineaction

import (
	"context"

	"solis/backend/core/domain"
	"solis/backend/core/domain/entity"
	"solis/backend/core/domain/gateway"
	"solis/backend/core/domain/valueobject"
)

type ListOfflineActionsUseCase struct {
	gateway gateway.OfflineActionGateway
}

type ListOfflineActionsInput struct {
	Category  *string
	RepName   *string
	Month     *string
	StartDate *string
	EndDate   *string
	Status    *string
	Page      int
	Limit     int
	SortBy    *string
	SortOrder *string
}

func NewListOfflineActions(gateway gateway.OfflineActionGateway) *ListOfflineActionsUseCase {
	return &ListOfflineActionsUseCase{gateway: gateway}
}

func (uc *ListOfflineActionsUseCase) Execute(ctx context.Context, input ListOfflineActionsInput) ([]*entity.OfflineAction, int64, error) {
	// Criar criteria
	crit := domain.NewOfflineActionCriteria()

	if input.Category != nil {
		category, err := valueobject.NewOfflineCategory(*input.Category)
		if err == nil {
			crit = crit.WithCategory(&category)
		}
	}

	if input.RepName != nil {
		crit = crit.WithRepName(input.RepName)
	}

	if input.Month != nil {
		crit = crit.WithMonth(input.Month)
	}

	if input.StartDate != nil {
		var err error
		crit, err = crit.WithStartDate(input.StartDate)
		if err != nil {
			return nil, 0, err
		}
	}

	if input.EndDate != nil {
		var err error
		crit, err = crit.WithEndDate(input.EndDate)
		if err != nil {
			return nil, 0, err
		}
	}

	if input.Status != nil {
		status, err := valueobject.NewOfflineStatus(*input.Status)
		if err == nil {
			crit = crit.WithStatus(&status)
		}
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
