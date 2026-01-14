package accountpayable

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// ListAccountsPayableInput define os dados de entrada para listar AccountPayables
type ListAccountsPayableInput struct {
	Supplier   *string
	MinAmount  *float64
	MaxAmount  *float64
	Status     *string
	Recurrence *string
	StartDate  *string
	EndDate    *string
	Page       *int
	Limit      *int
	SortBy     *string
	SortOrder  *string
}

// ListAccountsPayableOutput define os dados de saída após listar AccountPayables
type ListAccountsPayableOutput struct {
	ID            uuid.UUID
	Supplier      string
	Description   string
	Amount        float64
	DueDate       string
	NFArrived     bool
	BoletoArrived bool
	Status        string
	Recurrence    string
	CreatedAt     string
	UpdatedAt     string
}

// ListAccountsPayableUseCase define o use case para listar AccountPayables
type ListAccountsPayableUseCase struct {
	gateway gateway.AccountPayableGateway
}

// NewListAccountsPayable cria uma nova instância de ListAccountsPayableUseCase
func NewListAccountsPayable(gateway gateway.AccountPayableGateway) *ListAccountsPayableUseCase {
	return &ListAccountsPayableUseCase{gateway: gateway}
}

// Execute executa o use case para listar AccountPayables
func (uc *ListAccountsPayableUseCase) Execute(ctx context.Context, input ListAccountsPayableInput) ([]*ListAccountsPayableOutput, int64, error) {
	// Criar criteria
	crit := domain.NewAccountPayableCriteria()

	if input.Supplier != nil {
		supplier, err := valueobject.NewSupplierName(*input.Supplier)
		if err == nil {
			crit = crit.WithSupplier(supplier)
		}
	}

	if input.MinAmount != nil {
		crit = crit.WithMinAmount(input.MinAmount)
	}

	if input.MaxAmount != nil {
		crit = crit.WithMaxAmount(input.MaxAmount)
	}

	if input.Status != nil {
		status, err := valueobject.NewAccountStatus(*input.Status)
		if err != nil {
			return nil, 0, err
		}
		crit = crit.WithStatus(&status)
	}

	if input.Recurrence != nil {
		recurrence := valueobject.ReconstructRecurrenceType(*input.Recurrence)
		crit = crit.WithRecurrence(&recurrence)
	}

	if input.StartDate != nil {
		var err error
		crit, err = crit.WithStartDate(input.StartDate)
		if err != nil {
			return nil, 0, err
		}
	}

	if input.EndDate != nil {
		var err error
		crit, err = crit.WithEndDate(input.EndDate)
		if err != nil {
			return nil, 0, err
		}
	}

	// Paginação
	if input.Page != nil {
		crit = crit.WithPage(input.Page)
	}
	if input.Limit != nil {
		crit = crit.WithLimit(input.Limit)
	}

	// Ordenação
	if input.SortBy != nil {
		crit = crit.WithSortBy(input.SortBy)
	}
	if input.SortOrder != nil {
		crit = crit.WithSortOrder(input.SortOrder)
	}

	// Buscar contas
	accounts, err := uc.gateway.List(crit)
	if err != nil {
		return nil, 0, err
	}

	// Converter para output
	output := make([]*ListAccountsPayableOutput, len(accounts))
	for i, account := range accounts {
		output[i] = &ListAccountsPayableOutput{
			ID:            account.ID(),
			Supplier:      account.Supplier().String(),
			Description:   account.Description(),
			Amount:        account.Amount().Value(),
			DueDate:       account.DueDate().String(),
			NFArrived:     account.NFArrived(),
			BoletoArrived: account.BoletoArrived(),
			Status:        account.Status().String(),
			Recurrence:    account.Recurrence().String(),
			CreatedAt:     account.CreatedAt().Format("2006-01-02 15:04:05"),
			UpdatedAt:     account.UpdatedAt().Format("2006-01-02 15:04:05"),
		}
	}

	// Contar total de registros
	total, err := uc.gateway.Count(crit)
	if err != nil {
		return nil, 0, err
	}

	return output, total, nil
}
