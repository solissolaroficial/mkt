package kpis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	kpiErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// KPI slugs that require special calculation modes
const (
	AuthorityOnInternetSlug = "autoridade_na_internet_da"
)

// monthToAbbr maps Go's time.Month to Portuguese abbreviations used in the system
var monthToAbbr = map[time.Month]string{
	time.January:   "JAN",
	time.February:  "FEV",
	time.March:     "MAR",
	time.April:     "ABR",
	time.May:       "MAI",
	time.June:      "JUN",
	time.July:      "JUL",
	time.August:    "AGO",
	time.September: "SET",
	time.October:   "OUT",
	time.November:  "NOV",
	time.December:  "DEZ",
}

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

// AddDailyEntryInput represents input for adding a daily entry
type AddDailyEntryInput struct {
	KpiID   string  `json:"kpi_id,omitempty"`
	Year    int     `json:"year,omitempty"`
	Month   string  `json:"month,omitempty"`
	Date    string  `json:"date,omitempty"` // ISO date format (e.g., "2024-11-01")
	Value   float64 `json:"value,omitempty"`
	Context string  `json:"context,omitempty"`
	User    string  `json:"user,omitempty"`
}

// UpdateDailyEntryInput represents input for updating a daily entry
type UpdateDailyEntryInput struct {
	KpiID   string  `json:"kpi_id,omitempty"`
	Year    int     `json:"year,omitempty"`
	Month   string  `json:"month,omitempty"`
	Date    string  `json:"date,omitempty"` // ISO date format (e.g., "2024-11-01")
	Value   float64 `json:"value,omitempty"`
	Context string  `json:"context,omitempty"`
	User    string  `json:"user,omitempty"`
}

// DeleteDailyEntryInput represents input for deleting a daily entry
type DeleteDailyEntryInput struct {
	KpiID   string `json:"kpi_id,omitempty"`
	Year    int    `json:"year,omitempty"`
	Month   string `json:"month,omitempty"`
	Date    string `json:"date,omitempty"` // ISO date format (e.g., "2024-11-01")
	User    string `json:"user,omitempty"`
	Context string `json:"context,omitempty"`
}

// GetDailyEntriesInput represents input for getting daily entries
type GetDailyEntriesInput struct {
	KpiID string `json:"kpi_id,omitempty"`
	Year  int    `json:"year,omitempty"`
	Month string `json:"month,omitempty"`
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
		return nil, kpiErrors.ErrMonthDataNotFound
	}

	kpiID, err := uuid.Parse(input.KpiID)
	if err != nil {
		return nil, kpiErrors.ErrKpiNotFound
	}

	// 2. Validar que KPI existe
	kpi, err := uc.kpiGateway.FindByID(ctx, kpiID)
	if err != nil {
		if err == kpiErrors.ErrKpiNotFound {
			return nil, kpiErrors.ErrKpiNotFound
		}
		return nil, err
	}

	// 3. Buscar ou criar MonthlyData
	monthlyData, err := uc.monthlyDataGateway.FindByKpiAndMonth(ctx, kpiID, input.Year, input.Month)
	if err != nil {
		if err == kpiErrors.ErrMonthDataNotFound {
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

// AddDailyEntry adds a daily entry to the monthly data
func (uc *UpdateMonthlyDataUseCase) AddDailyEntry(ctx context.Context, input AddDailyEntryInput) (*entity.MonthlyData, error) {
	// 1. Validar parâmetros obrigatórios
	if input.KpiID == "" || input.Month == "" || input.Year == 0 || input.Date == "" {
		return nil, errors.New("kpi_id, month, year and date are required")
	}

	kpiID, err := uuid.Parse(input.KpiID)
	if err != nil {
		return nil, kpiErrors.ErrKpiNotFound
	}

	// 2. Buscar ou criar MonthlyData
	monthlyData, err := uc.monthlyDataGateway.FindByKpiAndMonth(ctx, kpiID, input.Year, input.Month)
	if err != nil {
		if err == kpiErrors.ErrMonthDataNotFound {
			// MonthlyData não existe, criar um novo
			monthlyData, err = entity.NewMonthlyData(kpiID, input.Year, input.Month)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// 3. Validar formato da data YYYY-MM-DD
	parsedDate, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return nil, errors.New("invalid date format, expected YYYY-MM-DD")
	}
	// Validar que a data pertence ao mês/ano esperado
	if parsedDate.Year() != monthlyData.Year() {
		return nil, errors.New("date year does not match monthly data year")
	}
	expectedMonth := monthToAbbr[parsedDate.Month()]
	if expectedMonth != monthlyData.Month() {
		return nil, fmt.Errorf("date month %s does not match monthly data month %s", expectedMonth, monthlyData.Month())
	}

	// 4. Adicionar daily entry
	if err := monthlyData.AddDailyEntry(input.Date, input.Value, input.Context, input.User); err != nil {
		return nil, err
	}

	// 5. Determinar modo de cálculo baseado no KPI
	calcMode := entity.RecalculateModeSum
	kpi, err := uc.kpiGateway.FindByID(ctx, kpiID)
	if err == nil && kpi != nil && kpi.Slug() == AuthorityOnInternetSlug {
		calcMode = entity.RecalculateModeLastValue
	}

	// 6. Recalcular realized value from daily entries
	if err := monthlyData.RecalculateFromDailyMode(calcMode); err != nil {
		return nil, err
	}

	// 7. Salvar o MonthlyData
	if err := uc.monthlyDataGateway.Save(ctx, monthlyData); err != nil {
		return nil, err
	}

	// 8. Adicionar log
	if input.User != "" && monthlyData.Realized() != nil {
		logEntry := entity.KpiLogEntry{
			ID:        uuid.New().String(),
			Date:      time.Now().Format(time.RFC3339),
			Timestamp: time.Now().Format("15:04"),
			User:      input.User,
			Month:     monthlyData.Month(),
			OldValue:  nil,
			NewValue:  *monthlyData.Realized(),
			Action:    "daily_entry",
			Context:   input.Context,
		}
		if err := monthlyData.AddLog(logEntry); err != nil {
			return nil, err
		}
		// Update after adding log
		if err := uc.monthlyDataGateway.Update(ctx, monthlyData); err != nil {
			return nil, err
		}
	}

	return monthlyData, nil
}

// UpdateDailyEntry updates an existing daily entry
func (uc *UpdateMonthlyDataUseCase) UpdateDailyEntry(ctx context.Context, input UpdateDailyEntryInput) (*entity.MonthlyData, error) {
	// 1. Validar parâmetros obrigatórios
	if input.KpiID == "" || input.Month == "" || input.Year == 0 || input.Date == "" {
		return nil, errors.New("kpi_id, month, year and date are required")
	}

	kpiID, err := uuid.Parse(input.KpiID)
	if err != nil {
		return nil, kpiErrors.ErrKpiNotFound
	}

	// 2. Buscar ou criar MonthlyData
	monthlyData, err := uc.monthlyDataGateway.FindByKpiAndMonth(ctx, kpiID, input.Year, input.Month)
	if err != nil {
		if err == kpiErrors.ErrMonthDataNotFound {
			// MonthlyData não existe, criar um novo
			monthlyData, err = entity.NewMonthlyData(kpiID, input.Year, input.Month)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// 4. Atualizar daily entry
	if err := monthlyData.UpdateDailyEntry(input.Date, input.Value, input.Context, input.User); err != nil {
		return nil, err
	}

	// 5. Determinar modo de cálculo baseado no KPI
	calcMode := entity.RecalculateModeSum
	kpi, err := uc.kpiGateway.FindByID(ctx, kpiID)
	if err == nil && kpi != nil && kpi.Slug() == AuthorityOnInternetSlug {
		calcMode = entity.RecalculateModeLastValue
	}

	// 6. Recalcular realized value
	if err := monthlyData.RecalculateFromDailyMode(calcMode); err != nil {
		return nil, err
	}

	// 7. Salvar o MonthlyData
	if err := uc.monthlyDataGateway.Save(ctx, monthlyData); err != nil {
		return nil, err
	}

	// 8. Adicionar log
	if input.User != "" && monthlyData.Realized() != nil {
		logEntry := entity.KpiLogEntry{
			ID:        uuid.New().String(),
			Date:      time.Now().Format(time.RFC3339),
			Timestamp: time.Now().Format("15:04"),
			User:      input.User,
			Month:     monthlyData.Month(),
			OldValue:  nil,
			NewValue:  *monthlyData.Realized(),
			Action:    "daily_entry_update",
			Context:   input.Context,
		}
		if err := monthlyData.AddLog(logEntry); err != nil {
			return nil, err
		}
		// Update after adding log
		if err := uc.monthlyDataGateway.Update(ctx, monthlyData); err != nil {
			return nil, err
		}
	}

	return monthlyData, nil
}

// DeleteDailyEntry removes a daily entry by date
func (uc *UpdateMonthlyDataUseCase) DeleteDailyEntry(ctx context.Context, input DeleteDailyEntryInput) (*entity.MonthlyData, error) {
	// 1. Validar parâmetros obrigatórios
	if input.KpiID == "" || input.Month == "" || input.Year == 0 || input.Date == "" {
		return nil, errors.New("kpi_id, month, year and date are required")
	}

	kpiID, err := uuid.Parse(input.KpiID)
	if err != nil {
		return nil, kpiErrors.ErrKpiNotFound
	}

	// 2. Buscar MonthlyData
	monthlyData, err := uc.monthlyDataGateway.FindByKpiAndMonth(ctx, kpiID, input.Year, input.Month)
	if err != nil {
		if err == kpiErrors.ErrMonthDataNotFound {
			return nil, kpiErrors.ErrMonthDataNotFound
		}
		return nil, err
	}

	// 4. Deletar daily entry
	if err := monthlyData.DeleteDailyEntry(input.Date); err != nil {
		return nil, err
	}

	// 5. Determinar modo de cálculo baseado no KPI
	calcMode := entity.RecalculateModeSum
	kpi, err := uc.kpiGateway.FindByID(ctx, kpiID)
	if err == nil && kpi != nil && kpi.Slug() == AuthorityOnInternetSlug {
		calcMode = entity.RecalculateModeLastValue
	}

	// 6. Recalcular realized value
	if err := monthlyData.RecalculateFromDailyMode(calcMode); err != nil {
		return nil, err
	}

	// 7. Salvar o MonthlyData
	if err := uc.monthlyDataGateway.Save(ctx, monthlyData); err != nil {
		return nil, err
	}

	// 8. Adicionar log
	if input.User != "" && monthlyData.Realized() != nil {
		logEntry := entity.KpiLogEntry{
			ID:        uuid.New().String(),
			Date:      time.Now().Format(time.RFC3339),
			Timestamp: time.Now().Format("15:04"),
			User:      input.User,
			Month:     monthlyData.Month(),
			OldValue:  nil,
			NewValue:  *monthlyData.Realized(),
			Action:    "daily_entry_delete",
			Context:   input.Context,
		}
		if err := monthlyData.AddLog(logEntry); err != nil {
			return nil, err
		}
		// Update after adding log
		if err := uc.monthlyDataGateway.Update(ctx, monthlyData); err != nil {
			return nil, err
		}
	}

	return monthlyData, nil
}

// GetDailyEntries retrieves all daily entries for a specific KPI, year and month
func (uc *UpdateMonthlyDataUseCase) GetDailyEntries(ctx context.Context, input GetDailyEntriesInput) ([]entity.DailyEntry, error) {
	// 1. Validar parâmetros obrigatórios
	if input.KpiID == "" || input.Month == "" || input.Year == 0 {
		return nil, errors.New("kpi_id, month, year are required")
	}

	kpiID, err := uuid.Parse(input.KpiID)
	if err != nil {
		return nil, kpiErrors.ErrKpiNotFound
	}

	// 2. Buscar MonthlyData
	monthlyData, err := uc.monthlyDataGateway.FindByKpiAndMonth(ctx, kpiID, input.Year, input.Month)
	if err != nil {
		if err == kpiErrors.ErrMonthDataNotFound {
			return []entity.DailyEntry{}, nil
		}
		return nil, err
	}

	// 3. Retornar daily entries
	return monthlyData.GetDailyEntries()
}
