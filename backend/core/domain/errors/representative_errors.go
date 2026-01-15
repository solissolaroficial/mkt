package errors

import "errors"

var (
	// ErrRepresentativeNotFound is returned when a representative is not found
	ErrRepresentativeNotFound = errors.New("representative not found")

	// ErrRepresentativeAlreadyExists is returned when trying to create a representative with existing code
	ErrRepresentativeAlreadyExists = errors.New("representative with this code already exists")

	// ErrInvalidRepresentativeData is returned when representative data is invalid
	ErrInvalidRepresentativeData = errors.New("invalid representative data")

	// ErrInvalidRepresentativeEmail is returned when email format is invalid
	ErrInvalidRepresentativeEmail = errors.New("invalid email format")

	// ErrInvalidRepresentativeCode is returned when code is out of valid range
	ErrInvalidRepresentativeCode = errors.New("representative code must be between 100 and 999")

	// ErrRepresentativeCodeRequired is returned when code is not provided
	ErrRepresentativeCodeRequired = errors.New("representative code is required")

	// ErrRepresentativeNameRequired is returned when name is not provided
	ErrRepresentativeNameRequired = errors.New("representative name is required")

	// ErrRepresentativeEmailRequired is returned when email is not provided
	ErrRepresentativeEmailRequired = errors.New("representative email is required")

	// ErrRepresentativeCompanyRequired is returned when company is not provided
	ErrRepresentativeCompanyRequired = errors.New("representative company is required")

	// ErrRepresentativeRegionRequired is returned when region is not provided
	ErrRepresentativeRegionRequired = errors.New("representative region is required")

	// ErrRepresentativeCityRequired is returned when city is not provided
	ErrRepresentativeCityRequired = errors.New("representative city is required")

	// ErrInvalidRepresentativeMonth is returned when month is out of valid range (1-12)
	ErrInvalidRepresentativeMonth = errors.New("representative month must be between 1 and 12")

	// ErrInvalidRepresentativeYear is returned when year is out of valid range (2000-2100)
	ErrInvalidRepresentativeYear = errors.New("representative year must be between 2000 and 2100")

	// ErrInvalidRepresentativeGoal is returned when goal value is invalid (negative)
	ErrInvalidRepresentativeGoal = errors.New("representative goal must be non-negative")

	// ErrRepresentativeGoalNotFound is returned when a monthly goal is not found
	ErrRepresentativeGoalNotFound = errors.New("representative monthly goal not found")
)
