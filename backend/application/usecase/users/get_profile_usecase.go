package users

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetProfile handles the business logic for getting a user's profile
type GetProfile struct {
	userGateway gateway.UserGateway
}

// NewGetProfile creates a new GetProfile use case
func NewGetProfile(userGateway gateway.UserGateway) *GetProfile {
	return &GetProfile{
		userGateway: userGateway,
	}
}

// Execute retrieves a user's profile information
func (uc *GetProfile) Execute(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	user, err := uc.userGateway.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
