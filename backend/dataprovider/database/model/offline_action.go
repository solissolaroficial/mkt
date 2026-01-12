package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OfflineActionModel struct {
	UUID             uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	RequestedAmount  float64        `gorm:"not null"`
	ActionDate       time.Time      `gorm:"not null;index:idx_action_date"`
	Category         string         `gorm:"not null;size:100;index:idx_category"`
	Month            string         `gorm:"not null;size:10;index:idx_month"`
	ApprovedAmount   *string        `gorm:"type:varchar(50)"`
	OrderNumber      *string        `gorm:"type:varchar(50);index:idx_order_number"`
	DepartureDate    *string        `gorm:"type:varchar(10)"`
	DeliveryForecast *string        `gorm:"type:varchar(10)"`
	DeliveryDate     *string        `gorm:"type:varchar(10)"`
	City             *string        `gorm:"type:varchar(100);index:idx_city"`
	UF               *string        `gorm:"type:varchar(2)"`
	Scored           string         `gorm:"not null;size:10;default:AINDA NÃO"`
	Status           string         `gorm:"not null;size:50;index:idx_status"`
	Observation      *string        `gorm:"type:varchar(500)"`
	PDV              *string        `gorm:"not null;size:200;index:idx_pdv"`
	RepName          string         `gorm:"not null;size:100;index:idx_rep_name"`
	CreatedAt        time.Time      `gorm:"not null;index:idx_created_at"`
	UpdatedAt        time.Time      `gorm:"not null"`
	DeletedAt        gorm.DeletedAt `gorm:"index:idx_deleted_at"`
}

func (OfflineActionModel) TableName() string {
	return "offline_actions"
}
