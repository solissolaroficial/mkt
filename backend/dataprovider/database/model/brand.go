package model

import (
	"time"

	"github.com/google/uuid"
)

// BrandModel represents a brand entity for social posts and benchmarkings
type BrandModel struct {
	UUID      uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string     `gorm:"not null;size:200;uniqueIndex"`
	CreatedAt time.Time  `gorm:"not null"`
	UpdatedAt time.Time  `gorm:"not null"`
	DeletedAt *time.Time `gorm:"index"`
}

// TableName specifies the table name for GORM
func (BrandModel) TableName() string {
	return "brands"
}
