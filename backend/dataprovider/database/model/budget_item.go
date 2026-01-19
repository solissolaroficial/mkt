package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// BudgetItemModel representa a tabela budget_items no banco de dados
type BudgetItemModel struct {
	UUID         uuid.UUID          `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ObjectUUID   *uuid.UUID         `gorm:"type:uuid;index:idx_object_uuid;constraint:fk_budget_items_object,foreignKey:ObjectUUID,references:UUID,onDelete:RESTRICT,onUpdate:CASCADE"`
	Object       *BudgetObjectModel `gorm:"foreignKey:ObjectUUID"`
	GroupUUID    *uuid.UUID         `gorm:"type:uuid;index:idx_group_uuid;constraint:fk_budget_items_group,foreignKey:GroupUUID,references:UUID,onDelete:RESTRICT,onUpdate:CASCADE"`
	Group        *BudgetGroupModel  `gorm:"foreignKey:GroupUUID"`
	Cod          string             `gorm:"not null;size:20;index:idx_cod"`
	Desc         string             `gorm:"not null;size:500;index:idx_desc"`
	Vals         datatypes.JSON     `gorm:"type:jsonb;not null"` // Array de 12 floats
	RealizedVals datatypes.JSON     `gorm:"type:jsonb;not null"` // Array de 12 floats
	Year         int                `gorm:"not null;index:idx_year;default:2025"`
	CreatedAt    time.Time          `gorm:"not null;index:idx_created_at"`
	UpdatedAt    time.Time          `gorm:"not null"`
	DeletedAt    gorm.DeletedAt     `gorm:"index:idx_deleted_at"`
}

// TableName define o nome da tabela
func (BudgetItemModel) TableName() string {
	return "budget_items"
}
