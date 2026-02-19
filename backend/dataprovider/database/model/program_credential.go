package model

import (
	"time"

	"github.com/google/uuid"
)

// ProgramCredentialModel represents a program credential entity in the database
type ProgramCredentialModel struct {
	UUID      uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string    `gorm:"not null;size:200;uniqueIndex"`
	User      string    `gorm:"size:200"`
	Password  string    `gorm:"size:500"`
	Access    string    `gorm:"size:200"`
	Notes     string    `gorm:"type:text"`
	Active    bool      `gorm:"default:true"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

// TableName specifies the table name for GORM
func (ProgramCredentialModel) TableName() string {
	return "program_credentials"
}
