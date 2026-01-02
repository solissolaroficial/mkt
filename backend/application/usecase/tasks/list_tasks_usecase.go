package tasks

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

type ListTasksUseCase struct {
	taskGateway gateway.TaskGateway
}

func NewListTasksUseCase(taskGateway gateway.TaskGateway) *ListTasksUseCase {
	return &ListTasksUseCase{
		taskGateway: taskGateway,
	}
}

func (uc *ListTasksUseCase) Execute(
	statuses []string,
	categories []string,
	priorities []string,
	assigneeID *string,
	flow *string,
	archived *bool,
	flows []string,
	dateFrom *string,
	dateTo *string,
	page int,
	pageSize int,
	sortBy string,
	sortOrder string,
) ([]*entity.Task, int64, error) {
	// Build criteria
	criteria := gateway.NewTaskCriteria()
	if len(statuses) > 0 {
		criteria = criteria.WithStatuses(statuses)
	}
	if len(categories) > 0 {
		criteria = criteria.WithCategories(categories)
	}
	if len(priorities) > 0 {
		criteria = criteria.WithPriorities(priorities)
	}
	if len(flows) > 0 {
		criteria = criteria.WithFlows(flows)
	}
	if archived != nil {
		criteria = criteria.WithArchived(*archived)
	}
	if dateFrom != nil && dateTo != nil {
		// Parse date strings to time.Time
		parsedFrom, err := time.Parse(time.RFC3339, *dateFrom)
		if err != nil {
			return nil, 0, err
		}
		parsedTo, err := time.Parse(time.RFC3339, *dateTo)
		if err != nil {
			return nil, 0, err
		}
		criteria = criteria.WithDueDateRange(parsedFrom, parsedTo)
	}

	// Build pagination
	pagination := valueobject.NewPagination(page, pageSize)

	// Find tasks (without sort order for now)
	tasks, total, err := uc.taskGateway.FindByCriteria(criteria, &pagination, nil)
	if err != nil {
		return nil, 0, err
	}

	return tasks, total, nil
}
