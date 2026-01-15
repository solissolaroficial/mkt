package model

import (
	"time"

	"github.com/google/uuid"
)

// RepresentativeMonthlyGoalModel represents a monthly goal for a representative
type RepresentativeMonthlyGoalModel struct {
	ID               uuid.UUID  `gorm:"type:uuid;primaryKey;column:id" json:"id"`
	RepresentativeID uuid.UUID  `gorm:"type:uuid;column:representative_id;not null;index" json:"representativeId"`
	Month            int        `gorm:"column:month;not null" json:"month"`
	Year             int        `gorm:"column:year;not null;index" json:"year"`
	Target           float64    `gorm:"column:target;not null" json:"target"`
	Realized         float64    `gorm:"column:realized;not null;default:0" json:"realized"`
	CreatedAt        time.Time  `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt        time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt        *time.Time `gorm:"column:deleted_at;index" json:"deletedAt,omitempty"`
}

// TableName specifies the table name for GORM
func (RepresentativeMonthlyGoalModel) TableName() string {
	return "representative_monthly_goals"
}
