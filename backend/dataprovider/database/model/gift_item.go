package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GiftItemModel struct {
	UUID      uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string         `gorm:"not null;size:200;uniqueIndex:idx_name_deleted_at"`
	Stock     int            `gorm:"not null;default:0"`
	Price     float64        `gorm:"not null;default:0"`
	CreatedAt time.Time      `gorm:"not null;index:idx_created_at"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_deleted_at"`
}

// Índices compostos para performance:
// - idx_name_deleted_at: (name, deleted_at) para queries filtradas por nome
// - idx_created_at: (created_at) para queries por data de criação
// - idx_deleted_at: (deleted_at) para soft deletes

func (GiftItemModel) TableName() string {
	return "gift_items"
}
