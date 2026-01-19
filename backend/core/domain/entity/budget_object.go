package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// BudgetObject representa um objeto de orçamento (ex: "Marketing Digital", "Eventos")
type BudgetObject struct {
	id        uuid.UUID
	code      *valueobject.BudgetCode
	name      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

// NewBudgetObject cria uma nova entidade BudgetObject
func NewBudgetObject(code string, name string) (*BudgetObject, error) {
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

	return &BudgetObject{
		id:        uuid.New(),
		code:      codeVO,
		name:      name,
		createdAt: now,
		updatedAt: now,
		deletedAt: nil,
	}, nil
}

// ReconstructBudgetObject reconstrói a entidade do banco de dados
func ReconstructBudgetObject(id uuid.UUID, code string, name string, createdAt time.Time, updatedAt time.Time, deletedAt *time.Time) *BudgetObject {
	return &BudgetObject{
		id:        id,
		code:      valueobject.ReconstructBudgetCode(code),
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}
}

// Getters
func (b *BudgetObject) ID() uuid.UUID {
	return b.id
}

func (b *BudgetObject) Code() string {
	if b.code == nil {
		return ""
	}
	return b.code.String()
}

func (b *BudgetObject) Name() string {
	return b.name
}

func (b *BudgetObject) CreatedAt() time.Time {
	return b.createdAt
}

func (b *BudgetObject) UpdatedAt() time.Time {
	return b.updatedAt
}

func (b *BudgetObject) DeletedAt() *time.Time {
	return b.deletedAt
}

// Métodos de Negócio

// Validate valida os dados da entidade
func (b *BudgetObject) Validate() error {
	if b.code == nil || b.code.String() == "" {
		return errors.New("code is required")
	}
	if b.name == "" {
		return errors.New("name is required")
	}
	return nil
}

// UpdateName atualiza o nome do objeto
func (b *BudgetObject) UpdateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	b.name = name
	b.updatedAt = time.Now()
	return nil
}

// UpdateCode atualiza o código do objeto
func (b *BudgetObject) UpdateCode(code string) error {
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
func (b *BudgetObject) SoftDelete() {
	now := time.Now()
	b.deletedAt = &now
	b.updatedAt = now
}

// IsActive retorna true se a entidade não foi deletada
func (b *BudgetObject) IsActive() bool {
	return b.deletedAt == nil
}
