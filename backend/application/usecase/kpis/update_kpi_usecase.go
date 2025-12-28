package kpis

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// UpdateKpiInput represents the input data for updating a KPI category
type UpdateKpiInput struct {
	Title *string `json:"title,omitempty"`
	Color *string `json:"color,omitempty"`
	Unit  *string `json:"unit,omitempty"`
}

// UpdateKpiUseCase handles the updating of KPI categories
type UpdateKpiUseCase struct {
	kpiGateway gateway.KpiGateway
}

// NewUpdateKpiUseCase creates a new UpdateKpiUseCase instance
func NewUpdateKpiUseCase(kpiGateway gateway.KpiGateway) *UpdateKpiUseCase {
	return &UpdateKpiUseCase{
		kpiGateway: kpiGateway,
	}
}

// Execute performs the KPI category update operation
func (uc *UpdateKpiUseCase) Execute(ctx context.Context, kpiID string, input UpdateKpiInput) (*entity.KpiCategory, error) {
	// 1. Converter string para UUID
	id, err := uuid.Parse(kpiID)
	if err != nil {
		return nil, errors.ErrKpiNotFound
	}

	// 2. Buscar KPI existente usando gateway
	kpi, err := uc.kpiGateway.FindByID(ctx, id)
	if err != nil {
		if err == errors.ErrKpiNotFound {
			return nil, errors.ErrKpiNotFound
		}
		return nil, err
	}

	// 3. Atualizar campos usando setters da entity
	if input.Title != nil {
		if err := kpi.UpdateTitle(*input.Title); err != nil {
			return nil, err
		}
	}

	// Update color if provided
	if input.Color != nil {
		if err := kpi.UpdateColor(*input.Color); err != nil {
			return nil, err
		}
	}

	// Update unit if provided
	if input.Unit != nil {
		if err := kpi.UpdateUnit(*input.Unit); err != nil {
			return nil, err
		}
	}

	// 4. Atualizar usando gateway
	if err := uc.kpiGateway.Update(ctx, kpi); err != nil {
		return nil, err
	}

	// 5. Retornar entity atualizada
	return kpi, nil
}
