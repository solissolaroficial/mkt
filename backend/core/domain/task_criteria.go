package domain

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/constants"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// TaskCriteria representa os critérios de busca para tarefas
type TaskCriteria struct {
	status     *constants.TaskStatus
	category   *constants.TaskCategory
	priority   *constants.TaskPriority
	assigneeID *string
	flowID     *string
	archived   *bool
	startDate  *time.Time
	endDate    *time.Time
	pagination *valueobject.Pagination
	sortOrders []*valueobject.SortOrder
}

// NewTaskCriteria cria um novo TaskCriteria
func NewTaskCriteria() *TaskCriteria {
	return &TaskCriteria{}
}

// WithStatus define o status da tarefa
func (c *TaskCriteria) WithStatus(status *constants.TaskStatus) *TaskCriteria {
	c.status = status
	return c
}

// WithCategory define a categoria da tarefa
func (c *TaskCriteria) WithCategory(category *constants.TaskCategory) *TaskCriteria {
	c.category = category
	return c
}

// WithPriority define a prioridade da tarefa
func (c *TaskCriteria) WithPriority(priority *constants.TaskPriority) *TaskCriteria {
	c.priority = priority
	return c
}

// WithAssigneeID define o ID do responsável
func (c *TaskCriteria) WithAssigneeID(assigneeID *string) *TaskCriteria {
	c.assigneeID = assigneeID
	return c
}

// WithFlowID define o ID do fluxo da tarefa
func (c *TaskCriteria) WithFlowID(flowID *string) *TaskCriteria {
	c.flowID = flowID
	return c
}

// WithArchived define se a tarefa está arquivada
func (c *TaskCriteria) WithArchived(archived *bool) *TaskCriteria {
	c.archived = archived
	return c
}

// WithStartDate define a data de início
func (c *TaskCriteria) WithStartDate(startDate *string) *TaskCriteria {
	if startDate != nil {
		parsedTime, err := time.Parse(time.RFC3339, *startDate)
		if err == nil {
			c.startDate = &parsedTime
		}
	}
	return c
}

// WithEndDate define a data de fim
func (c *TaskCriteria) WithEndDate(endDate *string) *TaskCriteria {
	if endDate != nil {
		parsedTime, err := time.Parse(time.RFC3339, *endDate)
		if err == nil {
			c.endDate = &parsedTime
		}
	}
	return c
}

// WithPagination define a paginação
func (c *TaskCriteria) WithPagination(pagination *valueobject.Pagination) *TaskCriteria {
	c.pagination = pagination
	return c
}

// WithSortOrders define as ordenações
func (c *TaskCriteria) WithSortOrders(sortOrders []*valueobject.SortOrder) *TaskCriteria {
	c.sortOrders = sortOrders
	return c
}

// GetStatus retorna o status
func (c *TaskCriteria) GetStatus() *constants.TaskStatus {
	return c.status
}

// GetCategory retorna a categoria
func (c *TaskCriteria) GetCategory() *constants.TaskCategory {
	return c.category
}

// GetPriority retorna a prioridade
func (c *TaskCriteria) GetPriority() *constants.TaskPriority {
	return c.priority
}

// GetAssigneeID retorna o ID do responsável
func (c *TaskCriteria) GetAssigneeID() *string {
	return c.assigneeID
}

// GetFlowID retorna o ID do fluxo
func (c *TaskCriteria) GetFlowID() *string {
	return c.flowID
}

// GetArchived retorna se está arquivada
func (c *TaskCriteria) GetArchived() *bool {
	return c.archived
}

// GetStartDate retorna a data de início
func (c *TaskCriteria) GetStartDate() *time.Time {
	return c.startDate
}

// GetEndDate retorna a data de fim
func (c *TaskCriteria) GetEndDate() *time.Time {
	return c.endDate
}

// GetPagination retorna a paginação
func (c *TaskCriteria) GetPagination() *valueobject.Pagination {
	return c.pagination
}

// GetSortOrders retorna as ordenações
func (c *TaskCriteria) GetSortOrders() []*valueobject.SortOrder {
	return c.sortOrders
}
