package migrations

import (
	"gorm.io/gorm"
)

// AddRepresentativeMonthlyGoalsTable creates the representative_monthly_goals table
func AddRepresentativeMonthlyGoalsTable(db *gorm.DB) error {
	return db.AutoMigrate(&RepresentativeMonthlyGoal{})
}

// RepresentativeMonthlyGoal represents the monthly goals table
type RepresentativeMonthlyGoal struct {
	ID               uint    `gorm:"primaryKey;autoIncrement"`
	GUID             string  `gorm:"type:uuid;not null;uniqueIndex"`
	RepresentativeID string  `gorm:"type:uuid;not null;index"`
	Month            int     `gorm:"not null;index"`
	Year             int     `gorm:"not null;index"`
	Target           float64 `gorm:"not null"`
	Realized         float64 `gorm:"not null;default:0"`
	CreatedAt        string  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt        string  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt        *string `gorm:"type:timestamp;index"`
}
