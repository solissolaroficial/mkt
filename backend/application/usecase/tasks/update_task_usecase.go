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
	taskGateway gateway.TaskGateway
}

func NewUpdateTaskUseCase(taskGateway gateway.TaskGateway) *UpdateTaskUseCase {
	return &UpdateTaskUseCase{
		taskGateway: taskGateway,
	}
}

func (uc *UpdateTaskUseCase) Execute(id string, title string, description *string, category *string, priority *string, status *string, assigneeID *string, archived *bool, flows []string, startDate *string, dueDate *string) (*entity.Task, error) {
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

	if flows != nil {
		task.SetFlows(flows)
	}

	// Validate task
	if err := task.Validate(); err != nil {
		return nil, err
	}

	// Save task
	if err := uc.taskGateway.Update(task); err != nil {
		return nil, err
	}

	return task, nil
}
