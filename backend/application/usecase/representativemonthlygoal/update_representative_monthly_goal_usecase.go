package representativemonthlygoal

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// UpdateRepresentativeMonthlyGoalInput defines input data for updating a monthly goal
type UpdateRepresentativeMonthlyGoalInput struct {
	ID       uuid.UUID
	Target   *float64
	Realized *float64
}

// UpdateRepresentativeMonthlyGoalOutput defines output data after updating a monthly goal
type UpdateRepresentativeMonthlyGoalOutput struct {
	ID                 uuid.UUID
	RepresentativeUUID uuid.UUID
	Month              int
	Year               int
	Target             float64
	Realized           float64
	Percentage         float64
	UpdatedAt          string
}

type UpdateRepresentativeMonthlyGoalUseCase struct {
	monthlyGoalGateway gateway.RepresentativeMonthlyGoalGateway
}

func NewUpdateRepresentativeMonthlyGoalUseCase(
	monthlyGoalGateway gateway.RepresentativeMonthlyGoalGateway,
) *UpdateRepresentativeMonthlyGoalUseCase {
	return &UpdateRepresentativeMonthlyGoalUseCase{
		monthlyGoalGateway: monthlyGoalGateway,
	}
}

func (uc *UpdateRepresentativeMonthlyGoalUseCase) Execute(ctx context.Context, input UpdateRepresentativeMonthlyGoalInput) (*UpdateRepresentativeMonthlyGoalOutput, error) {
	// Get existing goal
	goal, err := uc.monthlyGoalGateway.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if input.Target != nil {
		if err := goal.UpdateTarget(*input.Target); err != nil {
			return nil, err
		}
	}

	if input.Realized != nil {
		if err := goal.UpdateRealized(*input.Realized); err != nil {
			return nil, err
		}
	}

	// Save to database
	if err := uc.monthlyGoalGateway.Update(goal); err != nil {
		return nil, err
	}

	return &UpdateRepresentativeMonthlyGoalOutput{
		ID:                 goal.ID(),
		RepresentativeUUID: goal.RepresentativeUUID(),
		Month:              goal.Month(),
		Year:               goal.Year(),
		Target:             goal.Target(),
		Realized:           goal.Realized(),
		Percentage:         goal.PercentageAchieved(),
		UpdatedAt:          goal.UpdatedAt().Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
