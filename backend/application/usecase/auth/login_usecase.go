package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/service"
)

// LoginOutput represents the output of the login use case
type LoginOutput struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	UserID       uuid.UUID `json:"user_id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Role         string    `json:"role"`
}

// LoginInput represents the input data for login
type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginUseCase handles user authentication
type LoginUseCase struct {
	userGateway   gateway.UserGateway
	hasherService service.HasherService
	jwtService    service.JWTService
}

// NewLoginUseCase creates a new LoginUseCase instance
func NewLoginUseCase(
	userGateway gateway.UserGateway,
	hasherService service.HasherService,
	jwtService service.JWTService,
) *LoginUseCase {
	return &LoginUseCase{
		userGateway:   userGateway,
		hasherService: hasherService,
		jwtService:    jwtService,
	}
}

// Execute performs the login operation
func (uc *LoginUseCase) Execute(ctx context.Context, input LoginInput) (*LoginOutput, error) {
	// 1. Buscar usuário por email usando gateway
	user, err := uc.userGateway.FindByEmail(ctx, input.Email)
	if err != nil {
		if err == errors.ErrUserNotFound {
			return nil, errors.ErrInvalidCredentials
		}
		return nil, err
	}

	// 2. Verificar se usuário está ativo
	if !user.IsActive() {
		return nil, errors.ErrUserInactive
	}

	// 3. Comparar senha usando hasher service
	if err := uc.hasherService.Compare(user.Password(), input.Password); err != nil {
		return nil, errors.ErrInvalidCredentials
	}

	// 4. Gerar access token usando jwt service
	accessToken, err := uc.jwtService.GenerateAccessToken(user.ID(), user.Email(), user.Role())
	if err != nil {
		return nil, err
	}

	// 5. Gerar refresh token usando jwt service
	refreshToken, err := uc.jwtService.GenerateRefreshToken(user.ID())
	if err != nil {
		return nil, err
	}

	// 6. Retornar struct com tokens e informações do usuário
	return &LoginOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       user.ID(),
		Email:        user.Email(),
		Name:         user.Name(),
		Role:         user.Role(),
	}, nil
}
