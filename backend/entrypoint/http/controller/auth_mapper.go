package controller

import (
	"github.com/seu-usuario/solis-backend/application/usecase/auth"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// AuthMapper handles conversion between auth use case outputs and HTTP responses
type AuthMapper struct{}

// NewAuthMapper creates a new AuthMapper instance
func NewAuthMapper() *AuthMapper {
	return &AuthMapper{}
}

// ToLoginResponse converts use case output to HTTP response
func (m *AuthMapper) ToLoginResponse(output *auth.LoginOutput) *response.LoginResponse {
	return &response.LoginResponse{
		AccessToken:  output.AccessToken,
		RefreshToken: output.RefreshToken,
		User: response.UserData{
			ID:    output.UserID.String(),
			Email: output.Email,
			Name:  output.Name,
			Role:  output.Role, // Adicionar
		},
	}
}
