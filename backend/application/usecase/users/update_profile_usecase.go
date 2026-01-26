package users

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// UpdateProfileInput contains the data needed to update a user's profile
type UpdateProfileInput struct {
	UserID uuid.UUID
	Name   string
	Email  string
	Role   string
}

// UpdateProfile handles the business logic for updating a user's profile
type UpdateProfile struct {
	userGateway gateway.UserGateway
}

// NewUpdateProfile creates a new UpdateProfile use case
func NewUpdateProfile(userGateway gateway.UserGateway) *UpdateProfile {
	return &UpdateProfile{
		userGateway: userGateway,
	}
}

// Execute updates a user's profile information
func (uc *UpdateProfile) Execute(ctx context.Context, input UpdateProfileInput) (*entity.User, error) {
	// 1. Buscar usuário atual
	user, err := uc.userGateway.FindByID(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	// 2. Se o e-mail foi alterado, verificar se já existe
	if user.Email() != input.Email {
		existingUser, err := uc.userGateway.FindByEmail(ctx, input.Email)
		if err == nil && existingUser != nil && existingUser.ID() != input.UserID {
			return nil, errors.ErrUserEmailExists
		}
	}

	// 3. Atualizar dados (não atualizar o role, manter o valor atual)
	if err := user.Update(input.Name, input.Email, user.Role()); err != nil {
		return nil, err
	}

	// 4. Persistir alterações
	if err := uc.userGateway.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
