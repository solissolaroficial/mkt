package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Flow representa o model de fluxo no banco de dados
type Flow struct {
	UUID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"uuid"`
	Name        string         `gorm:"not null" json:"name"`
	Description *string        `gorm:"type:text" json:"description,omitempty"`
	Color       *string        `gorm:"type:varchar(7)" json:"color,omitempty"` // Hex code
	SortOrder   int            `gorm:"not null;default:0;index" json:"sort_order"`
	CreatedAt   time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// TableName define o nome da tabela
func (Flow) TableName() string {
	return "flows"
}
