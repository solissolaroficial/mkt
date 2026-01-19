package budget

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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

		// Parse UUIDs se fornecidos
		var objectUUID *uuid.UUID
		if item.ObjectUUID != nil && *item.ObjectUUID != "" {
			parsedUUID, err := uuid.Parse(*item.ObjectUUID)
			if err != nil {
				return nil, fmt.Errorf("invalid object UUID: %w", err)
			}
			objectUUID = &parsedUUID
		}

		var groupUUID *uuid.UUID
		if item.GroupUUID != nil && *item.GroupUUID != "" {
			parsedUUID, err := uuid.Parse(*item.GroupUUID)
			if err != nil {
				return nil, fmt.Errorf("invalid group UUID: %w", err)
			}
			groupUUID = &parsedUUID
		}

		// Criar a entidade
		budget, err := entity.NewBudgetItem(
			objectUUID,
			item.ObjectName,
			groupUUID,
			item.GroupName,
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
