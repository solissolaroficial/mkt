package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SocialDailyAggregationModel struct {
	ID              uuid.UUID  `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	BrandName       string     `gorm:"type:varchar(200);not null;uniqueIndex:idx_brand_date,priority:1;index:idx_brand"`
	AggregationDate time.Time  `gorm:"type:date;not null;uniqueIndex:idx_brand_date,priority:2;index:idx_date"`
	TotalPosts      int        `gorm:"type:int;not null;default:0"`
	TotalLikes      int        `gorm:"type:int;not null;default:0"`
	TotalComments   int        `gorm:"type:int;not null;default:0"`
	TotalShares     *int       `gorm:"type:int"`
	AvgLikes        float64    `gorm:"type:decimal(10,2);not null;default:0"`
	AvgComments     float64    `gorm:"type:decimal(10,2);not null;default:0"`
	AvgShares       *float64   `gorm:"type:decimal(10,2)"`
	FollowersAtDate *int       `gorm:"type:int"`
	EngagementRate  float64    `gorm:"type:decimal(10,2);not null;default:0"`
	CreatedAt       time.Time  `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt       time.Time  `gorm:"type:timestamp;not null;default:now()"`
	DeletedAt       *time.Time `gorm:"type:timestamp;index"`
}

func (SocialDailyAggregationModel) TableName() string {
	return "social_daily_aggregations"
}

func (m *SocialDailyAggregationModel) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}
