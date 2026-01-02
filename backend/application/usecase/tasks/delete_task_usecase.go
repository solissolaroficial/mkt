package tasks

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type DeleteTaskUseCase struct {
	taskGateway gateway.TaskGateway
}

func NewDeleteTaskUseCase(taskGateway gateway.TaskGateway) *DeleteTaskUseCase {
	return &DeleteTaskUseCase{
		taskGateway: taskGateway,
	}
}

func (uc *DeleteTaskUseCase) Execute(id string) error {
	// Parse ID
	taskID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid task ID: %w", err)
	}

	// Delete task
	return uc.taskGateway.Delete(taskID)
}
