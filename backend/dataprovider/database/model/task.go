package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Task representa o model de tarefa no banco de dados
type Task struct {
	UUID         uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"uuid"`
	Title        string         `gorm:"not null" json:"title"`
	Description  *string        `gorm:"type:text" json:"description,omitempty"`
	StartDate    *time.Time     `gorm:"index" json:"start_date,omitempty"`
	DueDate      time.Time      `gorm:"not null;index" json:"due_date"`
	Status       string         `gorm:"not null;index:idx_assignee_status,priority:2" json:"status"`
	Priority     string         `gorm:"not null;index" json:"priority"`
	Category     string         `gorm:"not null;index" json:"category"`
	AssigneeUUID *uuid.UUID     `gorm:"type:uuid;index:idx_assignee_status,priority:1" json:"assignee_uuid,omitempty"`
	Assignee     *User          `gorm:"foreignKey:AssigneeUUID;references:UUID;constraint:onDelete:SET NULL" json:"assignee,omitempty"`
	Archived     bool           `gorm:"not null;default:false;index" json:"archived"`
	SortOrder    int            `gorm:"not null;default:0;index" json:"sort_order"`
	Flows        string         `gorm:"type:jsonb" json:"flows"`
	CreatedAt    time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`

	// Relacionamentos
	Subtasks []*Subtask `gorm:"foreignKey:TaskUUID;references:UUID;constraint:OnDelete:CASCADE" json:"subtasks,omitempty"`
	Comments []*Comment `gorm:"foreignKey:TaskUUID;references:UUID;constraint:OnDelete:CASCADE" json:"comments,omitempty"`
}

// TableName define o nome da tabela
func (Task) TableName() string {
	return "tasks"
}
