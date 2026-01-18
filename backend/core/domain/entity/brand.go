package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Brand represents a brand for social posts and benchmarkings
type Brand struct {
	uuid      uuid.UUID
	name      string
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

// NewBrand creates a new Brand
func NewBrand(
	name string,
) (*Brand, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	now := time.Now()
	return &Brand{
		uuid:      uuid.New(),
		name:      name,
		createdAt: now,
		updatedAt: now,
	}, nil
}

// ReconstructBrand reconstructs a Brand from persistence
func ReconstructBrand(
	uuid uuid.UUID,
	name string,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
) *Brand {
	return &Brand{
		uuid:      uuid,
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
	}
}

// Getters
func (b *Brand) UUID() uuid.UUID {
	return b.uuid
}

func (b *Brand) Name() string {
	return b.name
}

func (b *Brand) CreatedAt() time.Time {
	return b.createdAt
}

func (b *Brand) UpdatedAt() time.Time {
	return b.updatedAt
}

func (b *Brand) DeletedAt() *time.Time {
	return b.deletedAt
}

// UpdateName updates the brand name
func (b *Brand) UpdateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	b.name = name
	b.updatedAt = time.Now()
	return nil
}

// SoftDelete marks the brand as deleted
func (b *Brand) SoftDelete() {
	now := time.Now()
	b.deletedAt = &now
	b.updatedAt = now
}

// IsActive checks if the brand is active (not deleted)
func (b *Brand) IsActive() bool {
	return b.deletedAt == nil
}
