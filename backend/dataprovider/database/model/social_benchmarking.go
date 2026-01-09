package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SocialBenchmarkingModel struct {
	UUID           uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	BrandName      string         `gorm:"not null;size:200;index:idx_brand_name"`
	AvgLikes       float64        `gorm:"not null"`
	AvgComments    float64        `gorm:"not null"`
	Followers      *int           `gorm:"type:integer"`
	EngagementRate float64        `gorm:"not null;index:idx_engagement_rate"`
	CreatedAt      time.Time      `gorm:"not null;index:idx_created_at"`
	UpdatedAt      time.Time      `gorm:"not null"`
	DeletedAt      gorm.DeletedAt `gorm:"index:idx_deleted_at"`
}

// Índices compostos para performance:
// - idx_brand_name_deleted_at: (brand_name, deleted_at) para queries filtradas por marca
// - idx_created_at_deleted_at: (created_at, deleted_at) para queries por data
// - idx_engagement_rate: (engagement_rate) para ordenação por taxa de engajamento

func (SocialBenchmarkingModel) TableName() string {
	return "social_benchmarkings"
}
