package budget

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type ListBudgetItemsUseCase struct {
	budgetGateway gateway.BudgetGateway
}

func NewListBudgetItemsUseCase(budgetGateway gateway.BudgetGateway) *ListBudgetItemsUseCase {
	return &ListBudgetItemsUseCase{
		budgetGateway: budgetGateway,
	}
}

func (uc *ListBudgetItemsUseCase) Execute(ctx context.Context, criteria interface{}) ([]*entity.BudgetItem, error) {
	return uc.budgetGateway.List(ctx, criteria)
}
