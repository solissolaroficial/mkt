package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GiftTransactionModel struct {
	UUID            uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	ItemName        string         `gorm:"not null;size:200;index:idx_item_name_deleted_at"`
	Quantity        int            `gorm:"not null"`
	TransactionType string         `gorm:"not null;size:10;index:idx_transaction_type_deleted_at"`
	Date            time.Time      `gorm:"not null;index:idx_date_deleted_at"`
	Time            string         `gorm:"not null;size:5"` // Formato HH:MM
	Price           *float64       `gorm:"type:decimal(10,2)"`
	Representative  *string        `gorm:"size:200"`
	Unit            string         `gorm:"not null;size:20;default:'unid.'"`
	CreatedAt       time.Time      `gorm:"not null;index:idx_created_at"`
	UpdatedAt       time.Time      `gorm:"not null"`
	DeletedAt       gorm.DeletedAt `gorm:"index:idx_deleted_at"`
}

// Índices compostos para performance:
// - idx_item_name_deleted_at: (item_name, deleted_at) para queries filtradas por item
// - idx_transaction_type_deleted_at: (transaction_type, deleted_at) para queries por tipo
// - idx_date_deleted_at: (date, deleted_at) para queries por data
// - idx_created_at: (created_at) para queries por data de criação
// - idx_deleted_at: (deleted_at) para soft deletes

func (GiftTransactionModel) TableName() string {
	return "gift_transactions"
}
