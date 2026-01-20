package tasks

import (
	"fmt"

	"github.com/seu-usuario/solis-backend/core/domain/constants"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// CreateTaskUseCase cria uma nova tarefa
type CreateTaskUseCase struct {
	taskGateway         gateway.TaskGateway
	subtaskGateway      gateway.SubtaskGateway
	notificationGateway gateway.NotificationGateway
}

// NewCreateTaskUseCase cria um novo CreateTaskUseCase
func NewCreateTaskUseCase(
	taskGateway gateway.TaskGateway,
	subtaskGateway gateway.SubtaskGateway,
	notificationGateway gateway.NotificationGateway,
) *CreateTaskUseCase {
	return &CreateTaskUseCase{
		taskGateway:         taskGateway,
		subtaskGateway:      subtaskGateway,
		notificationGateway: notificationGateway,
	}
}

// Execute cria uma nova tarefa
func (uc *CreateTaskUseCase) Execute(task *entity.Task) error {
	if err := task.Validate(); err != nil {
		return err
	}

	if err := uc.taskGateway.Create(task); err != nil {
		return err
	}

	// Criar notificação para o assignee (assíncrono)
	go func() {
		if err := uc.createTaskAssignedNotification(task); err != nil {
			// Logar erro mas não falhar a operação principal
			fmt.Printf("Erro ao criar notificação: %v\n", err)
		}
	}()

	return nil
}

// ExecuteWithSubtasks cria uma nova tarefa com subtarefas
func (uc *CreateTaskUseCase) ExecuteWithSubtasks(task *entity.Task, subtasks []*entity.Subtask) error {
	if err := task.Validate(); err != nil {
		return err
	}

	if err := uc.taskGateway.Create(task); err != nil {
		return err
	}

	// Criar notificação para o assignee (assíncrono)
	go func() {
		if err := uc.createTaskAssignedNotification(task); err != nil {
			fmt.Printf("Erro ao criar notificação: %v\n", err)
		}
	}()

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

// createTaskAssignedNotification cria uma notificação quando uma tarefa é atribuída a um usuário
func (uc *CreateTaskUseCase) createTaskAssignedNotification(task *entity.Task) error {
	if task.AssigneeUUID() == nil {
		return nil // Sem assignee, sem notificação
	}

	taskID := task.ID()
	notification, err := entity.NewNotification(
		*task.AssigneeUUID(),
		&taskID,
		constants.NotificationTypeTaskAssigned,
		"Nova tarefa atribuída",
		fmt.Sprintf("Você foi atribuído à tarefa: %s", task.Title()),
	)
	if err != nil {
		return err
	}

	return uc.notificationGateway.Create(notification)
}
