package kpis

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetKpisBySlugsInput represents input for getting KPIs by slugs
type GetKpisBySlugsInput struct {
	Slugs []string `json:"slugs"`
}

// GetKpisBySlugsOutput represents output of getting KPIs by slugs
type GetKpisBySlugsOutput struct {
	Kpis []*entity.KpiCategory `json:"kpis"`
}

// GetKpisBySlugsUseCase handles getting KPI categories by a list of slugs
type GetKpisBySlugsUseCase struct {
	kpiGateway gateway.KpiGateway
}

// NewGetKpisBySlugsUseCase creates a new GetKpisBySlugsUseCase instance
func NewGetKpisBySlugsUseCase(kpiGateway gateway.KpiGateway) *GetKpisBySlugsUseCase {
	return &GetKpisBySlugsUseCase{
		kpiGateway: kpiGateway,
	}
}

// Execute performs the KPI categories retrieval by slugs operation
func (uc *GetKpisBySlugsUseCase) Execute(ctx context.Context, input GetKpisBySlugsInput) (*GetKpisBySlugsOutput, error) {
	// 1. Buscar KPIs usando gateway com lista de slugs
	kpis, err := uc.kpiGateway.FindBySlugs(ctx, input.Slugs)
	if err != nil {
		return nil, err
	}

	// 2. Retornar slice de entities
	return &GetKpisBySlugsOutput{
		Kpis: kpis,
	}, nil
}
