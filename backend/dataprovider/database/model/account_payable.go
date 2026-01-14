package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountPayableModel struct {
	UUID          uuid.UUID      `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Supplier      string         `gorm:"not null;size:200;index:idx_supplier"`
	Description   string         `gorm:"not null;size:500;index:idx_description"`
	Amount        float64        `gorm:"not null;index:idx_amount"`
	DueDate       time.Time      `gorm:"not null;index:idx_due_date"`
	NFArrived     bool           `gorm:"not null;default:false;index:idx_nf_arrived"`
	BoletoArrived bool           `gorm:"not null;default:false;index:idx_boleto_arrived"`
	Status        string         `gorm:"not null;size:20;default:'pending';index:idx_status"`
	Recurrence    string         `gorm:"not null;size:10;default:'none';index:idx_recurrence"`
	CreatedAt     time.Time      `gorm:"not null;index:idx_created_at"`
	UpdatedAt     time.Time      `gorm:"not null"`
	DeletedAt     gorm.DeletedAt `gorm:"index:idx_deleted_at"`
}

func (AccountPayableModel) TableName() string {
	return "accounts_payable"
}
