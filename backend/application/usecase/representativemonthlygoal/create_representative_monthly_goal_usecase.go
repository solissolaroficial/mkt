package representativemonthlygoal

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// CreateRepresentativeMonthlyGoalInput defines input data for creating a monthly goal
type CreateRepresentativeMonthlyGoalInput struct {
	RepresentativeUUID uuid.UUID
	Month              int
	Year               int
	Target             float64
}

// CreateRepresentativeMonthlyGoalOutput defines output data after creating a monthly goal
type CreateRepresentativeMonthlyGoalOutput struct {
	ID                 uuid.UUID
	RepresentativeUUID uuid.UUID
	Month              int
	Year               int
	Target             float64
	Realized           float64
	Percentage         float64
	CreatedAt          string
	UpdatedAt          string
}

type CreateRepresentativeMonthlyGoalUseCase struct {
	monthlyGoalGateway    gateway.RepresentativeMonthlyGoalGateway
	representativeGateway gateway.RepresentativeGateway
}

func NewCreateRepresentativeMonthlyGoalUseCase(
	monthlyGoalGateway gateway.RepresentativeMonthlyGoalGateway,
	representativeGateway gateway.RepresentativeGateway,
) *CreateRepresentativeMonthlyGoalUseCase {
	return &CreateRepresentativeMonthlyGoalUseCase{
		monthlyGoalGateway:    monthlyGoalGateway,
		representativeGateway: representativeGateway,
	}
}

func (uc *CreateRepresentativeMonthlyGoalUseCase) Execute(ctx context.Context, input CreateRepresentativeMonthlyGoalInput) (*CreateRepresentativeMonthlyGoalOutput, error) {
	// Check if representative exists
	_, err := uc.representativeGateway.FindByID(input.RepresentativeUUID)
	if err != nil {
		if err == errors.ErrRepresentativeNotFound {
			return nil, errors.ErrRepresentativeNotFound
		}
		return nil, err
	}

	// Check if goal already exists for this representative and month/year
	existing, err := uc.monthlyGoalGateway.GetByRepresentativeAndMonth(input.RepresentativeUUID, input.Month, input.Year)
	if err == nil && existing != nil {
		return nil, errors.ErrRepresentativeAlreadyExists
	}

	// Create entity
	goal, err := entity.NewRepresentativeMonthlyGoal(
		input.RepresentativeUUID,
		input.Month,
		input.Year,
		input.Target,
	)
	if err != nil {
		return nil, err
	}

	// Save to database
	if err := uc.monthlyGoalGateway.Create(goal); err != nil {
		return nil, err
	}

	return &CreateRepresentativeMonthlyGoalOutput{
		ID:                 goal.ID(),
		RepresentativeUUID: goal.RepresentativeUUID(),
		Month:              goal.Month(),
		Year:               goal.Year(),
		Target:             goal.Target(),
		Realized:           goal.Realized(),
		Percentage:         goal.PercentageAchieved(),
		CreatedAt:          goal.CreatedAt().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:          goal.UpdatedAt().Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
