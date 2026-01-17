package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RecurrentPdvModel representa o modelo de banco de dados para PDVs recorrentes
type RecurrentPdvModel struct {
	UUID               uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	RepresentativeUUID uuid.UUID      `gorm:"not null;type:uuid;index:idx_representative_uuid;constraint:fk_recurrent_pdvs_representative,foreignKey:RepresentativeUUID,references:UUID,onDelete:RESTRICT,onUpdate:CASCADE"`
	Name               string         `gorm:"not null;size:200;index:idx_name"`
	City               *string        `gorm:"type:varchar(100);index:idx_city"`
	Followers          *int           `gorm:"type:integer"`
	InstagramProfile   *string        `gorm:"type:varchar(100)"`
	CreatedAt          time.Time      `gorm:"not null"`
	UpdatedAt          time.Time      `gorm:"not null"`
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

// TableName define o nome da tabela no banco de dados
func (RecurrentPdvModel) TableName() string {
	return "recurrent_pdvs"
}
