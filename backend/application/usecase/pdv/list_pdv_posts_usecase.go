package pdv

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// ListPdvPostsUseCase lista posts de PDV baseado em critérios
type ListPdvPostsUseCase struct {
	gateway gateway.PdvPostGateway
}

// ListPdvPostsInput representa os dados de entrada para listar posts de PDV
type ListPdvPostsInput struct {
	RepName   *string
	Month     *string
	Platform  *string
	Status    *string
	StartDate *string
	EndDate   *string
	Page      int
	Limit     int
	SortBy    *string
	SortOrder *string
}

// NewListPdvPosts cria uma nova instância do ListPdvPostsUseCase
func NewListPdvPosts(gateway gateway.PdvPostGateway) *ListPdvPostsUseCase {
	return &ListPdvPostsUseCase{gateway: gateway}
}

// Execute lista posts de PDV baseado em critérios
func (uc *ListPdvPostsUseCase) Execute(ctx context.Context, input ListPdvPostsInput) ([]*entity.PdvPost, int64, error) {
	// Criar criteria - Criteria está no package domain
	crit := domain.NewPdvPostCriteria()

	if input.RepName != nil {
		crit = crit.WithRepName(input.RepName)
	}

	if input.Month != nil {
		crit = crit.WithMonth(input.Month)
	}

	if input.Platform != nil {
		crit = crit.WithPlatform(input.Platform)
	}

	if input.Status != nil {
		status, err := valueobject.NewPdvStatus(*input.Status)
		if err != nil {
			return nil, 0, err
		}
		crit = crit.WithStatus(status)
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

	// Validar critérios
	if err := crit.Validate(); err != nil {
		return nil, 0, err
	}

	// Criar paginação
	pagination := valueobject.NewPagination(input.Page, input.Limit)

	// Criar ordenação
	var sortBy string
	if input.SortBy != nil {
		sortBy = *input.SortBy
	} else {
		sortBy = "post_date"
	}

	var sortOrderStr string
	if input.SortOrder != nil {
		sortOrderStr = *input.SortOrder
	} else {
		sortOrderStr = "DESC"
	}

	sortOrder, err := valueobject.NewSortOrder(sortBy, valueobject.SortDirection(sortOrderStr))
	if err != nil {
		return nil, 0, err
	}

	// Buscar posts
	posts, err := uc.gateway.FindByCriteria(ctx, crit, &pagination, sortOrder)
	if err != nil {
		return nil, 0, err
	}

	// Contar total
	total, err := uc.gateway.CountByCriteria(ctx, crit)
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}
