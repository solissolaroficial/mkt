package budget

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	budgeterrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type BatchCreateBudgetItemsInput struct {
	Items []CreateBudgetItemInput
}

type BatchCreateBudgetItemsUseCase struct {
	budgetGateway gateway.BudgetGateway
}

func NewBatchCreateBudgetItemsUseCase(budgetGateway gateway.BudgetGateway) *BatchCreateBudgetItemsUseCase {
	return &BatchCreateBudgetItemsUseCase{
		budgetGateway: budgetGateway,
	}
}

func (uc *BatchCreateBudgetItemsUseCase) Execute(ctx context.Context, input BatchCreateBudgetItemsInput) ([]*entity.BudgetItem, error) {
	if len(input.Items) == 0 {
		return []*entity.BudgetItem{}, nil
	}

	// Validar e criar entidades
	budgets := make([]*entity.BudgetItem, 0, len(input.Items))
	year := 0

	for i, item := range input.Items {
		// Validar dados de entrada
		if item.CodObj == "" {
			return nil, budgeterrors.ErrInvalidCodObj
		}
		if item.Obj == "" {
			return nil, budgeterrors.ErrInvalidObj
		}
		if item.CodGrp == "" {
			return nil, budgeterrors.ErrInvalidCodGrp
		}
		if item.Grp == "" {
			return nil, budgeterrors.ErrInvalidGrp
		}
		if item.Cod == "" {
			return nil, budgeterrors.ErrInvalidCod
		}
		if item.Desc == "" {
			return nil, budgeterrors.ErrInvalidDescription
		}
		if item.Year < 2000 || item.Year > 2100 {
			return nil, budgeterrors.ErrInvalidYear
		}

		// Validar tamanho dos arrays de valores
		if len(item.Vals) != 12 {
			return nil, budgeterrors.ErrInvalidMonthCount
		}
		if len(item.RealizedVals) != 12 {
			return nil, budgeterrors.ErrInvalidMonthCount
		}

		// Validar valores não negativos
		for _, v := range item.Vals {
			if v < 0 {
				return nil, budgeterrors.ErrNegativeValue
			}
		}
		for _, v := range item.RealizedVals {
			if v < 0 {
				return nil, budgeterrors.ErrNegativeValue
			}
		}

		// Verificar se todos os itens são do mesmo ano
		if i == 0 {
			year = item.Year
		} else if item.Year != year {
			return nil, budgeterrors.ErrInvalidBudgetData
		}

		// Criar a entidade
		budget, err := entity.NewBudgetItem(
			item.CodObj,
			item.Obj,
			item.CodGrp,
			item.Grp,
			item.Cod,
			item.Desc,
			item.Vals,
			item.RealizedVals,
			item.Year,
		)
		if err != nil {
			return nil, err
		}

		budgets = append(budgets, budget)
	}

	// Salvar em lote
	if err := uc.budgetGateway.BatchCreate(ctx, budgets); err != nil {
		return nil, err
	}

	return budgets, nil
}
