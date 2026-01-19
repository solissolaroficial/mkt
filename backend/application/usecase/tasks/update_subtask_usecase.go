package tasks

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// UpdateSubtaskUseCase atualiza uma subtarefa existente
type UpdateSubtaskUseCase struct {
	subtaskGateway gateway.SubtaskGateway
}

// NewUpdateSubtaskUseCase cria um novo UpdateSubtaskUseCase
func NewUpdateSubtaskUseCase(subtaskGateway gateway.SubtaskGateway) *UpdateSubtaskUseCase {
	return &UpdateSubtaskUseCase{
		subtaskGateway: subtaskGateway,
	}
}

// Execute atualiza uma subtarefa
func (uc *UpdateSubtaskUseCase) Execute(id, title string, completed *bool, assigneeID *string, dueDate *string) (*entity.Subtask, error) {
	// Parse ID
	subtaskID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid subtask ID: %w", err)
	}

	// Find existing subtask
	subtask, err := uc.subtaskGateway.FindByID(subtaskID)
	if err != nil {
		return nil, err
	}

	// Update subtask fields using setters
	if title != "" {
		if err := subtask.UpdateTitle(title); err != nil {
			return nil, err
		}
	}

	if completed != nil {
		subtask.SetCompleted(*completed)
	}

	if assigneeID != nil {
		assigneeUUID, err := uuid.Parse(*assigneeID)
		if err != nil {
			return nil, fmt.Errorf("invalid assignee ID: %w", err)
		}
		subtask.SetAssigneeUUID(&assigneeUUID)
	}

	// Save subtask
	if err := uc.subtaskGateway.Update(subtask); err != nil {
		return nil, err
	}

	return subtask, nil
}
