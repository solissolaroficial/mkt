package service

import "github.com/google/uuid"

// TokenType represents the type of JWT token
type TokenType string

const (
	// AccessToken represents a short-lived access token
	AccessToken TokenType = "access"
	// RefreshToken represents a long-lived refresh token
	RefreshToken TokenType = "refresh"
)

// TokenClaims represents the claims extracted from a JWT token
type TokenClaims struct {
	UserID uuid.UUID
	Email  string
	Role   string
}

// JWTService defines the interface for JWT token operations
type JWTService interface {
	// GenerateAccessToken creates a new access token for the user
	GenerateAccessToken(userID uuid.UUID, email, role string) (string, error)

	// GenerateRefreshToken creates a new refresh token for the user
	GenerateRefreshToken(userID uuid.UUID) (string, error)

	// ValidateToken validates a token and returns the claims
	ValidateToken(token string, tokenType TokenType) (*TokenClaims, error)
}
