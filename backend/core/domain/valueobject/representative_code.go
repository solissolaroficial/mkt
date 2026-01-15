package valueobject

import (
	"errors"
)

// RepresentativeCode represents a representative's unique code
type RepresentativeCode struct {
	value int
}

// NewRepresentativeCode creates a new RepresentativeCode with validation
func NewRepresentativeCode(code int) (*RepresentativeCode, error) {
	if code < 100 || code > 999 {
		return nil, errors.New("representative code must be between 100 and 999")
	}
	return &RepresentativeCode{value: code}, nil
}

// ReconstructRepresentativeCode creates a RepresentativeCode without validation (for data from DB)
func ReconstructRepresentativeCode(code int) *RepresentativeCode {
	return &RepresentativeCode{value: code}
}

// Value returns the code value
func (rc *RepresentativeCode) Value() int {
	return rc.value
}

// Equals checks if two RepresentativeCode are equal
func (rc *RepresentativeCode) Equals(other *RepresentativeCode) bool {
	if rc == nil || other == nil {
		return false
	}
	return rc.value == other.value
}
