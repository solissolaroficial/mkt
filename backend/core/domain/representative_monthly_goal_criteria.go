package domain

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// RepresentativeMonthlyGoalCriteria defines criteria for querying representative monthly goals
type RepresentativeMonthlyGoalCriteria struct {
	RepresentativeID *uuid.UUID
	Month            *int
	Year             *int
	Pagination       valueobject.Pagination
	SortBy           string
	SortOrder        *valueobject.SortOrder
}

// NewRepresentativeMonthlyGoalCriteria creates a new RepresentativeMonthlyGoalCriteria
func NewRepresentativeMonthlyGoalCriteria() *RepresentativeMonthlyGoalCriteria {
	sortOrder, _ := valueobject.NewSortOrder("created_at", valueobject.SortDirectionDesc)
	return &RepresentativeMonthlyGoalCriteria{
		Pagination: valueobject.NewPagination(1, 10),
		SortBy:     "created_at",
		SortOrder:  sortOrder,
	}
}

// WithRepresentativeID sets the representative ID filter
func (c *RepresentativeMonthlyGoalCriteria) WithRepresentativeID(id uuid.UUID) *RepresentativeMonthlyGoalCriteria {
	c.RepresentativeID = &id
	return c
}

// WithMonth sets the month filter
func (c *RepresentativeMonthlyGoalCriteria) WithMonth(month int) *RepresentativeMonthlyGoalCriteria {
	c.Month = &month
	return c
}

// WithYear sets the year filter
func (c *RepresentativeMonthlyGoalCriteria) WithYear(year int) *RepresentativeMonthlyGoalCriteria {
	c.Year = &year
	return c
}

// WithPagination sets the pagination
func (c *RepresentativeMonthlyGoalCriteria) WithPagination(page, pageSize int) *RepresentativeMonthlyGoalCriteria {
	c.Pagination = valueobject.NewPagination(page, pageSize)
	return c
}

// WithSort sets the sort field and order
func (c *RepresentativeMonthlyGoalCriteria) WithSort(sortBy string, direction valueobject.SortDirection) *RepresentativeMonthlyGoalCriteria {
	sortOrder, _ := valueobject.NewSortOrder(sortBy, direction)
	c.SortBy = sortBy
	c.SortOrder = sortOrder
	return c
}
