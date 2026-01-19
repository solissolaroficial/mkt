package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// BudgetGroup representa um grupo de orçamento (ex: "Marketing", "Vendas", "Operacional")
type BudgetGroup struct {
	id        uuid.UUID
	code      *valueobject.BudgetCode
	name      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

// NewBudgetGroup cria uma nova entidade BudgetGroup
func NewBudgetGroup(code string, name string) (*BudgetGroup, error) {
	// Validar código
	if code == "" {
		return nil, errors.New("code is required")
	}

	codeVO, err := valueobject.NewBudgetCode(code)
	if err != nil {
		return nil, err
	}

	// Validar nome
	if name == "" {
		return nil, errors.New("name is required")
	}

	now := time.Now()

	return &BudgetGroup{
		id:        uuid.New(),
		code:      codeVO,
		name:      name,
		createdAt: now,
		updatedAt: now,
		deletedAt: nil,
	}, nil
}

// ReconstructBudgetGroup reconstrói a entidade do banco de dados
func ReconstructBudgetGroup(id uuid.UUID, code string, name string, createdAt time.Time, updatedAt time.Time, deletedAt *time.Time) *BudgetGroup {
	return &BudgetGroup{
		id:        id,
		code:      valueobject.ReconstructBudgetCode(code),
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}
}

// Getters
func (b *BudgetGroup) ID() uuid.UUID {
	return b.id
}

func (b *BudgetGroup) Code() string {
	if b.code == nil {
		return ""
	}
	return b.code.String()
}

func (b *BudgetGroup) Name() string {
	return b.name
}

func (b *BudgetGroup) CreatedAt() time.Time {
	return b.createdAt
}

func (b *BudgetGroup) UpdatedAt() time.Time {
	return b.updatedAt
}

func (b *BudgetGroup) DeletedAt() *time.Time {
	return b.deletedAt
}

// Métodos de Negócio

// Validate valida os dados da entidade
func (b *BudgetGroup) Validate() error {
	if b.code == nil || b.code.String() == "" {
		return errors.New("code is required")
	}
	if b.name == "" {
		return errors.New("name is required")
	}
	return nil
}

// UpdateName atualiza o nome do grupo
func (b *BudgetGroup) UpdateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	b.name = name
	b.updatedAt = time.Now()
	return nil
}

// UpdateCode atualiza o código do grupo
func (b *BudgetGroup) UpdateCode(code string) error {
	if code == "" {
		return errors.New("code is required")
	}

	codeVO, err := valueobject.NewBudgetCode(code)
	if err != nil {
		return err
	}

	b.code = codeVO
	b.updatedAt = time.Now()
	return nil
}

// SoftDelete marca a entidade como deletada
func (b *BudgetGroup) SoftDelete() {
	now := time.Now()
	b.deletedAt = &now
	b.updatedAt = now
}

// IsActive retorna true se a entidade não foi deletada
func (b *BudgetGroup) IsActive() bool {
	return b.deletedAt == nil
}
