package budget

import (
	"context"

	"github.com/google/uuid"
	budgeterrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type DeleteBudgetItemUseCase struct {
	budgetGateway gateway.BudgetGateway
}

func NewDeleteBudgetItemUseCase(budgetGateway gateway.BudgetGateway) *DeleteBudgetItemUseCase {
	return &DeleteBudgetItemUseCase{
		budgetGateway: budgetGateway,
	}
}

func (uc *DeleteBudgetItemUseCase) Execute(ctx context.Context, id string) error {
	// Parse UUID
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return budgeterrors.ErrInvalidBudgetData
	}

	// Deletar item
	if err := uc.budgetGateway.Delete(ctx, uuidID); err != nil {
		return err
	}

	return nil
}
