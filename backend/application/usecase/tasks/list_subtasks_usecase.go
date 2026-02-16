package tasks

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// ListSubtasksUseCase lista subtarefas
type ListSubtasksUseCase struct {
	subtaskGateway gateway.SubtaskGateway
}

// NewListSubtasksUseCase cria um novo ListSubtasksUseCase
func NewListSubtasksUseCase(subtaskGateway gateway.SubtaskGateway) *ListSubtasksUseCase {
	return &ListSubtasksUseCase{
		subtaskGateway: subtaskGateway,
	}
}

// Execute lista subtarefas
func (uc *ListSubtasksUseCase) Execute(page, pageSize int) ([]*entity.Subtask, int64, error) {
	pagination := valueobject.NewPagination(page, pageSize)
	subtasks, total, err := uc.subtaskGateway.FindAll(&pagination, nil)
	return subtasks, total, err
}

// ExecuteByTaskID lista subtarefas filtradas por task ID
func (uc *ListSubtasksUseCase) ExecuteByTaskID(taskID string, page, pageSize int) ([]*entity.Subtask, int64, error) {
	taskUUID, err := uuid.Parse(taskID)
	if err != nil {
		return nil, 0, err
	}

	pagination := valueobject.NewPagination(page, pageSize)
	subtasks, total, err := uc.subtaskGateway.FindByTaskID(taskUUID, &pagination, nil)
	return subtasks, total, err
}
