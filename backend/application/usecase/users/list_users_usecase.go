package users

import (
	"context"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// ListUsersUseCase handles the business logic for listing all active users
type ListUsersUseCase struct {
	userGateway gateway.UserGateway
}

// NewListUsersUseCase creates a new ListUsersUseCase instance
func NewListUsersUseCase(userGateway gateway.UserGateway) *ListUsersUseCase {
	return &ListUsersUseCase{
		userGateway: userGateway,
	}
}

// Execute retrieves all active users from the database
func (uc *ListUsersUseCase) Execute(ctx context.Context) ([]*entity.User, error) {
	users, err := uc.userGateway.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
