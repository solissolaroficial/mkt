package gateway

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

// MonthlyDataGateway defines the interface for monthly data operations
type MonthlyDataGateway interface {
	// Save creates new monthly data or updates an existing one
	Save(ctx context.Context, data *entity.MonthlyData) error

	// FindByKpiAndMonth retrieves monthly data for a specific KPI and month
	FindByKpiAndMonth(ctx context.Context, kpiID uuid.UUID, month string) (*entity.MonthlyData, error)

	// FindAllByKpi retrieves all monthly data for a specific KPI
	FindAllByKpi(ctx context.Context, kpiID uuid.UUID) ([]*entity.MonthlyData, error)

	// Update modifies existing monthly data
	Update(ctx context.Context, data *entity.MonthlyData) error

	// UpdateMeta updates only the meta value for a specific KPI and month
	UpdateMeta(ctx context.Context, kpiID uuid.UUID, month string, meta float64) error

	// UpdateRealized updates only the realized value for a specific KPI and month
	UpdateRealized(ctx context.Context, kpiID uuid.UUID, month string, realized float64) error

	// UpdateBreakdown updates only the breakdown data for a specific KPI and month
	UpdateBreakdown(ctx context.Context, kpiID uuid.UUID, month string, breakdown interface{}) error
}
