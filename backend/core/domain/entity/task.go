package entity

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/constants"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"

	"github.com/google/uuid"
)

// Task representa uma tarefa do sistema
type Task struct {
	id           uuid.UUID
	title        string
	description  *string
	startDate    *time.Time
	dueDate      time.Time
	status       constants.TaskStatus
	priority     constants.TaskPriority
	category     constants.TaskCategory
	assigneeUUID *uuid.UUID
	flowUUID     *uuid.UUID
	archived     bool
	sortOrder    int
	createdAt    time.Time
	updatedAt    time.Time
	deletedAt    *time.Time
}

// NewTask cria uma nova tarefa
func NewTask(
	title string,
	description *string,
	startDate *time.Time,
	dueDate time.Time,
	priority constants.TaskPriority,
	status constants.TaskStatus,
	category constants.TaskCategory,
	sortOrder int,
) (*Task, error) {
	if title == "" {
		return nil, &domainerrors.TaskEmptyTitleError{}
	}

	if !constants.IsValidTaskStatus(status) {
		return nil, &domainerrors.TaskInvalidStatusError{Status: string(status)}
	}

	if !constants.IsValidTaskPriority(priority) {
		return nil, &domainerrors.TaskInvalidPriorityError{Priority: string(priority)}
	}

	if !constants.IsValidTaskCategory(category) {
		return nil, &domainerrors.TaskInvalidCategoryError{Category: string(category)}
	}

	now := time.Now()
	return &Task{
		id:          uuid.New(),
		title:       title,
		description: description,
		startDate:   startDate,
		dueDate:     dueDate,
		priority:    priority,
		status:      status,
		category:    category,
		flowUUID:    nil,
		archived:    false,
		sortOrder:   sortOrder,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

// ReconstructTask reconstrói uma tarefa a partir de dados persistidos
func ReconstructTask(
	id uuid.UUID,
	title string,
	description *string,
	startDate *time.Time,
	dueDate time.Time,
	priority constants.TaskPriority,
	status constants.TaskStatus,
	category constants.TaskCategory,
	assigneeUUID *uuid.UUID,
	flowUUID *uuid.UUID,
	archived bool,
	sortOrder int,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *Task {
	return &Task{
		id:           id,
		title:        title,
		description:  description,
		startDate:    startDate,
		dueDate:      dueDate,
		priority:     priority,
		status:       status,
		category:     category,
		assigneeUUID: assigneeUUID,
		flowUUID:     flowUUID,
		archived:     archived,
		sortOrder:    sortOrder,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
		deletedAt:    deletedAt,
	}
}

// Getters

func (t *Task) ID() uuid.UUID {
	return t.id
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) Description() *string {
	return t.description
}

func (t *Task) StartDate() *time.Time {
	return t.startDate
}

func (t *Task) DueDate() time.Time {
	return t.dueDate
}

func (t *Task) Priority() constants.TaskPriority {
	return t.priority
}

func (t *Task) Status() constants.TaskStatus {
	return t.status
}

func (t *Task) Category() constants.TaskCategory {
	return t.category
}

func (t *Task) AssigneeUUID() *uuid.UUID {
	return t.assigneeUUID
}

func (t *Task) FlowUUID() *uuid.UUID {
	return t.flowUUID
}

func (t *Task) Archived() bool {
	return t.archived
}

func (t *Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *Task) DeletedAt() *time.Time {
	return t.deletedAt
}

func (t *Task) SortOrder() int {
	return t.sortOrder
}

// Setters (business logic)

// UpdateTitle atualiza o título da tarefa
func (t *Task) UpdateTitle(title string) error {
	if title == "" {
		return &domainerrors.TaskEmptyTitleError{}
	}
	t.title = title
	t.updatedAt = time.Now()
	return nil
}

// UpdateDescription atualiza a descrição da tarefa
func (t *Task) UpdateDescription(description *string) {
	t.description = description
	t.updatedAt = time.Now()
}

// SetStartDate atualiza a data de início da tarefa
func (t *Task) SetStartDate(startDate *time.Time) {
	t.startDate = startDate
	t.updatedAt = time.Now()
}

// SetDueDate atualiza a data de vencimento da tarefa
func (t *Task) SetDueDate(dueDate time.Time) {
	t.dueDate = dueDate
	t.updatedAt = time.Now()
}

// SetPriority atualiza a prioridade da tarefa
func (t *Task) SetPriority(priority constants.TaskPriority) error {
	if !constants.IsValidTaskPriority(priority) {
		return &domainerrors.TaskInvalidPriorityError{Priority: string(priority)}
	}
	t.priority = priority
	t.updatedAt = time.Now()
	return nil
}

// SetStatus atualiza o status da tarefa
func (t *Task) SetStatus(status constants.TaskStatus) error {
	if !constants.IsValidTaskStatus(status) {
		return &domainerrors.TaskInvalidStatusError{Status: string(status)}
	}
	t.status = status
	t.updatedAt = time.Now()
	return nil
}

// SetCategory atualiza a categoria da tarefa
func (t *Task) SetCategory(category constants.TaskCategory) error {
	if !constants.IsValidTaskCategory(category) {
		return &domainerrors.TaskInvalidCategoryError{Category: string(category)}
	}
	t.category = category
	t.updatedAt = time.Now()
	return nil
}

// SetAssigneeUUID atualiza o responsável pela tarefa
func (t *Task) SetAssigneeUUID(assigneeUUID *uuid.UUID) {
	t.assigneeUUID = assigneeUUID
	t.updatedAt = time.Now()
}

// SetFlowUUID atualiza o flow UUID da tarefa
func (t *Task) SetFlowUUID(flowUUID *uuid.UUID) {
	t.flowUUID = flowUUID
	t.updatedAt = time.Now()
}

// Archive arquiva a tarefa
func (t *Task) Archive() {
	t.archived = true
	t.updatedAt = time.Now()
}

// Unarchive desarquiva a tarefa
func (t *Task) Unarchive() {
	t.archived = false
	t.updatedAt = time.Now()
}

// SetSortOrder atualiza a ordem da tarefa
func (t *Task) SetSortOrder(order int) {
	t.sortOrder = order
	t.updatedAt = time.Now()
}

// SoftDelete marca a tarefa como deletada
func (t *Task) SoftDelete() {
	now := time.Now()
	t.deletedAt = &now
	t.updatedAt = now
}

// IsActive verifica se a tarefa está ativa (não foi deletada)
func (t *Task) IsActive() bool {
	return t.deletedAt == nil
}

// Validate valida os dados da tarefa
func (t *Task) Validate() error {
	if t.title == "" {
		return &domainerrors.TaskEmptyTitleError{}
	}

	if t.dueDate.IsZero() {
		return &domainerrors.TaskInvalidStatusError{Status: "due date is required"}
	}

	// Validar prioridade
	validPriorities := map[constants.TaskPriority]bool{
		constants.TaskPriorityLow:    true,
		constants.TaskPriorityMedium: true,
		constants.TaskPriorityHigh:   true,
		constants.TaskPriorityUrgent: true,
	}
	if !validPriorities[t.priority] {
		return &domainerrors.TaskInvalidPriorityError{Priority: string(t.priority)}
	}

	// Validar status
	validStatuses := map[constants.TaskStatus]bool{
		constants.TaskStatusPending:    true,
		constants.TaskStatusInProgress: true,
		constants.TaskStatusCompleted:  true,
		constants.TaskStatusCancelled:  true,
	}
	if !validStatuses[t.status] {
		return &domainerrors.TaskInvalidStatusError{Status: string(t.status)}
	}

	// Validar categoria
	validCategories := map[constants.TaskCategory]bool{
		constants.TaskCategoryCommercial:     true,
		constants.TaskCategoryMarketing:      true,
		constants.TaskCategoryTraining:       true,
		constants.TaskCategoryAdministrative: true,
		constants.TaskCategorySupport:        true,
	}
	if !validCategories[t.category] {
		return &domainerrors.TaskInvalidCategoryError{Category: string(t.category)}
	}

	return nil
}

// HasFlowUUID verifica se a tarefa pertence a um fluxo específico
func (t *Task) HasFlowUUID(flowUUID uuid.UUID) bool {
	if t.flowUUID == nil {
		return false
	}
	return *t.flowUUID == flowUUID
}

// IsOverdue verifica se a tarefa está atrasada
func (t *Task) IsOverdue() bool {
	if t.status == constants.TaskStatusCompleted {
		return false
	}
	return time.Now().After(t.dueDate)
}

// IsDueToday verifica se a tarefa vence hoje
func (t *Task) IsDueToday() bool {
	now := time.Now()
	return now.Year() == t.dueDate.Year() &&
		now.Month() == t.dueDate.Month() &&
		now.Day() == t.dueDate.Day()
}
