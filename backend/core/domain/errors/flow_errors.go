package errors

import "errors"

// Flow related error variables
var (
	// ErrFlowNotFound is returned when a flow is not found in the system
	ErrFlowNotFound = errors.New("flow not found")

	// ErrFlowEmptyName is returned when flow name is empty
	ErrFlowEmptyName = errors.New("flow name is required")

	// ErrFlowInvalidSortOrder is returned when sort order is invalid
	ErrFlowInvalidSortOrder = errors.New("invalid sort order")
)

// FlowNotFoundError represents an error when a flow is not found
type FlowNotFoundError struct{}

func (e *FlowNotFoundError) Error() string {
	return ErrFlowNotFound.Error()
}

// FlowEmptyNameError represents an error when flow name is empty
type FlowEmptyNameError struct{}

func (e *FlowEmptyNameError) Error() string {
	return ErrFlowEmptyName.Error()
}

// FlowInvalidSortOrderError represents an error when sort order is invalid
type FlowInvalidSortOrderError struct {
	SortOrder int
}

func (e *FlowInvalidSortOrderError) Error() string {
	return ErrFlowInvalidSortOrder.Error()
}
