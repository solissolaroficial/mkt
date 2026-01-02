package tasks

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// CreateSubtaskUseCase cria uma nova subtarefa
type CreateSubtaskUseCase struct {
	subtaskGateway gateway.SubtaskGateway
}

// NewCreateSubtaskUseCase cria um novo CreateSubtaskUseCase
func NewCreateSubtaskUseCase(subtaskGateway gateway.SubtaskGateway) *CreateSubtaskUseCase {
	return &CreateSubtaskUseCase{
		subtaskGateway: subtaskGateway,
	}
}

// Execute cria uma nova subtarefa
func (uc *CreateSubtaskUseCase) Execute(subtask *entity.Subtask) error {
	if err := subtask.Validate(); err != nil {
		return err
	}

	return uc.subtaskGateway.Create(subtask)
}
