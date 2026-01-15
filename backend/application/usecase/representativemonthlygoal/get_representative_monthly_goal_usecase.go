package representativemonthlygoal

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetRepresentativeMonthlyGoalInput defines input data for getting a monthly goal
type GetRepresentativeMonthlyGoalInput struct {
	ID uuid.UUID
}

// GetRepresentativeMonthlyGoalOutput defines output data for a monthly goal
type GetRepresentativeMonthlyGoalOutput struct {
	ID               uuid.UUID
	RepresentativeID uuid.UUID
	Month            int
	Year             int
	Target           float64
	Realized         float64
	Percentage       float64
	Remaining        float64
	IsTargetMet      bool
	CreatedAt        string
	UpdatedAt        string
}

type GetRepresentativeMonthlyGoalUseCase struct {
	monthlyGoalGateway gateway.RepresentativeMonthlyGoalGateway
}

func NewGetRepresentativeMonthlyGoalUseCase(
	monthlyGoalGateway gateway.RepresentativeMonthlyGoalGateway,
) *GetRepresentativeMonthlyGoalUseCase {
	return &GetRepresentativeMonthlyGoalUseCase{
		monthlyGoalGateway: monthlyGoalGateway,
	}
}

func (uc *GetRepresentativeMonthlyGoalUseCase) Execute(ctx context.Context, input GetRepresentativeMonthlyGoalInput) (*GetRepresentativeMonthlyGoalOutput, error) {
	goal, err := uc.monthlyGoalGateway.GetByID(input.ID)
	if err != nil {
		return nil, err
	}

	return &GetRepresentativeMonthlyGoalOutput{
		ID:               goal.ID(),
		RepresentativeID: goal.RepresentativeID(),
		Month:            goal.Month(),
		Year:             goal.Year(),
		Target:           goal.Target(),
		Realized:         goal.Realized(),
		Percentage:       goal.PercentageAchieved(),
		Remaining:        goal.Remaining(),
		IsTargetMet:      goal.IsTargetMet(),
		CreatedAt:        goal.CreatedAt().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:        goal.UpdatedAt().Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}
