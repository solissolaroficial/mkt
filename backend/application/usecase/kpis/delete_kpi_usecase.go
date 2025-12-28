package kpis

import (
	"context"

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
func (uc *DeleteKpiUseCase) Execute(ctx context.Context, kpiID string) error {
	// 1. Converter string para UUID
	id, err := uuid.Parse(kpiID)
	if err != nil {
		return errors.ErrKpiNotFound
	}

	// 2. Verificar se KPI existe antes de deletar (opcional, mas bom para validação)
	_, err = uc.kpiGateway.FindByID(ctx, id)
	if err != nil {
		if err == errors.ErrKpiNotFound {
			return errors.ErrKpiNotFound
		}
		return err
	}

	// 3. Deletar usando gateway
	if err := uc.kpiGateway.Delete(ctx, id); err != nil {
		return err
	}

	// 4. Retornar nil se sucesso
	return nil
}
