package kpis

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// CreateKpiInput represents the input data for creating a KPI category
type CreateKpiInput struct {
	Title string `json:"title"`
	Color string `json:"color"`
	Unit  string `json:"unit"`
}

// CreateKpiUseCase handles the creation of new KPI categories
type CreateKpiUseCase struct {
	kpiGateway gateway.KpiGateway
}

// NewCreateKpiUseCase creates a new CreateKpiUseCase instance
func NewCreateKpiUseCase(kpiGateway gateway.KpiGateway) *CreateKpiUseCase {
	return &CreateKpiUseCase{
		kpiGateway: kpiGateway,
	}
}

// Execute performs the KPI category creation operation
func (uc *CreateKpiUseCase) Execute(ctx context.Context, input CreateKpiInput) (*entity.KpiCategory, error) {
	// 1. Criar entidade KpiCategory já construída com validação
	kpi, err := entity.NewKpiCategory(input.Title, input.Color, input.Unit)
	if err != nil {
		return nil, err
	}

	// 2. Criar usando gateway
	if err := uc.kpiGateway.Save(ctx, kpi); err != nil {
		return nil, err
	}

	// 3. Retornar entity criada
	return kpi, nil
}
