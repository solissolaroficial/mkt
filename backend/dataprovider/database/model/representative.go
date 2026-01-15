package model

import (
	"time"

	"github.com/google/uuid"
)

// RepresentativeModel represents the database model for representatives
type RepresentativeModel struct {
	UUID      uuid.UUID  `gorm:"type:uuid;primaryKey;column:uuid" json:"uuid"`
	Code      int        `gorm:"column:code;not null" json:"code"`
	Name      string     `gorm:"column:name;not null" json:"name"`
	Email     string     `gorm:"column:email;not null" json:"email"`
	Phone     string     `gorm:"column:phone" json:"phone"`
	Company   string     `gorm:"column:company;not null" json:"company"`
	Region    string     `gorm:"column:region;not null" json:"region"`
	City      string     `gorm:"column:city" json:"city"`
	Attendant string     `gorm:"column:attendant" json:"attendant"`
	Active    bool       `gorm:"column:active;not null;default:true" json:"active"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index" json:"deletedAt,omitempty"`
}

// TableName specifies the table name for GORM
func (RepresentativeModel) TableName() string {
	return "representatives"
}
