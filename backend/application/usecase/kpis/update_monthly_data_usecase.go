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
	KpiID     string      `json:"kpi_id,omitempty"`
	Year      int         `json:"year,omitempty"` // Year (e.g., 2024, 2025)
	Month     string      `json:"month,omitempty"`
	Realized  *float64    `json:"realized,omitempty"`
	Meta      *float64    `json:"meta,omitempty"`
	Breakdown interface{} `json:"breakdown,omitempty"`
	UserID    string      `json:"user_id,omitempty"` // User making the change
	Context   string      `json:"context,omitempty"` // Context of the change
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

// Execute performs monthly data update operation (creates if doesn't exist)
func (uc *UpdateMonthlyDataUseCase) Execute(ctx context.Context, input UpdateMonthlyDataInput) (*entity.MonthlyData, error) {
	// 1. Validar parâmetros obrigatórios
	if input.KpiID == "" || input.Month == "" || input.Year == 0 {
		return nil, errors.ErrMonthDataNotFound
	}

	kpiID, err := uuid.Parse(input.KpiID)
	if err != nil {
		return nil, errors.ErrKpiNotFound
	}

	// 2. Validar que KPI existe
	kpi, err := uc.kpiGateway.FindByID(ctx, kpiID)
	if err != nil {
		if err == errors.ErrKpiNotFound {
			return nil, errors.ErrKpiNotFound
		}
		return nil, err
	}

	// 3. Buscar ou criar MonthlyData
	monthlyData, err := uc.monthlyDataGateway.FindByKpiAndMonth(ctx, kpiID, input.Year, input.Month)
	if err != nil {
		if err == errors.ErrMonthDataNotFound {
			// MonthlyData não existe, criar um novo
			monthlyData, err = entity.NewMonthlyData(kpiID, input.Year, input.Month)
			if err != nil {
				return nil, err
			}

			// Definir valores usando setters
			if input.Realized != nil {
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

			// Adicionar log de criação se houver usuário
			if input.UserID != "" && input.Realized != nil {
				logEntry := entity.KpiLogEntry{
					ID:        uuid.New().String(),
					Date:      time.Now().Format(time.RFC3339),
					Timestamp: time.Now().Format("15:04"),
					User:      input.UserID,
					Month:     input.Month,
					OldValue:  nil,
					NewValue:  *input.Realized,
					Action:    "create",
					Context:   input.Context,
				}
				if err := monthlyData.AddLog(logEntry); err != nil {
					return nil, err
				}
			}

			// Salvar o novo MonthlyData
			if err := uc.monthlyDataGateway.Save(ctx, monthlyData); err != nil {
				return nil, err
			}

			// Adicionar o MonthlyData ao KPI
			kpi.AddMonthlyData(monthlyData)

			return monthlyData, nil
		}
		return nil, err
	}

	// 4. Atualizar campos existentes usando setters da entity
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
