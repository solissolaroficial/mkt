package model

import (
	"time"

	"github.com/google/uuid"
)

// Notification representa o model de notificação no banco de dados
type Notification struct {
	UUID      uuid.UUID  `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"uuid"`
	UserUUID  uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_uuid"`
	User      *User      `gorm:"foreignKey:UserUUID;references:UUID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	TaskUUID  *uuid.UUID `gorm:"type:uuid;index" json:"task_uuid"`
	Task      *Task      `gorm:"foreignKey:TaskUUID;references:UUID;constraint:OnDelete:SET NULL" json:"task,omitempty"`
	Type      string     `gorm:"not null;index" json:"type"`
	Title     string     `gorm:"not null" json:"title"`
	Message   string     `gorm:"type:text;not null" json:"message"`
	Read      bool       `gorm:"not null;default:false;index" json:"read"`
	Archived  bool       `gorm:"not null;default:false;index" json:"archived"`
	Timestamp time.Time  `gorm:"not null;index" json:"timestamp"`
}

// TableName define o nome da tabela
func (Notification) TableName() string {
	return "notifications"
}
