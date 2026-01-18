package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SocialPostModel struct {
	ID              uuid.UUID   `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	BrandID         uuid.UUID   `gorm:"type:uuid;not null;index:idx_brand_id;constraint:fk_social_posts_brand,foreignKey:BrandID,references:UUID,onDelete:RESTRICT,onUpdate:CASCADE"`
	Brand           *BrandModel `gorm:"foreignKey:BrandID"`
	Platform        string      `gorm:"type:varchar(50);not null;index"`
	PostDate        time.Time   `gorm:"type:date;not null;index:idx_brand_id,priority:2"`
	PostTime        *time.Time  `gorm:"type:time"`
	Likes           int         `gorm:"type:int;not null;default:0"`
	Comments        int         `gorm:"type:int;not null;default:0"`
	Shares          *int        `gorm:"type:int"`
	PostType        string      `gorm:"type:varchar(50);not null;index"`
	Caption         *string     `gorm:"type:text"`
	FollowersAtPost *int        `gorm:"type:int"`
	CreatedAt       time.Time   `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt       time.Time   `gorm:"type:timestamp;not null;default:now()"`
	DeletedAt       *time.Time  `gorm:"type:timestamp;index"`
}

func (SocialPostModel) TableName() string {
	return "social_posts"
}

func (m *SocialPostModel) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return nil
}
