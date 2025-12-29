package kpis

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// UpdateMonthlyDataInput represents input data for updating monthly data
type UpdateMonthlyDataInput struct {
	MonthlyDataID string      `json:"monthly_data_id"`
	KpiID         string      `json:"kpi_id,omitempty"`
	Month         string      `json:"month,omitempty"`
	Realized      *float64    `json:"realized,omitempty"`
	Meta          *float64    `json:"meta,omitempty"`
	Breakdown     interface{} `json:"breakdown,omitempty"`
	UserID        string      `json:"user_id,omitempty"` // User making the change
	Context       string      `json:"context,omitempty"` // Context of the change
}

// UpdateMonthlyDataUseCase handles updating of monthly data for KPIs
type UpdateMonthlyDataUseCase struct {
	monthlyDataGateway gateway.MonthlyDataGateway
	kpiGateway         gateway.KpiGateway
}

// NewUpdateMonthlyDataUseCase creates a new UpdateMonthlyDataUseCase instance
func NewUpdateMonthlyDataUseCase(
	monthlyDataGateway gateway.MonthlyDataGateway,
	kpiGateway gateway.KpiGateway,
) *UpdateMonthlyDataUseCase {
	return &UpdateMonthlyDataUseCase{
		monthlyDataGateway: monthlyDataGateway,
		kpiGateway:         kpiGateway,
	}
}

// Execute performs monthly data update operation
func (uc *UpdateMonthlyDataUseCase) Execute(ctx context.Context, input UpdateMonthlyDataInput) (*entity.MonthlyData, error) {
	var monthlyData *entity.MonthlyData

	// 1. Buscar MonthlyData existente pelo ID se fornecido
	if input.MonthlyDataID != "" {
		_, err := uuid.Parse(input.MonthlyDataID)
		if err != nil {
			return nil, errors.ErrMonthDataNotFound
		}

		// Nota: O gateway atual não tem FindByID, então precisamos buscar por KPI e mês
		// Por enquanto, vamos assumir que temos o KPI e mês
		if input.KpiID == "" || input.Month == "" {
			return nil, errors.ErrMonthDataNotFound
		}

		kpiID, err := uuid.Parse(input.KpiID)
		if err != nil {
			return nil, errors.ErrKpiNotFound
		}

		monthlyData, err = uc.monthlyDataGateway.FindByKpiAndMonth(ctx, kpiID, input.Month)
		if err != nil {
			if err == errors.ErrMonthDataNotFound {
				return nil, errors.ErrMonthDataNotFound
			}
			return nil, err
		}
	} else if input.KpiID != "" && input.Month != "" {
		// 2. Se não tiver ID mas tiver KPI e mês, buscar por esses campos
		kpiID, err := uuid.Parse(input.KpiID)
		if err != nil {
			return nil, errors.ErrKpiNotFound
		}

		monthlyData, err = uc.monthlyDataGateway.FindByKpiAndMonth(ctx, kpiID, input.Month)
		if err != nil {
			if err == errors.ErrMonthDataNotFound {
				return nil, errors.ErrMonthDataNotFound
			}
			return nil, err
		}
	} else {
		return nil, errors.ErrMonthDataNotFound
	}

	// 3. Validar que KPI existe (se tiver KPI ID)
	if input.KpiID != "" {
		kpiID, err := uuid.Parse(input.KpiID)
		if err != nil {
			return nil, errors.ErrKpiNotFound
		}

		_, err = uc.kpiGateway.FindByID(ctx, kpiID)
		if err != nil {
			if err == errors.ErrKpiNotFound {
				return nil, errors.ErrKpiNotFound
			}
			return nil, err
		}
	}

	// 4. Atualizar campos usando setters da entity
	if input.Realized != nil {
		// Create log entry before updating
		if input.UserID != "" {
			logEntry := entity.KpiLogEntry{
				ID:        uuid.New().String(),
				Date:      time.Now().Format(time.RFC3339),
				Timestamp: time.Now().Format("15:04"),
				User:      input.UserID,
				Month:     monthlyData.Month(),
				OldValue:  monthlyData.Realized(),
				NewValue:  *input.Realized,
				Action:    "update",
				Context:   input.Context,
			}
			if err := monthlyData.AddLog(logEntry); err != nil {
				return nil, err
			}
		}
		monthlyData.SetRealized(*input.Realized)
	}

	if input.Meta != nil {
		monthlyData.SetMeta(*input.Meta)
	}

	if input.Breakdown != nil {
		if err := monthlyData.SetBreakdown(input.Breakdown); err != nil {
			return nil, err
		}
	}

	// 5. Atualizar usando gateway
	if err := uc.monthlyDataGateway.Update(ctx, monthlyData); err != nil {
		return nil, err
	}

	// 6. Retornar entity atualizada
	return monthlyData, nil
}
