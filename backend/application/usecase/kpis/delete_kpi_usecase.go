package kpis

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// DeleteKpiUseCase handles the deletion of KPI categories
type DeleteKpiUseCase struct {
	kpiGateway gateway.KpiGateway
}

// NewDeleteKpiUseCase creates a new DeleteKpiUseCase instance
func NewDeleteKpiUseCase(kpiGateway gateway.KpiGateway) *DeleteKpiUseCase {
	return &DeleteKpiUseCase{
		kpiGateway: kpiGateway,
	}
}

// Execute performs the KPI category deletion operation
func (uc *DeleteKpiUseCase) Execute(ctx context.Context, kpiID string, requestingUserID string, isAdmin bool) error {
	// 1. Converter string para UUID
	kpiUUID, err := uuid.Parse(kpiID)
	if err != nil {
		return errors.ErrKpiNotFound
	}

	// 2. Verificar se KPI existe antes de deletar
	kpi, err := uc.kpiGateway.FindByID(ctx, kpiUUID)
	if err != nil {
		if err == errors.ErrKpiNotFound {
			return errors.ErrKpiNotFound
		}
		return err
	}

	// 3. Verificar permissão
	// System KPIs (createdBy = nil) - só admin podem deletar
	// User KPIs - owner ou admin podem deletar
	if kpi.CreatedBy() == nil {
		if !isAdmin {
			return fmt.Errorf("system KPIs can only be deleted by admins")
		}
	} else {
		requestingUserUUID, err := uuid.Parse(requestingUserID)
		if err != nil {
			return fmt.Errorf("invalid user ID")
		}
		if *kpi.CreatedBy() != requestingUserUUID && !isAdmin {
			return fmt.Errorf("you can only delete your own KPIs")
		}
	}

	// 4. Deletar usando gateway
	if err := uc.kpiGateway.Delete(ctx, kpiUUID); err != nil {
		return err
	}

	// 5. Retornar nil se sucesso
	return nil
}
