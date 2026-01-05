package gateway

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"

	"github.com/google/uuid"
)

// TaskGateway define a interface para persistência de tarefas
type TaskGateway interface {
	// Create cria uma nova tarefa
	Create(task *entity.Task) error

	// Update atualiza uma tarefa existente
	Update(task *entity.Task) error

	// Delete deleta uma tarefa (soft delete)
	Delete(id uuid.UUID) error

	// FindByID busca uma tarefa pelo ID
	FindByID(id uuid.UUID) (*entity.Task, error)

	// FindAll busca todas as tarefas com paginação
	FindAll(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error)

	// FindByCriteria busca tarefas usando o padrão Criteria
	FindByCriteria(criteria *TaskCriteria, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error)

	// CountByCriteria conta tarefas usando o padrão Criteria
	CountByCriteria(criteria *TaskCriteria) (int64, error)

	// FindByAssigneeID busca tarefas por assignee
	FindByAssigneeID(assigneeID uuid.UUID, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error)

	// FindByStatus busca tarefas por status
	FindByStatus(status string, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error)

	// FindByCategory busca tarefas por categoria
	FindByCategory(category string, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error)

	// FindByFlow busca tarefas por fluxo
	FindByFlow(flow string, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error)

	// FindArchived busca tarefas arquivadas
	FindArchived(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error)

	// FindOverdue busca tarefas atrasadas
	FindOverdue(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Task, int64, error)

	// Reorder atualiza a ordem de múltiplas tarefas
	Reorder(taskIDs []uuid.UUID) error
}

// TaskCriteria define os critérios de busca para tarefas
type TaskCriteria struct {
	Statuses    []string
	Categories  []string
	Priorities  []string
	Flows       []string
	AssigneeIDs []uuid.UUID
	Archived    *bool
	DueDateFrom *time.Time
	DueDateTo   *time.Time
	Title       *string
}

// NewTaskCriteria cria um novo TaskCriteria vazio
func NewTaskCriteria() *TaskCriteria {
	return &TaskCriteria{}
}

// WithStatuses adiciona filtros de status
func (c *TaskCriteria) WithStatuses(statuses []string) *TaskCriteria {
	c.Statuses = statuses
	return c
}

// WithCategories adiciona filtros de categoria
func (c *TaskCriteria) WithCategories(categories []string) *TaskCriteria {
	c.Categories = categories
	return c
}

// WithPriorities adiciona filtros de prioridade
func (c *TaskCriteria) WithPriorities(priorities []string) *TaskCriteria {
	c.Priorities = priorities
	return c
}

// WithFlows adiciona filtros de fluxo
func (c *TaskCriteria) WithFlows(flows []string) *TaskCriteria {
	c.Flows = flows
	return c
}

// WithAssigneeIDs adiciona filtros de assignee
func (c *TaskCriteria) WithAssigneeIDs(assigneeIDs []uuid.UUID) *TaskCriteria {
	c.AssigneeIDs = assigneeIDs
	return c
}

// WithArchived adiciona filtro de arquivado
func (c *TaskCriteria) WithArchived(archived bool) *TaskCriteria {
	c.Archived = &archived
	return c
}

// WithDueDateRange adiciona filtro de intervalo de data de vencimento
func (c *TaskCriteria) WithDueDateRange(from, to time.Time) *TaskCriteria {
	c.DueDateFrom = &from
	c.DueDateTo = &to
	return c
}

// WithTitle adiciona filtro de título (busca parcial)
func (c *TaskCriteria) WithTitle(title string) *TaskCriteria {
	c.Title = &title
	return c
}
