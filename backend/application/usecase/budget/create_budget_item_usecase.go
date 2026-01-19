package budget

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	budgeterrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type CreateBudgetItemInput struct {
	ObjectUUID   *string // UUID string para objeto (opcional)
	ObjectName   string  // Nome do objeto (para compatibilidade)
	GroupUUID    *string // UUID string para grupo (opcional)
	GroupName    string  // Nome do grupo (para compatibilidade)
	Cod          string
	Desc         string
	Vals         []float64
	RealizedVals []float64
	Year         int
}

type CreateBudgetItemUseCase struct {
	budgetGateway gateway.BudgetGateway
}

func NewCreateBudgetItemUseCase(budgetGateway gateway.BudgetGateway) *CreateBudgetItemUseCase {
	return &CreateBudgetItemUseCase{
		budgetGateway: budgetGateway,
	}
}

func (uc *CreateBudgetItemUseCase) Execute(ctx context.Context, input CreateBudgetItemInput) (*entity.BudgetItem, error) {
	// Validar dados de entrada
	if input.Cod == "" {
		return nil, budgeterrors.ErrInvalidCod
	}
	if input.Desc == "" {
		return nil, budgeterrors.ErrInvalidDescription
	}
	if input.Year < 2000 || input.Year > 2100 {
		return nil, budgeterrors.ErrInvalidYear
	}

	// Validar tamanho dos arrays de valores
	if len(input.Vals) != 12 {
		return nil, budgeterrors.ErrInvalidMonthCount
	}
	if len(input.RealizedVals) != 12 {
		return nil, budgeterrors.ErrInvalidMonthCount
	}

	// Validar valores não negativos
	for _, v := range input.Vals {
		if v < 0 {
			return nil, budgeterrors.ErrNegativeValue
		}
	}
	for _, v := range input.RealizedVals {
		if v < 0 {
			return nil, budgeterrors.ErrNegativeValue
		}
	}

	// Parse UUIDs se fornecidos
	var objectUUID *uuid.UUID
	if input.ObjectUUID != nil && *input.ObjectUUID != "" {
		parsedUUID, err := uuid.Parse(*input.ObjectUUID)
		if err != nil {
			return nil, fmt.Errorf("invalid object UUID: %w", err)
		}
		objectUUID = &parsedUUID
	}

	var groupUUID *uuid.UUID
	if input.GroupUUID != nil && *input.GroupUUID != "" {
		parsedUUID, err := uuid.Parse(*input.GroupUUID)
		if err != nil {
			return nil, fmt.Errorf("invalid group UUID: %w", err)
		}
		groupUUID = &parsedUUID
	}

	// Verificar se já existe um item com o mesmo código
	exists, err := uc.budgetGateway.ExistsByCode(ctx, objectUUID, groupUUID, input.Cod, input.Year)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, budgeterrors.ErrBudgetAlreadyExists
	}

	// Criar a entidade
	budget, err := entity.NewBudgetItem(
		objectUUID,
		input.ObjectName,
		groupUUID,
		input.GroupName,
		input.Cod,
		input.Desc,
		input.Vals,
		input.RealizedVals,
		input.Year,
	)
	if err != nil {
		return nil, err
	}

	// Salvar no banco de dados
	if err := uc.budgetGateway.Create(ctx, budget); err != nil {
		return nil, err
	}

	return budget, nil
}
