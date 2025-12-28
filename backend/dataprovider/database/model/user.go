package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UUID      uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"uuid"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Name      string    `gorm:"not null" json:"name"`
	Role      string    `gorm:"not null" json:"role"` // "admin", "marketing", "commercial"
	Active    bool      `gorm:"default:true" json:"active"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
