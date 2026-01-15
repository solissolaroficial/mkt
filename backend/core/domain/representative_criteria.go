package domain

import (
	"regexp"

	"github.com/seu-usuario/solis-backend/core/domain/errors"
)

// RepresentativeCriteria defines filter options for querying representatives
type RepresentativeCriteria struct {
	name    *string
	company *string
	email   *string
	region  *string
	city    *string
	active  *bool
	code    *int
}

// NewRepresentativeCriteria creates a new RepresentativeCriteria with default values
func NewRepresentativeCriteria() *RepresentativeCriteria {
	return &RepresentativeCriteria{}
}

// WithName filters by representative name (partial match)
func (rc *RepresentativeCriteria) WithName(name string) *RepresentativeCriteria {
	rc.name = &name
	return rc
}

// WithCompany filters by company name (partial match)
func (rc *RepresentativeCriteria) WithCompany(company string) *RepresentativeCriteria {
	rc.company = &company
	return rc
}

// WithEmail filters by email (partial match)
func (rc *RepresentativeCriteria) WithEmail(email string) *RepresentativeCriteria {
	rc.email = &email
	return rc
}

// WithRegion filters by region (exact match)
func (rc *RepresentativeCriteria) WithRegion(region string) *RepresentativeCriteria {
	rc.region = &region
	return rc
}

// WithCity filters by city (partial match)
func (rc *RepresentativeCriteria) WithCity(city string) *RepresentativeCriteria {
	rc.city = &city
	return rc
}

// WithActive filters by active status
func (rc *RepresentativeCriteria) WithActive(active bool) *RepresentativeCriteria {
	rc.active = &active
	return rc
}

// WithCode filters by representative code (exact match)
func (rc *RepresentativeCriteria) WithCode(code int) *RepresentativeCriteria {
	rc.code = &code
	return rc
}

// GetName returns the name filter
func (rc *RepresentativeCriteria) GetName() *string {
	return rc.name
}

// GetCompany returns the company filter
func (rc *RepresentativeCriteria) GetCompany() *string {
	return rc.company
}

// GetEmail returns the email filter
func (rc *RepresentativeCriteria) GetEmail() *string {
	return rc.email
}

// GetRegion returns the region filter
func (rc *RepresentativeCriteria) GetRegion() *string {
	return rc.region
}

// GetCity returns the city filter
func (rc *RepresentativeCriteria) GetCity() *string {
	return rc.city
}

// GetActive returns the active filter
func (rc *RepresentativeCriteria) GetActive() *bool {
	return rc.active
}

// GetCode returns the code filter
func (rc *RepresentativeCriteria) GetCode() *int {
	return rc.code
}

// Validate checks if the criteria is valid
func (rc *RepresentativeCriteria) Validate() error {
	// Validate email format if provided
	if rc.email != nil {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(*rc.email) {
			return errors.ErrInvalidRepresentativeEmail
		}
	}

	// Validate code range if provided
	if rc.code != nil {
		if *rc.code < 100 || *rc.code > 999 {
			return errors.ErrInvalidRepresentativeCode
		}
	}

	return nil
}
