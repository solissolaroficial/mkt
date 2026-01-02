package model

import (
	"time"

	"github.com/google/uuid"
)

// Comment representa o model de comentário no banco de dados
type Comment struct {
	UUID      uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"uuid"`
	TaskID    uuid.UUID `gorm:"type:uuid;not null;index" json:"task_id"`
	Task      *Task     `gorm:"foreignKey:TaskID;references:UUID;constraint:OnDelete:CASCADE" json:"task,omitempty"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	User      *User     `gorm:"foreignKey:UserID;references:UUID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Text      string    `gorm:"type:text;not null" json:"text"`
	Timestamp time.Time `gorm:"not null" json:"timestamp"`
}

// TableName define o nome da tabela
func (Comment) TableName() string {
	return "comments"
}
