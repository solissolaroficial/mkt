package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// User represents a system user with authentication credentials
type User struct {
	id              uuid.UUID
	email           string
	password        string // hash bcrypt
	name            string
	role            string // "admin", "marketing", "commercial"
	active          bool
	profilePhotoKey string // Key da foto de perfil no S3 (não URL)
	createdAt       time.Time
	updatedAt       time.Time
}

// NewUser creates a new user with validation
func NewUser(email, password, name, role string) (*User, error) {
	user := &User{
		id:              uuid.New(),
		email:           email,
		password:        password, // Should be already hashed when passed to this constructor
		name:            name,
		role:            role,
		active:          true,
		profilePhotoKey: "", // Inicializar vazio
		createdAt:       time.Now(),
		updatedAt:       time.Now(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

// ReconstructUser reconstructs a user from database (without creation validation)
func ReconstructUser(id uuid.UUID, email, password, name, role string, active bool, profilePhotoKey string, createdAt, updatedAt time.Time) *User {
	return &User{
		id:              id,
		email:           email,
		password:        password,
		name:            name,
		role:            role,
		active:          active,
		profilePhotoKey: profilePhotoKey,
		createdAt:       createdAt,
		updatedAt:       updatedAt,
	}
}

// Getters (encapsulation)
func (u *User) ID() uuid.UUID           { return u.id }
func (u *User) Email() string           { return u.email }
func (u *User) Password() string        { return u.password }
func (u *User) Name() string            { return u.name }
func (u *User) Role() string            { return u.role }
func (u *User) IsActive() bool          { return u.active }
func (u *User) ProfilePhotoKey() string { return u.profilePhotoKey }
func (u *User) CreatedAt() time.Time    { return u.createdAt }
func (u *User) UpdatedAt() time.Time    { return u.updatedAt }

// Update updates user's profile information
func (u *User) Update(name, email, role string) error {
	u.name = name
	u.email = email
	u.role = role
	u.updatedAt = time.Now()

	return u.Validate()
}

// UpdateProfilePhotoKey atualiza a key da foto de perfil
func (u *User) UpdateProfilePhotoKey(profilePhotoKey string) {
	u.profilePhotoKey = profilePhotoKey
	u.updatedAt = time.Now()
}

// Validate performs business validation
func (u *User) Validate() error {
	if u.email == "" {
		return errors.New("email is required")
	}
	if u.name == "" {
		return errors.New("name is required")
	}
	if u.role != "admin" && u.role != "marketing" && u.role != "commercial" {
		return errors.New("invalid role: must be 'admin', 'marketing', or 'commercial'")
	}
	if u.password == "" {
		return errors.New("password is required")
	}
	return nil
}
