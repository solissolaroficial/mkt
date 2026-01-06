package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type CalendarPostModel struct {
	UUID               uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title              string         `gorm:"not null;size:500"`
	Description        *string        `gorm:"type:text"`
	Date               time.Time      `gorm:"type:date;not null;index:idx_date"`
	Time               string         `gorm:"not null;size:5"`
	Caption            *string        `gorm:"type:text"`
	Category           string         `gorm:"not null;size:50;index:idx_category"`
	Type               string         `gorm:"not null;size:50;index:idx_type"`
	Status             string         `gorm:"not null;size:50;index:idx_status"`
	AssigneeID         *uuid.UUID     `gorm:"type:uuid;index"`
	Assignee           *User          `gorm:"foreignKey:AssigneeID;references:UUID;constraint:OnDelete:SET NULL"`
	Platforms          datatypes.JSON `gorm:"type:jsonb"`
	PublishedPlatforms datatypes.JSON `gorm:"type:jsonb"`
	ImageURL           *string        `gorm:"type:varchar(500)"`
	History            datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt          time.Time      `gorm:"not null"`
	UpdatedAt          time.Time      `gorm:"not null"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

func (CalendarPostModel) TableName() string {
	return "calendar_posts"
}
