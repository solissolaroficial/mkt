package model

import (
	"time"

	"github.com/google/uuid"
)

type KpiCategory struct {
	UUID      uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"uuid"`
	Title     string    `gorm:"not null" json:"title"`
	Slug      string    `gorm:"not null;uniqueIndex" json:"slug"`
	Color     string    `gorm:"not null" json:"color"`
	Unit      string    `gorm:"not null" json:"unit"` // "currency", "percent", "number"
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`

	// Relacionamentos Has Many (inverso)
	MonthlyData []MonthlyData `gorm:"foreignKey:KpiCategoryID;constraint:OnDelete:CASCADE" json:"monthly_data,omitempty"`
}

func (KpiCategory) TableName() string {
	return "kpi_categories"
}
