package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RecurrentPdvModel representa o modelo de banco de dados para PDVs recorrentes
type RecurrentPdvModel struct {
	UUID             uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name             string         `gorm:"not null;size:200;index:idx_name"`
	RepName          string         `gorm:"not null;size:100;index:idx_rep_name"`
	City             *string        `gorm:"type:varchar(100);index:idx_city"`
	Followers        *int           `gorm:"type:integer"`
	InstagramProfile *string        `gorm:"type:varchar(100)"`
	CreatedAt        time.Time      `gorm:"not null"`
	UpdatedAt        time.Time      `gorm:"not null"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

// TableName define o nome da tabela no banco de dados
func (RecurrentPdvModel) TableName() string {
	return "recurrent_pdvs"
}
