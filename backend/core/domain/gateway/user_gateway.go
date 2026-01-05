package gateway

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

// UserGateway defines the interface for user data operations
type UserGateway interface {
	// Save creates a new user or updates an existing one
	Save(ctx context.Context, user *entity.User) error

	// FindByID retrieves a user by their ID
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)

	// FindByEmail retrieves a user by their email address
	FindByEmail(ctx context.Context, email string) (*entity.User, error)

	// Update modifies an existing user
	Update(ctx context.Context, user *entity.User) error

	// ExistsByEmail checks if a user with the given email already exists
	ExistsByEmail(ctx context.Context, email string) (bool, error)

	// FindByName retrieves a user by their name
	FindByName(ctx context.Context, name string) (*entity.User, error)

	// ListAll retrieves all active users
	ListAll(ctx context.Context) ([]*entity.User, error)
}
