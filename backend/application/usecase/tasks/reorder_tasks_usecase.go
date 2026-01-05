package tasks

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// ReorderTasksUseCase reordena múltiplas tarefas
type ReorderTasksUseCase struct {
	taskGateway gateway.TaskGateway
}

// NewReorderTasksUseCase cria um novo ReorderTasksUseCase
func NewReorderTasksUseCase(taskGateway gateway.TaskGateway) *ReorderTasksUseCase {
	return &ReorderTasksUseCase{
		taskGateway: taskGateway,
	}
}

// Execute reordena as tarefas na ordem especificada
func (uc *ReorderTasksUseCase) Execute(taskIDs []uuid.UUID) error {
	return uc.taskGateway.Reorder(taskIDs)
}
