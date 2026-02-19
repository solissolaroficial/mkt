package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// ProgramCredential represents a credential for accessing external programs/tools
type ProgramCredential struct {
	id        uuid.UUID
	name      string
	user      string
	password  string
	access    string
	notes     string
	active    bool
	createdAt time.Time
	updatedAt time.Time
}

// NewProgramCredential creates a new ProgramCredential
func NewProgramCredential(
	name string,
	user string,
	password string,
	access string,
	notes string,
) (*ProgramCredential, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	now := time.Now()
	return &ProgramCredential{
		id:        uuid.New(),
		name:      name,
		user:      user,
		password:  password,
		access:    access,
		notes:     notes,
		active:    true,
		createdAt: now,
		updatedAt: now,
	}, nil
}

// ReconstructProgramCredential reconstructs a ProgramCredential from persistence
func ReconstructProgramCredential(
	id uuid.UUID,
	name string,
	user string,
	password string,
	access string,
	notes string,
	active bool,
	createdAt time.Time,
	updatedAt time.Time,
) *ProgramCredential {
	return &ProgramCredential{
		id:        id,
		name:      name,
		user:      user,
		password:  password,
		access:    access,
		notes:     notes,
		active:    active,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

// Getters
func (c *ProgramCredential) ID() uuid.UUID {
	return c.id
}

func (c *ProgramCredential) Name() string {
	return c.name
}

func (c *ProgramCredential) User() string {
	return c.user
}

func (c *ProgramCredential) Password() string {
	return c.password
}

func (c *ProgramCredential) Access() string {
	return c.access
}

func (c *ProgramCredential) Notes() string {
	return c.notes
}

func (c *ProgramCredential) Active() bool {
	return c.active
}

func (c *ProgramCredential) CreatedAt() time.Time {
	return c.createdAt
}

func (c *ProgramCredential) UpdatedAt() time.Time {
	return c.updatedAt
}

// Update updates the credential fields
func (c *ProgramCredential) Update(
	name string,
	user string,
	password string,
	access string,
	notes string,
) error {
	if name == "" {
		return errors.New("name is required")
	}

	c.name = name
	c.user = user
	c.password = password
	c.access = access
	c.notes = notes
	c.updatedAt = time.Now()
	return nil
}

// Deactivate marks the credential as inactive
func (c *ProgramCredential) Deactivate() {
	c.active = false
	c.updatedAt = time.Now()
}

// Activate marks the credential as active
func (c *ProgramCredential) Activate() {
	c.active = true
	c.updatedAt = time.Now()
}

// Validate validates the credential data
func (c *ProgramCredential) Validate() error {
	if c.name == "" {
		return errors.New("name is required")
	}
	return nil
}
