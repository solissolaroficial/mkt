package entity

import (
	"time"

	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"

	"github.com/google/uuid"
)

// Subtask representa uma subtarefa de uma tarefa
type Subtask struct {
	id           uuid.UUID
	taskUUID     uuid.UUID
	title        string
	completed    bool
	assigneeUUID *uuid.UUID
	dueDate      *time.Time
	createdAt    time.Time
	updatedAt    time.Time
	deletedAt    *time.Time
}

// NewSubtask cria uma nova subtarefa
func NewSubtask(
	taskUUID uuid.UUID,
	title string,
) (*Subtask, error) {
	if title == "" {
		return nil, &domainerrors.SubtaskEmptyTitleError{}
	}

	now := time.Now()
	return &Subtask{
		id:        uuid.New(),
		taskUUID:  taskUUID,
		title:     title,
		completed: false,
		createdAt: now,
		updatedAt: now,
	}, nil
}

// ReconstructSubtask reconstrói uma subtarefa a partir de dados persistidos
func ReconstructSubtask(
	id uuid.UUID,
	taskUUID uuid.UUID,
	title string,
	completed bool,
	assigneeUUID *uuid.UUID,
	dueDate *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *Subtask {
	return &Subtask{
		id:           id,
		taskUUID:     taskUUID,
		title:        title,
		completed:    completed,
		assigneeUUID: assigneeUUID,
		dueDate:      dueDate,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
		deletedAt:    deletedAt,
	}
}

// Getters

func (s *Subtask) ID() uuid.UUID {
	return s.id
}

func (s *Subtask) TaskUUID() uuid.UUID {
	return s.taskUUID
}

func (s *Subtask) Title() string {
	return s.title
}

func (s *Subtask) Completed() bool {
	return s.completed
}

func (s *Subtask) AssigneeUUID() *uuid.UUID {
	return s.assigneeUUID
}

func (s *Subtask) DueDate() *time.Time {
	return s.dueDate
}

func (s *Subtask) CreatedAt() time.Time {
	return s.createdAt
}

func (s *Subtask) UpdatedAt() time.Time {
	return s.updatedAt
}

func (s *Subtask) DeletedAt() *time.Time {
	return s.deletedAt
}

// Setters (business logic)

// UpdateTitle atualiza o título da subtarefa
func (s *Subtask) UpdateTitle(title string) error {
	if title == "" {
		return &domainerrors.SubtaskEmptyTitleError{}
	}
	s.title = title
	s.updatedAt = time.Now()
	return nil
}

// ToggleCompleted alterna o status de conclusão da subtarefa
func (s *Subtask) ToggleCompleted() {
	s.completed = !s.completed
	s.updatedAt = time.Now()
}

// SetCompleted define o status de conclusão da subtarefa
func (s *Subtask) SetCompleted(completed bool) {
	s.completed = completed
	s.updatedAt = time.Now()
}

// SetAssigneeUUID define o responsável pela subtarefa
func (s *Subtask) SetAssigneeUUID(assigneeUUID *uuid.UUID) {
	s.assigneeUUID = assigneeUUID
	s.updatedAt = time.Now()
}

// SetDueDate define a data de vencimento da subtarefa
func (s *Subtask) SetDueDate(dueDate *time.Time) {
	s.dueDate = dueDate
	s.updatedAt = time.Now()
}

// SoftDelete marca a subtarefa como deletada
func (s *Subtask) SoftDelete() {
	now := time.Now()
	s.deletedAt = &now
	s.updatedAt = now
}

// IsActive verifica se a subtarefa está ativa (não foi deletada)
func (s *Subtask) IsActive() bool {
	return s.deletedAt == nil
}

// Validate valida os dados da subtarefa
func (s *Subtask) Validate() error {
	if s.title == "" {
		return &domainerrors.SubtaskEmptyTitleError{}
	}
	if len(s.title) < 3 {
		return &domainerrors.SubtaskTitleTooShortError{}
	}
	if len(s.title) > 200 {
		return &domainerrors.SubtaskTitleTooLongError{}
	}

	return nil
}
