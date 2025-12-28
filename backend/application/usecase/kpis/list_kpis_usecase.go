package kpis

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// ListKpisOutput represents the output of listing KPI categories
type ListKpisOutput struct {
	Kpis  []*entity.KpiCategory `json:"kpis"`
	Total int                   `json:"total"`
}

// ListKpisUseCase handles the listing of all KPI categories
type ListKpisUseCase struct {
	kpiGateway gateway.KpiGateway
}

// NewListKpisUseCase creates a new ListKpisUseCase instance
func NewListKpisUseCase(kpiGateway gateway.KpiGateway) *ListKpisUseCase {
	return &ListKpisUseCase{
		kpiGateway: kpiGateway,
	}
}

// Execute performs the KPI categories listing operation
func (uc *ListKpisUseCase) Execute(ctx context.Context, pagination valueobject.Pagination) (*ListKpisOutput, error) {
	// 1. Listar usando gateway com paginação
	kpis, total, err := uc.kpiGateway.FindAll(ctx, pagination)
	if err != nil {
		return nil, err
	}

	// 2. Retornar slice de entities e total
	return &ListKpisOutput{
		Kpis:  kpis,
		Total: int(total),
	}, nil
}
