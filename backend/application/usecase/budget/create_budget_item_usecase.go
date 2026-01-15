package budget

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	budgeterrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type CreateBudgetItemInput struct {
	CodObj       string
	Obj          string
	CodGrp       string
	Grp          string
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
	if input.CodObj == "" {
		return nil, budgeterrors.ErrInvalidCodObj
	}
	if input.Obj == "" {
		return nil, budgeterrors.ErrInvalidObj
	}
	if input.CodGrp == "" {
		return nil, budgeterrors.ErrInvalidCodGrp
	}
	if input.Grp == "" {
		return nil, budgeterrors.ErrInvalidGrp
	}
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

	// Verificar se já existe um item com o mesmo código
	exists, err := uc.budgetGateway.ExistsByCode(ctx, input.CodObj, input.CodGrp, input.Cod, input.Year)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, budgeterrors.ErrBudgetAlreadyExists
	}

	// Criar a entidade
	budget, err := entity.NewBudgetItem(
		input.CodObj,
		input.Obj,
		input.CodGrp,
		input.Grp,
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
