package social

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type ListSocialBenchmarkingsUseCase struct {
	gateway gateway.SocialBenchmarkingGateway
}

type ListSocialBenchmarkingsInput struct {
	BrandID   *string
	Active    *bool
	StartDate *string
	EndDate   *string
	Page      int
	Limit     int
	SortBy    *string
	SortOrder *string
}

func NewListSocialBenchmarkings(gateway gateway.SocialBenchmarkingGateway) *ListSocialBenchmarkingsUseCase {
	return &ListSocialBenchmarkingsUseCase{gateway: gateway}
}

func (uc *ListSocialBenchmarkingsUseCase) Execute(ctx context.Context, input ListSocialBenchmarkingsInput) ([]*entity.SocialBenchmarking, int64, error) {
	// Criar criteria
	crit := domain.NewSocialBenchmarkingCriteria()

	if input.BrandID != nil {
		brandID, err := uuid.Parse(*input.BrandID)
		if err == nil {
			crit = crit.WithBrandID(&brandID)
		}
	}

	if input.Active != nil {
		crit = crit.WithActive(input.Active)
	}

	if input.StartDate != nil {
		crit = crit.WithStartDate(input.StartDate)
	}

	if input.EndDate != nil {
		crit = crit.WithEndDate(input.EndDate)
	}

	// Criar paginação
	pagination := valueobject.NewPagination(input.Page, input.Limit)

	// Criar ordenação
	var sortOrder *valueobject.SortOrder
	if input.SortBy != nil && input.SortOrder != nil {
		sortOrder, _ = valueobject.NewSortOrder(*input.SortBy, valueobject.SortDirection(*input.SortOrder))
	} else {
		// Default: ordenar por engagement_rate DESC
		sortOrder, _ = valueobject.NewSortOrder("engagement_rate", valueobject.SortDirectionDesc)
	}

	// Buscar benchmarkings
	benchmarkings, err := uc.gateway.FindByCriteria(ctx, crit, &pagination, sortOrder)
	if err != nil {
		return nil, 0, err
	}

	// Contar total
	total, err := uc.gateway.CountByCriteria(ctx, crit)
	if err != nil {
		return nil, 0, err
	}

	return benchmarkings, total, nil
}
