package tasks

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// DeleteSubtaskUseCase deleta uma subtarefa
type DeleteSubtaskUseCase struct {
	subtaskGateway gateway.SubtaskGateway
}

// NewDeleteSubtaskUseCase cria um novo DeleteSubtaskUseCase
func NewDeleteSubtaskUseCase(subtaskGateway gateway.SubtaskGateway) *DeleteSubtaskUseCase {
	return &DeleteSubtaskUseCase{
		subtaskGateway: subtaskGateway,
	}
}

// Execute deleta uma subtarefa
func (uc *DeleteSubtaskUseCase) Execute(id string) error {
	// Parse ID
	subtaskID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid subtask ID: %w", err)
	}

	// Delete subtask
	return uc.subtaskGateway.Delete(subtaskID)
}
