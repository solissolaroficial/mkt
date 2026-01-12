package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShowroomItemModel struct {
	UUID             uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	PDV              string         `gorm:"not null;size:200;index:idx_pdv"`
	City             *string        `gorm:"type:varchar(100);index:idx_city"`
	Contact          *string        `gorm:"type:varchar(100)"`
	RepName          string         `gorm:"not null;size:100;index:idx_rep_name"`
	DeliveryForecast *string        `gorm:"type:varchar(10)"`
	WorkshopDate     *string        `gorm:"type:varchar(10)"`
	Delivered        bool           `gorm:"not null;default:false;index:idx_delivered"`
	CreatedAt        time.Time      `gorm:"not null;index:idx_created_at"`
	UpdatedAt        time.Time      `gorm:"not null"`
	DeletedAt        gorm.DeletedAt `gorm:"index:idx_deleted_at"`
}

func (ShowroomItemModel) TableName() string {
	return "showroom_items"
}
