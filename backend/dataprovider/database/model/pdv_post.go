package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PdvPostModel representa o modelo de banco de dados para posts de PDV
type PdvPostModel struct {
	UUID      uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	RepName   string         `gorm:"not null;size:100;index:idx_rep_name"`
	PdvName   string         `gorm:"not null;size:200;index:idx_pdv_name"`
	PostDate  time.Time      `gorm:"not null;index:idx_post_date"`
	Month     string         `gorm:"not null;size:3;index:idx_month;check:month IN ('JAN','FEV','MAR','ABR','MAI','JUN','JUL','AGO','SET','OUT','NOV','DEZ')"`
	Platform  string         `gorm:"not null;size:50;index:idx_platform;check:platform IN ('instagram','facebook','linkedin','youtube','tiktok')"`
	Link      *string        `gorm:"type:varchar(500)"`
	ProofUrl  *string        `gorm:"type:varchar(500)"`
	Status    string         `gorm:"not null;size:50;index:idx_status;check:status IN ('pending','approved','rejected','published','cancelled')"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName define o nome da tabela no banco de dados
func (PdvPostModel) TableName() string {
	return "pdv_posts"
}
