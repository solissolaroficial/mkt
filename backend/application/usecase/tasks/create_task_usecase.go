package tasks

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// CreateTaskUseCase cria uma nova tarefa
type CreateTaskUseCase struct {
	taskGateway    gateway.TaskGateway
	subtaskGateway gateway.SubtaskGateway
}

// NewCreateTaskUseCase cria um novo CreateTaskUseCase
func NewCreateTaskUseCase(taskGateway gateway.TaskGateway, subtaskGateway gateway.SubtaskGateway) *CreateTaskUseCase {
	return &CreateTaskUseCase{
		taskGateway:    taskGateway,
		subtaskGateway: subtaskGateway,
	}
}

// Execute cria uma nova tarefa
func (uc *CreateTaskUseCase) Execute(task *entity.Task) error {
	if err := task.Validate(); err != nil {
		return err
	}

	return uc.taskGateway.Create(task)
}

// ExecuteWithSubtasks cria uma nova tarefa com subtarefas
func (uc *CreateTaskUseCase) ExecuteWithSubtasks(task *entity.Task, subtasks []*entity.Subtask) error {
	if err := task.Validate(); err != nil {
		return err
	}

	if err := uc.taskGateway.Create(task); err != nil {
		return err
	}

	for _, subtask := range subtasks {
		if err := subtask.Validate(); err != nil {
			return err
		}
		if err := uc.subtaskGateway.Create(subtask); err != nil {
			return err
		}
	}

	return nil
}
