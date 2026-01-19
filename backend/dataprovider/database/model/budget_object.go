package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BudgetObjectModel representa a tabela budget_objects no banco de dados
type BudgetObjectModel struct {
	UUID      uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Code      string         `gorm:"not null;size:20;uniqueIndex:idx_budget_object_code"`
	Name      string         `gorm:"not null;size:200;index:idx_budget_object_name"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_budget_object_deleted_at"`
}

// TableName define o nome da tabela
func (BudgetObjectModel) TableName() string {
	return "budget_objects"
}
