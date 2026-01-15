package representativemonthlygoal

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// DeleteRepresentativeMonthlyGoalInput defines input data for deleting a monthly goal
type DeleteRepresentativeMonthlyGoalInput struct {
	ID uuid.UUID
}

type DeleteRepresentativeMonthlyGoalUseCase struct {
	monthlyGoalGateway gateway.RepresentativeMonthlyGoalGateway
}

func NewDeleteRepresentativeMonthlyGoalUseCase(
	monthlyGoalGateway gateway.RepresentativeMonthlyGoalGateway,
) *DeleteRepresentativeMonthlyGoalUseCase {
	return &DeleteRepresentativeMonthlyGoalUseCase{
		monthlyGoalGateway: monthlyGoalGateway,
	}
}

func (uc *DeleteRepresentativeMonthlyGoalUseCase) Execute(ctx context.Context, input DeleteRepresentativeMonthlyGoalInput) error {
	// Check if goal exists
	_, err := uc.monthlyGoalGateway.GetByID(input.ID)
	if err != nil {
		return err
	}

	// Delete the goal
	return uc.monthlyGoalGateway.Delete(input.ID)
}
