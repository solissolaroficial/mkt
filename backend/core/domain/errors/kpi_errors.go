package errors

import "errors"

// KPI related error variables
var (
	// ErrKpiNotFound is returned when KPI category is not found
	ErrKpiNotFound = errors.New("kpi not found")

	// ErrMonthDataNotFound is returned when monthly data is not found
	ErrMonthDataNotFound = errors.New("monthly data not found")

	// ErrInvalidMonth is returned when month value is invalid
	ErrInvalidMonth = errors.New("invalid month")

	// ErrInvalidUnit is returned when unit type is invalid
	ErrInvalidUnit = errors.New("invalid unit type")
)
