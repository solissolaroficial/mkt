package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type MonthlyData struct {
	UUID          uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"uuid"`
	KpiCategoryID uuid.UUID      `gorm:"type:uuid;not null;uniqueIndex:idx_kpi_year_month_active,priority:1,where:deleted_at IS NULL" json:"kpi_category_id"`
	KpiCategory   *KpiCategory   `gorm:"foreignKey:KpiCategoryID;references:UUID" json:"kpi_category,omitempty"`
	Year          int            `gorm:"not null;uniqueIndex:idx_kpi_year_month_active,priority:2,where:deleted_at IS NULL" json:"year"`         // Year (e.g., 2024, 2025)
	Month         string         `gorm:"not null;size:3;uniqueIndex:idx_kpi_year_month_active,priority:3,where:deleted_at IS NULL" json:"month"` // 'JAN', 'FEV', etc
	Realized      *float64       `gorm:"type:decimal(12,2)" json:"realized"`                                                                     // nullable
	Meta          *float64       `gorm:"type:decimal(12,2)" json:"meta"`                                                                         // nullable
	Breakdown     datatypes.JSON `gorm:"type:jsonb" json:"breakdown"`                                                                            // PostgreSQL JSONB
	Logs          datatypes.JSON `gorm:"type:jsonb" json:"logs"`                                                                                 // PostgreSQL JSONB - array of log entries
	CreatedAt     time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt     *time.Time     `gorm:"index" json:"deleted_at,omitempty"`
}

func (MonthlyData) TableName() string {
	return "monthly_data"
}
