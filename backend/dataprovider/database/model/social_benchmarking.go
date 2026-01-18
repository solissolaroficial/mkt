package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SocialBenchmarkingModel struct {
	UUID           uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	BrandID        uuid.UUID      `gorm:"not null;type:uuid;index:idx_brand_id;constraint:fk_social_benchmarkings_brand,foreignKey:BrandID,references:UUID,onDelete:RESTRICT,onUpdate:CASCADE"`
	Brand          *BrandModel    `gorm:"foreignKey:BrandID"`
	AvgLikes       float64        `gorm:"not null"`
	AvgComments    float64        `gorm:"not null"`
	Followers      *int           `gorm:"type:integer"`
	EngagementRate float64        `gorm:"not null;index:idx_engagement_rate"`
	CreatedAt      time.Time      `gorm:"not null;index:idx_created_at"`
	UpdatedAt      time.Time      `gorm:"not null"`
	DeletedAt      gorm.DeletedAt `gorm:"index:idx_deleted_at"`
}

// Índices compostos para performance:
// - idx_brand_id_deleted_at: (brand_id, deleted_at) para queries filtradas por marca
// - idx_created_at_deleted_at: (created_at, deleted_at) para queries por data
// - idx_engagement_rate: (engagement_rate) para ordenação por taxa de engajamento

func (SocialBenchmarkingModel) TableName() string {
	return "social_benchmarkings"
}
