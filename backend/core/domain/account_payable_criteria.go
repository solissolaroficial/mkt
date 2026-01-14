package domain

import (
	"errors"
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// AccountPayableCriteria representa filtros para busca de contas a pagar
// NOTA: Este criteria NÃO depende de GORM, seguindo Clean Architecture
type AccountPayableCriteria struct {
	supplier   *valueobject.SupplierName
	minAmount  *float64
	maxAmount  *float64
	status     *valueobject.AccountStatus
	recurrence *valueobject.RecurrenceType
	startDate  *string
	endDate    *string
	page       *int
	limit      *int
	sortBy     *string
	sortOrder  *string
}

func NewAccountPayableCriteria() *AccountPayableCriteria {
	return &AccountPayableCriteria{}
}

func (c *AccountPayableCriteria) WithSupplier(supplier *valueobject.SupplierName) *AccountPayableCriteria {
	c.supplier = supplier
	return c
}

func (c *AccountPayableCriteria) WithMinAmount(minAmount *float64) *AccountPayableCriteria {
	c.minAmount = minAmount
	return c
}

func (c *AccountPayableCriteria) WithMaxAmount(maxAmount *float64) *AccountPayableCriteria {
	c.maxAmount = maxAmount
	return c
}

func (c *AccountPayableCriteria) WithStatus(status *valueobject.AccountStatus) *AccountPayableCriteria {
	c.status = status
	return c
}

func (c *AccountPayableCriteria) WithRecurrence(recurrence *valueobject.RecurrenceType) *AccountPayableCriteria {
	c.recurrence = recurrence
	return c
}

func (c *AccountPayableCriteria) WithStartDate(startDate *string) (*AccountPayableCriteria, error) {
	if startDate != nil {
		_, err := time.Parse("2006-01-02", *startDate)
		if err != nil {
			return nil, errors.New("invalid start_date format, expected YYYY-MM-DD")
		}
	}
	c.startDate = startDate
	return c, nil
}

func (c *AccountPayableCriteria) WithEndDate(endDate *string) (*AccountPayableCriteria, error) {
	if endDate != nil {
		_, err := time.Parse("2006-01-02", *endDate)
		if err != nil {
			return nil, errors.New("invalid end_date format, expected YYYY-MM-DD")
		}
	}
	c.endDate = endDate
	return c, nil
}

func (c *AccountPayableCriteria) WithPage(page *int) *AccountPayableCriteria {
	c.page = page
	return c
}

func (c *AccountPayableCriteria) WithLimit(limit *int) *AccountPayableCriteria {
	c.limit = limit
	return c
}

func (c *AccountPayableCriteria) WithSortBy(sortBy *string) *AccountPayableCriteria {
	c.sortBy = sortBy
	return c
}

func (c *AccountPayableCriteria) WithSortOrder(sortOrder *string) *AccountPayableCriteria {
	c.sortOrder = sortOrder
	return c
}

// Getters para o gateway aplicar os filtros
func (c *AccountPayableCriteria) Supplier() *valueobject.SupplierName     { return c.supplier }
func (c *AccountPayableCriteria) MinAmount() *float64                     { return c.minAmount }
func (c *AccountPayableCriteria) MaxAmount() *float64                     { return c.maxAmount }
func (c *AccountPayableCriteria) Status() *valueobject.AccountStatus      { return c.status }
func (c *AccountPayableCriteria) Recurrence() *valueobject.RecurrenceType { return c.recurrence }
func (c *AccountPayableCriteria) StartDate() *string                      { return c.startDate }
func (c *AccountPayableCriteria) EndDate() *string                        { return c.endDate }
func (c *AccountPayableCriteria) Page() *int                              { return c.page }
func (c *AccountPayableCriteria) Limit() *int                             { return c.limit }
func (c *AccountPayableCriteria) SortBy() *string                         { return c.sortBy }
func (c *AccountPayableCriteria) SortOrder() *string                      { return c.sortOrder }

// GetOffset calcula o offset baseado na página e limite
func (c *AccountPayableCriteria) GetOffset() int {
	if c.page == nil || c.limit == nil {
		return 0
	}
	return (*c.page - 1) * *c.limit
}
