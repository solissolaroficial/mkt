package budget

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type GetBudgetSummaryUseCase struct {
	budgetGateway gateway.BudgetGateway
}

func NewGetBudgetSummaryUseCase(budgetGateway gateway.BudgetGateway) *GetBudgetSummaryUseCase {
	return &GetBudgetSummaryUseCase{
		budgetGateway: budgetGateway,
	}
}

func (uc *GetBudgetSummaryUseCase) Execute(ctx context.Context, criteria interface{}) ([]*gateway.BudgetSummary, error) {
	return uc.budgetGateway.GetSummary(ctx, criteria)
}
