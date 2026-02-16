package entity

import (
	"time"

	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"

	"github.com/google/uuid"
)

// Flow representa um fluxo (Kanban independente) do sistema
type Flow struct {
	id          uuid.UUID
	name        string
	description *string
	color       *string
	sortOrder   int
	createdAt   time.Time
	updatedAt   time.Time
	deletedAt   *time.Time
}

// NewFlow cria um novo fluxo com validação
func NewFlow(name string, description *string, color *string, sortOrder int) (*Flow, error) {
	if name == "" {
		return nil, &domainerrors.FlowEmptyNameError{}
	}

	if sortOrder < 0 {
		return nil, &domainerrors.FlowInvalidSortOrderError{SortOrder: sortOrder}
	}

	now := time.Now()
	return &Flow{
		id:          uuid.New(),
		name:        name,
		description: description,
		color:       color,
		sortOrder:   sortOrder,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

// ReconstructFlow reconstrói um fluxo a partir de dados persistidos
func ReconstructFlow(
	id uuid.UUID,
	name string,
	description *string,
	color *string,
	sortOrder int,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *Flow {
	return &Flow{
		id:          id,
		name:        name,
		description: description,
		color:       color,
		sortOrder:   sortOrder,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
		deletedAt:   deletedAt,
	}
}

// Getters

func (f *Flow) ID() uuid.UUID {
	return f.id
}

func (f *Flow) Name() string {
	return f.name
}

func (f *Flow) Description() *string {
	return f.description
}

func (f *Flow) Color() *string {
	return f.color
}

func (f *Flow) SortOrder() int {
	return f.sortOrder
}

func (f *Flow) CreatedAt() time.Time {
	return f.createdAt
}

func (f *Flow) UpdatedAt() time.Time {
	return f.updatedAt
}

func (f *Flow) DeletedAt() *time.Time {
	return f.deletedAt
}

// Setters (business logic)

// UpdateName atualiza o nome do fluxo
func (f *Flow) UpdateName(name string) error {
	if name == "" {
		return &domainerrors.FlowEmptyNameError{}
	}
	f.name = name
	f.updatedAt = time.Now()
	return nil
}

// UpdateDescription atualiza a descrição do fluxo
func (f *Flow) UpdateDescription(description *string) {
	f.description = description
	f.updatedAt = time.Now()
}

// UpdateColor atualiza a cor do fluxo
func (f *Flow) UpdateColor(color *string) {
	f.color = color
	f.updatedAt = time.Now()
}

// UpdateSortOrder atualiza a ordem do fluxo
func (f *Flow) UpdateSortOrder(sortOrder int) error {
	if sortOrder < 0 {
		return &domainerrors.FlowInvalidSortOrderError{SortOrder: sortOrder}
	}
	f.sortOrder = sortOrder
	f.updatedAt = time.Now()
	return nil
}

// SoftDelete marca o fluxo como deletado
func (f *Flow) SoftDelete() {
	now := time.Now()
	f.deletedAt = &now
	f.updatedAt = now
}

// IsActive verifica se o fluxo está ativo (não foi deletado)
func (f *Flow) IsActive() bool {
	return f.deletedAt == nil
}

// Validate valida os dados do fluxo
func (f *Flow) Validate() error {
	if f.name == "" {
		return &domainerrors.FlowEmptyNameError{}
	}
	if f.sortOrder < 0 {
		return &domainerrors.FlowInvalidSortOrderError{SortOrder: f.sortOrder}
	}
	return nil
}
