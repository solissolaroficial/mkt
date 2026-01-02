package gateway

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"

	"github.com/google/uuid"
)

// SubtaskGateway define a interface para persistência de subtarefas
type SubtaskGateway interface {
	// Create cria uma nova subtarefa
	Create(subtask *entity.Subtask) error

	// Update atualiza uma subtarefa existente
	Update(subtask *entity.Subtask) error

	// Delete deleta uma subtarefa
	Delete(id uuid.UUID) error

	// FindByID busca uma subtarefa pelo ID
	FindByID(id uuid.UUID) (*entity.Subtask, error)

	// FindAll busca todas as subtarefas com paginação
	FindAll(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Subtask, int64, error)

	// FindByTaskID busca subtarefas por task ID
	FindByTaskID(taskID uuid.UUID, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Subtask, int64, error)

	// FindByAssigneeID busca subtarefas por assignee ID
	FindByAssigneeID(assigneeID uuid.UUID, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Subtask, int64, error)

	// FindCompleted busca subtarefas concluídas
	FindCompleted(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Subtask, int64, error)

	// FindPending busca subtarefas pendentes
	FindPending(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Subtask, int64, error)
}
