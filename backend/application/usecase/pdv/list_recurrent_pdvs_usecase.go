package pdv

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// ListRecurrentPdvsUseCase lista PDVs recorrentes baseado em critérios
type ListRecurrentPdvsUseCase struct {
	gateway gateway.RecurrentPdvGateway
}

// ListRecurrentPdvsInput representa os dados de entrada para listar PDVs recorrentes
type ListRecurrentPdvsInput struct {
	RepName   *string
	City      *string
	Page      int
	Limit     int
	SortBy    *string
	SortOrder *string
}

// NewListRecurrentPdvs cria uma nova instância do ListRecurrentPdvsUseCase
func NewListRecurrentPdvs(gateway gateway.RecurrentPdvGateway) *ListRecurrentPdvsUseCase {
	return &ListRecurrentPdvsUseCase{gateway: gateway}
}

// Execute lista PDVs recorrentes baseado em critérios
func (uc *ListRecurrentPdvsUseCase) Execute(ctx context.Context, input ListRecurrentPdvsInput) ([]*entity.RecurrentPdv, int64, error) {
	// Criar criteria - Criteria está no package domain
	crit := domain.NewRecurrentPdvCriteria()

	if input.RepName != nil {
		crit = crit.WithRepName(input.RepName)
	}

	if input.City != nil {
		crit = crit.WithCity(input.City)
	}

	// Criar paginação
	pagination := valueobject.NewPagination(input.Page, input.Limit)

	// Criar ordenação
	var sortBy string
	if input.SortBy != nil {
		sortBy = *input.SortBy
	} else {
		sortBy = "name"
	}

	var sortOrderStr string
	if input.SortOrder != nil {
		sortOrderStr = *input.SortOrder
	} else {
		sortOrderStr = "ASC"
	}

	sortOrder, err := valueobject.NewSortOrder(sortBy, valueobject.SortDirection(sortOrderStr))
	if err != nil {
		// Usar valor padrão em caso de erro
		sortOrder, _ = valueobject.NewSortOrder("name", valueobject.SortDirectionAsc)
	}

	// Buscar PDVs
	pdvs, err := uc.gateway.FindByCriteria(ctx, crit, &pagination, sortOrder)
	if err != nil {
		return nil, 0, err
	}

	// Contar total
	total, err := uc.gateway.CountByCriteria(ctx, crit)
	if err != nil {
		return nil, 0, err
	}

	return pdvs, total, nil
}
