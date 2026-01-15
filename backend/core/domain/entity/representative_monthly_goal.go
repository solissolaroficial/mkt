package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
)

// RepresentativeMonthlyGoal represents a monthly goal for a representative
type RepresentativeMonthlyGoal struct {
	id               uuid.UUID
	representativeID uuid.UUID
	month            int
	year             int
	target           float64
	realized         float64
	createdAt        time.Time
	updatedAt        time.Time
	deletedAt        *time.Time
}

// NewRepresentativeMonthlyGoal creates a new RepresentativeMonthlyGoal
func NewRepresentativeMonthlyGoal(
	representativeID uuid.UUID,
	month int,
	year int,
	target float64,
) (*RepresentativeMonthlyGoal, error) {
	if month < 1 || month > 12 {
		return nil, errors.ErrInvalidRepresentativeMonth
	}
	if year < 2000 || year > 2100 {
		return nil, errors.ErrInvalidRepresentativeYear
	}
	if target < 0 {
		return nil, errors.ErrInvalidRepresentativeGoal
	}

	now := time.Now()
	return &RepresentativeMonthlyGoal{
		id:               uuid.New(),
		representativeID: representativeID,
		month:            month,
		year:             year,
		target:           target,
		realized:         0,
		createdAt:        now,
		updatedAt:        now,
		deletedAt:        nil,
	}, nil
}

// ReconstructRepresentativeMonthlyGoal reconstructs a RepresentativeMonthlyGoal from persistence
func ReconstructRepresentativeMonthlyGoal(
	id uuid.UUID,
	representativeID uuid.UUID,
	month int,
	year int,
	target float64,
	realized float64,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) (*RepresentativeMonthlyGoal, error) {
	if month < 1 || month > 12 {
		return nil, errors.ErrInvalidRepresentativeMonth
	}
	if year < 2000 || year > 2100 {
		return nil, errors.ErrInvalidRepresentativeYear
	}
	if target < 0 {
		return nil, errors.ErrInvalidRepresentativeGoal
	}
	if realized < 0 {
		return nil, errors.ErrInvalidRepresentativeGoal
	}

	return &RepresentativeMonthlyGoal{
		id:               id,
		representativeID: representativeID,
		month:            month,
		year:             year,
		target:           target,
		realized:         realized,
		createdAt:        createdAt,
		updatedAt:        updatedAt,
		deletedAt:        deletedAt,
	}, nil
}

// Getters

func (rmg *RepresentativeMonthlyGoal) ID() uuid.UUID {
	return rmg.id
}

func (rmg *RepresentativeMonthlyGoal) RepresentativeID() uuid.UUID {
	return rmg.representativeID
}

func (rmg *RepresentativeMonthlyGoal) Month() int {
	return rmg.month
}

func (rmg *RepresentativeMonthlyGoal) Year() int {
	return rmg.year
}

func (rmg *RepresentativeMonthlyGoal) Target() float64 {
	return rmg.target
}

func (rmg *RepresentativeMonthlyGoal) Realized() float64 {
	return rmg.realized
}

func (rmg *RepresentativeMonthlyGoal) CreatedAt() time.Time {
	return rmg.createdAt
}

func (rmg *RepresentativeMonthlyGoal) UpdatedAt() time.Time {
	return rmg.updatedAt
}

func (rmg *RepresentativeMonthlyGoal) DeletedAt() *time.Time {
	return rmg.deletedAt
}

// UpdateTarget updates the target value
func (rmg *RepresentativeMonthlyGoal) UpdateTarget(target float64) error {
	if target < 0 {
		return errors.ErrInvalidRepresentativeGoal
	}
	rmg.target = target
	rmg.updatedAt = time.Now()
	return nil
}

// UpdateRealized updates the realized value
func (rmg *RepresentativeMonthlyGoal) UpdateRealized(realized float64) error {
	if realized < 0 {
		return errors.ErrInvalidRepresentativeGoal
	}
	rmg.realized = realized
	rmg.updatedAt = time.Now()
	return nil
}

// IsTargetMet returns true if the realized value meets or exceeds the target
func (rmg *RepresentativeMonthlyGoal) IsTargetMet() bool {
	return rmg.realized >= rmg.target
}

// PercentageAchieved returns the percentage of the target achieved
func (rmg *RepresentativeMonthlyGoal) PercentageAchieved() float64 {
	if rmg.target == 0 {
		return 0
	}
	return (rmg.realized / rmg.target) * 100
}

// Remaining returns the remaining amount to reach the target
func (rmg *RepresentativeMonthlyGoal) Remaining() float64 {
	remaining := rmg.target - rmg.realized
	if remaining < 0 {
		return 0
	}
	return remaining
}

// MarkAsDeleted marks the goal as deleted
func (rmg *RepresentativeMonthlyGoal) MarkAsDeleted() {
	now := time.Now()
	rmg.deletedAt = &now
	rmg.updatedAt = now
}

// IsDeleted returns true if the goal is marked as deleted
func (rmg *RepresentativeMonthlyGoal) IsDeleted() bool {
	return rmg.deletedAt != nil
}
