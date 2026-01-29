package kpis

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	kpiErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// DeleteMonthlyData handles the deletion of a monthly data record
type DeleteMonthlyData struct {
	kpiGateway gateway.KpiGateway
}

// NewDeleteMonthlyData creates a new DeleteMonthlyData use case
func NewDeleteMonthlyData(kpiGateway gateway.KpiGateway) *DeleteMonthlyData {
	return &DeleteMonthlyData{
		kpiGateway: kpiGateway,
	}
}

// Execute deletes a monthly data record by its ID
// It adds a deletion log before soft deleting the record for audit purposes
func (uc *DeleteMonthlyData) Execute(ctx context.Context, monthlyDataID string) error {
	// 1. Convert string to UUID
	id, err := uuid.Parse(monthlyDataID)
	if err != nil {
		return kpiErrors.ErrMonthDataNotFound
	}

	// 2. Find MonthlyData BEFORE deleting to capture data for log
	monthlyData, err := uc.kpiGateway.FindMonthlyDataByID(ctx, id)
	if err != nil {
		return err
	}

	// 3. Add deletion log BEFORE soft delete
	// This ensures the log is saved before the record is marked as deleted
	userName := "System" // TODO: Extract user from JWT context
	zeroValue := 0.0
	if err := monthlyData.AddLog(entity.KpiLogEntry{
		ID:        uuid.New().String(),
		Date:      time.Now().Format("2006-01-02"),
		Timestamp: time.Now().Format("15:04"),
		User:      userName,
		Month:     monthlyData.Month(),
		OldValue:  monthlyData.Realized(),
		NewValue:  zeroValue,
		Action:    "delete",
		Context:   "Admin deletion",
	}); err != nil {
		return err
	}

	// 4. Save the monthly data with the deletion log
	if err := uc.kpiGateway.UpdateMonthlyData(ctx, monthlyData); err != nil {
		return err
	}

	// 5. Perform soft delete
	if err := uc.kpiGateway.DeleteMonthlyData(ctx, id); err != nil {
		return err
	}

	return nil
}
