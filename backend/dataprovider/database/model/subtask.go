package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Subtask representa o model de subtarefa no banco de dados
type Subtask struct {
	UUID       uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"uuid"`
	TaskID     uuid.UUID      `gorm:"type:uuid;not null;index" json:"task_id"`
	Task       *Task          `gorm:"foreignKey:TaskID;references:UUID;constraint:OnDelete:CASCADE" json:"task,omitempty"`
	Title      string         `gorm:"not null" json:"title"`
	Completed  bool           `gorm:"not null;default:false;index" json:"completed"`
	AssigneeID *uuid.UUID     `gorm:"type:uuid;index" json:"assignee_id,omitempty"`
	Assignee   *User          `gorm:"foreignKey:AssigneeID;references:UUID;constraint:OnDelete:SET NULL" json:"assignee,omitempty"`
	DueDate    *time.Time     `gorm:"index" json:"due_date,omitempty"`
	CreatedAt  time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName define o nome da tabela
func (Subtask) TableName() string {
	return "subtasks"
}
