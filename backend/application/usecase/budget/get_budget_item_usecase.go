package budget

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	budgeterrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type GetBudgetItemUseCase struct {
	budgetGateway gateway.BudgetGateway
}

func NewGetBudgetItemUseCase(budgetGateway gateway.BudgetGateway) *GetBudgetItemUseCase {
	return &GetBudgetItemUseCase{
		budgetGateway: budgetGateway,
	}
}

func (uc *GetBudgetItemUseCase) Execute(ctx context.Context, id string) (*entity.BudgetItem, error) {
	// Parse UUID
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, budgeterrors.ErrInvalidBudgetData
	}

	// Buscar item
	budget, err := uc.budgetGateway.FindByID(ctx, uuidID)
	if err != nil {
		return nil, err
	}

	return budget, nil
}
