package kpis

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetKpiUseCase handles the retrieval of a KPI category by ID
type GetKpiUseCase struct {
	kpiGateway gateway.KpiGateway
}

// NewGetKpiUseCase creates a new GetKpiUseCase instance
func NewGetKpiUseCase(kpiGateway gateway.KpiGateway) *GetKpiUseCase {
	return &GetKpiUseCase{
		kpiGateway: kpiGateway,
	}
}

// Execute performs the KPI category retrieval operation
func (uc *GetKpiUseCase) Execute(ctx context.Context, kpiID string) (*entity.KpiCategory, error) {
	// 1. Converter string para UUID
	id, err := uuid.Parse(kpiID)
	if err != nil {
		return nil, errors.ErrKpiNotFound
	}

	// 2. Buscar usando gateway
	kpi, err := uc.kpiGateway.FindByID(ctx, id)
	if err != nil {
		if err == errors.ErrKpiNotFound {
			return nil, errors.ErrKpiNotFound
		}
		return nil, err
	}

	// 3. Retornar entity
	return kpi, nil
}
