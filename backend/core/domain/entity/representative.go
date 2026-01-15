package entity

import (
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// Representative represents a sales representative
type Representative struct {
	uuid      uuid.UUID
	code      *valueobject.RepresentativeCode
	name      string
	email     string
	phone     string
	company   string
	region    string
	city      string
	attendant string
	active    bool
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

// NewRepresentative creates a new Representative with validation
func NewRepresentative(
	code *valueobject.RepresentativeCode,
	name string,
	email string,
	phone string,
	company string,
	region string,
	city string,
	attendant string,
) (*Representative, error) {
	if err := validateRepresentativeData(code, name, email, company, region, city); err != nil {
		return nil, err
	}

	now := time.Now()
	return &Representative{
		uuid:      uuid.New(),
		code:      code,
		name:      name,
		email:     email,
		phone:     phone,
		company:   company,
		region:    region,
		city:      city,
		attendant: attendant,
		active:    true,
		createdAt: now,
		updatedAt: now,
	}, nil
}

// ReconstructRepresentative reconstructs a Representative from persisted data
func ReconstructRepresentative(
	uuid uuid.UUID,
	code *valueobject.RepresentativeCode,
	name string,
	email string,
	phone string,
	company string,
	region string,
	city string,
	attendant string,
	active bool,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *Representative {
	return &Representative{
		uuid:      uuid,
		code:      code,
		name:      name,
		email:     email,
		phone:     phone,
		company:   company,
		region:    region,
		city:      city,
		attendant: attendant,
		active:    active,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}
}

// Getters

func (r *Representative) UUID() uuid.UUID {
	return r.uuid
}

func (r *Representative) Code() *valueobject.RepresentativeCode {
	return r.code
}

func (r *Representative) Name() string {
	return r.name
}

func (r *Representative) Email() string {
	return r.email
}

func (r *Representative) Company() string {
	return r.company
}

func (r *Representative) Region() string {
	return r.region
}

func (r *Representative) City() string {
	return r.city
}

func (r *Representative) Phone() string {
	return r.phone
}

func (r *Representative) Attendant() string {
	return r.attendant
}

func (r *Representative) Active() bool {
	return r.active
}

func (r *Representative) CreatedAt() time.Time {
	return r.createdAt
}

func (r *Representative) UpdatedAt() time.Time {
	return r.updatedAt
}

func (r *Representative) DeletedAt() *time.Time {
	return r.deletedAt
}

// Setters (business logic)

// UpdateName updates the representative's name
func (r *Representative) UpdateName(name string) error {
	if name == "" {
		return errors.ErrRepresentativeNameRequired
	}
	r.name = name
	r.updatedAt = time.Now()
	return nil
}

// UpdateEmail updates the representative's email
func (r *Representative) UpdateEmail(email string) error {
	if email == "" {
		return errors.ErrRepresentativeEmailRequired
	}
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.ErrInvalidRepresentativeEmail
	}
	r.email = email
	r.updatedAt = time.Now()
	return nil
}

// UpdateCompany updates the representative's company
func (r *Representative) UpdateCompany(company string) error {
	if company == "" {
		return errors.ErrRepresentativeCompanyRequired
	}
	r.company = company
	r.updatedAt = time.Now()
	return nil
}

// UpdateRegion updates the representative's region
func (r *Representative) UpdateRegion(region string) error {
	if region == "" {
		return errors.ErrRepresentativeRegionRequired
	}
	r.region = region
	r.updatedAt = time.Now()
	return nil
}

// UpdateCity updates the representative's city
func (r *Representative) UpdateCity(city string) error {
	if city == "" {
		return errors.ErrRepresentativeCityRequired
	}
	r.city = city
	r.updatedAt = time.Now()
	return nil
}

// UpdatePhone updates the representative's phone
func (r *Representative) UpdatePhone(phone string) error {
	r.phone = phone
	r.updatedAt = time.Now()
	return nil
}

// UpdateAttendant updates the representative's attendant
func (r *Representative) UpdateAttendant(attendant string) error {
	r.attendant = attendant
	r.updatedAt = time.Now()
	return nil
}

// Activate activates the representative
func (r *Representative) Activate() {
	r.active = true
	r.updatedAt = time.Now()
}

// Deactivate deactivates the representative
func (r *Representative) Deactivate() {
	r.active = false
	r.updatedAt = time.Now()
}

// SoftDelete marks the representative as deleted
func (r *Representative) SoftDelete() {
	now := time.Now()
	r.deletedAt = &now
	r.updatedAt = now
}

// IsActive checks if the representative is active (not deleted)
func (r *Representative) IsActive() bool {
	return r.deletedAt == nil
}

// Validate validates the representative data
func (r *Representative) Validate() error {
	return validateRepresentativeData(r.code, r.name, r.email, r.company, r.region, r.city)
}

// validateRepresentativeData validates representative data
func validateRepresentativeData(
	code *valueobject.RepresentativeCode,
	name string,
	email string,
	company string,
	region string,
	city string,
) error {
	if code == nil {
		return errors.ErrRepresentativeCodeRequired
	}

	if name == "" {
		return errors.ErrRepresentativeNameRequired
	}

	if email == "" {
		return errors.ErrRepresentativeEmailRequired
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.ErrInvalidRepresentativeEmail
	}

	if company == "" {
		return errors.ErrRepresentativeCompanyRequired
	}

	if region == "" {
		return errors.ErrRepresentativeRegionRequired
	}

	return nil
}
