package request

import (
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/constants"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// CreateTaskRequest representa a requisição para criar uma tarefa
type CreateTaskRequest struct {
	Title       string                 `json:"title" validate:"required"`
	Description string                 `json:"description"`
	StartDate   *string                `json:"start_date"`
	Category    constants.TaskCategory `json:"category" validate:"required"`
	Priority    constants.TaskPriority `json:"priority" validate:"required"`
	Status      constants.TaskStatus   `json:"status"`
	DueDate     *string                `json:"due_date"`
	AssigneeID  *string                `json:"assignee_id"`
	FlowID      *string                `json:"flow_id"`
	Archived    bool                   `json:"archived"`
}

// UpdateTaskRequest representa a requisição para atualizar uma tarefa
type UpdateTaskRequest struct {
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	StartDate   *string                 `json:"start_date"`
	Category    *constants.TaskCategory `json:"category"`
	Priority    *constants.TaskPriority `json:"priority"`
	Status      *constants.TaskStatus   `json:"status"`
	DueDate     *string                 `json:"due_date"`
	AssigneeID  *string                 `json:"assignee_id"`
	FlowID      *string                 `json:"flow_id"`
	Archived    *bool                   `json:"archived"`
}

// ListTasksRequest representa a requisição para listar tarefas
type ListTasksRequest struct {
	Page       int                     `json:"page"`
	Limit      int                     `json:"limit"`
	Status     *constants.TaskStatus   `json:"status"`
	Category   *constants.TaskCategory `json:"category"`
	Priority   *constants.TaskPriority `json:"priority"`
	AssigneeID *string                 `json:"assignee_id"`
	FlowID     *string                 `json:"flow_id"`
	Archived   *bool                   `json:"archived"`
	StartDate  *string                 `json:"start_date"`
	EndDate    *string                 `json:"end_date"`
	SortBy     string                  `json:"sort_by"`
	SortOrder  string                  `json:"sort_order"`
}

// CreateSubtaskRequest representa a requisição para criar uma subtarefa
type CreateSubtaskRequest struct {
	TaskID      string                 `json:"task_id" validate:"required"`
	Title       string                 `json:"title" validate:"required"`
	Description string                 `json:"description"`
	Priority    constants.TaskPriority `json:"priority"`
	Status      constants.TaskStatus   `json:"status"`
	DueDate     *string                `json:"due_date"`
	AssigneeID  *string                `json:"assignee_id"`
}

// UpdateSubtaskRequest representa a requisição para atualizar uma subtarefa
type UpdateSubtaskRequest struct {
	Title       *string                 `json:"title"`
	Description *string                 `json:"description"`
	Priority    *constants.TaskPriority `json:"priority"`
	Status      *constants.TaskStatus   `json:"status"`
	DueDate     *string                 `json:"due_date"`
	AssigneeID  *string                 `json:"assignee_id"`
}

// CreateCommentRequest representa a requisição para criar um comentário
type CreateCommentRequest struct {
	TaskID  string `json:"task_id" validate:"required"`
	UserID  string `json:"user_id"` // ID do usuário que está criando o comentário
	Content string `json:"content" validate:"required"`
	Text    string `json:"text"` // Campo adicional para compatibilidade com frontend
}

// UpdateCommentRequest representa a requisição para atualizar um comentário
type UpdateCommentRequest struct {
	Content string `json:"content" validate:"required"`
}

// CreateNotificationRequest representa a requisição para criar uma notificação
type CreateNotificationRequest struct {
	UserID           string `json:"user_id" validate:"required"`
	Title            string `json:"title" validate:"required"`
	Message          string `json:"message" validate:"required"`
	NotificationType string `json:"notification_type" validate:"required"`
	TaskID           string `json:"task_id"`
}

// UpdateNotificationRequest representa a requisição para atualizar uma notificação
type UpdateNotificationRequest struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

// TaskCriteriaRequest representa os critérios de busca para tarefas
type TaskCriteriaRequest struct {
	Status     *constants.TaskStatus   `json:"status"`
	Category   *constants.TaskCategory `json:"category"`
	Priority   *constants.TaskPriority `json:"priority"`
	AssigneeID *string                 `json:"assignee_id"`
	FlowID     *string                 `json:"flow_id"`
	Archived   *bool                   `json:"archived"`
	StartDate  *string                 `json:"start_date"`
	EndDate    *string                 `json:"end_date"`
}

// SortOrderRequest representa a ordenação
type SortOrderRequest struct {
	Field     string `json:"field" validate:"required"`
	Direction string `json:"direction" validate:"required,oneof=ASC DESC"`
}

// PaginationRequest representa a paginação
type PaginationRequest struct {
	Page  int `json:"page" validate:"min=0"`
	Limit int `json:"limit" validate:"min=0"`
}

// ToCriteria converte a requisição para TaskCriteria
func (r *TaskCriteriaRequest) ToCriteria() (*domain.TaskCriteria, error) {
	criteria := domain.NewTaskCriteria()
	criteria.WithStatus(r.Status)
	criteria.WithCategory(r.Category)
	criteria.WithPriority(r.Priority)
	criteria.WithAssigneeID(r.AssigneeID)
	criteria.WithFlowID(r.FlowID)
	criteria.WithArchived(r.Archived)
	criteria.WithStartDate(r.StartDate)
	criteria.WithEndDate(r.EndDate)
	return criteria, nil
}

// ToSortOrder converte a requisição para SortOrder
func (r *SortOrderRequest) ToSortOrder() (*valueobject.SortOrder, error) {
	return valueobject.NewSortOrder(r.Field, valueobject.SortDirection(r.Direction))
}

// ToPagination converte a requisição para Pagination
func (r *PaginationRequest) ToPagination() valueobject.Pagination {
	return valueobject.NewPagination(r.Page, r.Limit)
}

// ReorderTasksRequest representa a requisição para reordenar tarefas
type ReorderTasksRequest struct {
	TaskIDs []string `json:"task_ids" validate:"required"`
}
