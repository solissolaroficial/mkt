package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RepMarketingActionModel struct {
	UUID               uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	RepresentativeUUID uuid.UUID      `gorm:"not null;type:uuid;index:idx_representative_uuid;constraint:fk_rep_marketing_actions_representative,foreignKey:RepresentativeUUID,references:UUID,onDelete:CASCADE,onUpdate:CASCADE"`
	Date               time.Time      `gorm:"not null;index:idx_date"`
	Description        string         `gorm:"not null;size:500"`
	Month              string         `gorm:"not null;size:10;index:idx_month"`
	CreatedAt          time.Time      `gorm:"not null;index:idx_created_at"`
	UpdatedAt          time.Time      `gorm:"not null"`
	DeletedAt          gorm.DeletedAt `gorm:"index:idx_deleted_at"`
}

func (RepMarketingActionModel) TableName() string {
	return "rep_marketing_actions"
}
