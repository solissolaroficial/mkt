package tasks

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type GetTaskUseCase struct {
	taskGateway gateway.TaskGateway
}

func NewGetTaskUseCase(taskGateway gateway.TaskGateway) *GetTaskUseCase {
	return &GetTaskUseCase{
		taskGateway: taskGateway,
	}
}

func (uc *GetTaskUseCase) Execute(id string) (*entity.Task, error) {
	// Parse ID
	taskID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID: %w", err)
	}

	// Find task
	task, err := uc.taskGateway.FindByID(taskID)
	if err != nil {
		return nil, err
	}

	return task, nil
}
