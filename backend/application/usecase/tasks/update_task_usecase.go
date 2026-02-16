package tasks

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/constants"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

type UpdateTaskUseCase struct {
	taskGateway         gateway.TaskGateway
	notificationGateway gateway.NotificationGateway
}

func NewUpdateTaskUseCase(
	taskGateway gateway.TaskGateway,
	notificationGateway gateway.NotificationGateway,
) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{
		taskGateway:         taskGateway,
		notificationGateway: notificationGateway,
	}
}

func (uc *UpdateTaskUseCase) Execute(id string, title string, description *string, category *string, priority *string, status *string, assigneeID *string, flowID *string, archived *bool, startDate *string, dueDate *string) (*entity.Task, error) {
	// Parse ID
	taskID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid task ID: %w", err)
	}

	// Find existing task
	task, err := uc.taskGateway.FindByID(taskID)
	if err != nil {
		return nil, err
	}

	// Guardar valores anteriores
	oldAssigneeID := task.AssigneeUUID()
	oldStatus := task.Status()

	// Update task fields using setters
	if title != "" {
		if err := task.UpdateTitle(title); err != nil {
			return nil, err
		}
	}

	if description != nil {
		task.UpdateDescription(description)
	}

	if startDate != nil {
		parsedTime, err := time.Parse(time.RFC3339, *startDate)
		if err != nil {
			return nil, fmt.Errorf("invalid start date: %w", err)
		}
		task.SetStartDate(&parsedTime)
	}

	if dueDate != nil {
		parsedTime, err := time.Parse(time.RFC3339, *dueDate)
		if err != nil {
			return nil, fmt.Errorf("invalid due date: %w", err)
		}
		task.SetDueDate(parsedTime)
	}

	if category != nil {
		if err := task.SetCategory(constants.TaskCategory(*category)); err != nil {
			return nil, err
		}
	}

	if priority != nil {
		if err := task.SetPriority(constants.TaskPriority(*priority)); err != nil {
			return nil, err
		}
	}

	if status != nil {
		if err := task.SetStatus(constants.TaskStatus(*status)); err != nil {
			return nil, err
		}
	}

	if archived != nil {
		if *archived {
			task.Archive()
		} else {
			task.Unarchive()
		}
	}

	if assigneeID != nil {
		assigneeUUID, err := uuid.Parse(*assigneeID)
		if err != nil {
			return nil, fmt.Errorf("invalid assignee ID: %w", err)
		}
		task.SetAssigneeUUID(&assigneeUUID)
	}

	if flowID != nil {
		flowUUID, err := uuid.Parse(*flowID)
		if err != nil {
			return nil, fmt.Errorf("invalid flow ID: %w", err)
		}
		task.SetFlowUUID(&flowUUID)
	}

	// Validate task
	if err := task.Validate(); err != nil {
		return nil, err
	}

	// Save task
	if err := uc.taskGateway.Update(task); err != nil {
		return nil, err
	}

	// Verificar se assignee foi alterado e criar notificação (assíncrono)
	if assigneeID != nil && oldAssigneeID != nil {
		newAssigneeUUID, err := uuid.Parse(*assigneeID)
		if err == nil && *oldAssigneeID != newAssigneeUUID {
			go func() {
				if err := uc.createTaskAssignedNotification(task); err != nil {
					fmt.Printf("Erro ao criar notificação: %v\n", err)
				}
			}()
		}
	}

	// Verificar se status mudou para completed e criar notificação (assíncrono)
	if status != nil && oldStatus != constants.TaskStatusCompleted &&
		constants.TaskStatus(*status) == constants.TaskStatusCompleted {
		go func() {
			if err := uc.createTaskCompletedNotification(task); err != nil {
				fmt.Printf("Erro ao criar notificação: %v\n", err)
			}
		}()
	}

	return task, nil
}

// createTaskAssignedNotification cria uma notificação quando uma tarefa é atribuída a um usuário
func (uc *UpdateTaskUseCase) createTaskAssignedNotification(task *entity.Task) error {
	if task.AssigneeUUID() == nil {
		return nil
	}

	taskID := task.ID()
	notification, err := entity.NewNotification(
		*task.AssigneeUUID(),
		&taskID,
		constants.NotificationTypeTaskAssigned,
		"Tarefa atribuída a você",
		fmt.Sprintf("Você foi atribuído à tarefa: %s", task.Title()),
	)
	if err != nil {
		return err
	}

	return uc.notificationGateway.Create(notification)
}

// createTaskCompletedNotification cria uma notificação quando uma tarefa é concluída
func (uc *UpdateTaskUseCase) createTaskCompletedNotification(task *entity.Task) error {
	if task.AssigneeUUID() == nil {
		return nil
	}

	taskID := task.ID()
	notification, err := entity.NewNotification(
		*task.AssigneeUUID(),
		&taskID,
		constants.NotificationTypeTaskCompleted,
		"Tarefa concluída",
		fmt.Sprintf("A tarefa foi concluída: %s", task.Title()),
	)
	if err != nil {
		return err
	}

	return uc.notificationGateway.Create(notification)
}
