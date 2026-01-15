package errors

import "errors"

var (
	// Erros de Validação
	ErrInvalidBudgetData  = errors.New("invalid budget data")
	ErrInvalidMonthCount  = errors.New("vals and realizedVals must have exactly 12 values")
	ErrInvalidMonthIndex  = errors.New("month index must be between 0 and 11")
	ErrInvalidYear        = errors.New("year must be between 2000 and 2100")
	ErrInvalidCodObj      = errors.New("invalid codObj")
	ErrInvalidCodGrp      = errors.New("invalid codGrp")
	ErrInvalidCod         = errors.New("invalid cod")
	ErrInvalidDescription = errors.New("description cannot be empty")
	ErrInvalidObj         = errors.New("obj cannot be empty")
	ErrInvalidGrp         = errors.New("grp cannot be empty")
	ErrNegativeValue      = errors.New("budget values cannot be negative")

	// Erros de Negócio
	ErrBudgetNotFound      = errors.New("budget item not found")
	ErrBudgetAlreadyExists = errors.New("budget item with same code already exists")
)
