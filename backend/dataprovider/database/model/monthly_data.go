package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type MonthlyData struct {
	UUID          uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"uuid"`
	KpiCategoryID uuid.UUID      `gorm:"type:uuid;not null;index" json:"kpi_category_id"`
	Month         string         `gorm:"not null;size:3" json:"month"`       // 'JAN', 'FEV', etc
	Realized      *float64       `gorm:"type:decimal(12,2)" json:"realized"` // nullable
	Meta          *float64       `gorm:"type:decimal(12,2)" json:"meta"`     // nullable
	Breakdown     datatypes.JSON `gorm:"type:jsonb" json:"breakdown"`        // PostgreSQL JSONB
	CreatedAt     time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"not null" json:"updated_at"`
	KpiCategory   KpiCategory    `gorm:"foreignKey:KpiCategoryID" json:"kpi_category,omitempty"`
}

func (MonthlyData) TableName() string {
	return "monthly_data"
}
