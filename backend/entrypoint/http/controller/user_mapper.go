package controller

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// UserMapper converts between domain entities and HTTP responses
type UserMapper struct{}

// ToResponse converts a User entity to a UserResponse
func (m *UserMapper) ToResponse(user *entity.User) response.UserResponse {
	return response.UserResponse{
		ID:        user.ID().String(),
		Email:     user.Email(),
		Name:      user.Name(),
		Role:      user.Role(),
		Active:    user.IsActive(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}
}

// ToResponseList converts a slice of User entities to a slice of UserResponse
func (m *UserMapper) ToResponseList(users []*entity.User) []response.UserResponse {
	responses := make([]response.UserResponse, len(users))
	for i, user := range users {
		responses[i] = m.ToResponse(user)
	}
	return responses
}

// ToPublicResponse converts a User entity to a PublicUserResponse (without email)
func (m *UserMapper) ToPublicResponse(user *entity.User) response.PublicUserResponse {
	return response.PublicUserResponse{
		ID:        user.ID().String(),
		Name:      user.Name(),
		Role:      user.Role(),
		Active:    user.IsActive(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}
}

// ToPublicResponseList converts a slice of User entities to a slice of PublicUserResponse
func (m *UserMapper) ToPublicResponseList(users []*entity.User) []response.PublicUserResponse {
	responses := make([]response.PublicUserResponse, len(users))
	for i, user := range users {
		responses[i] = m.ToPublicResponse(user)
	}
	return responses
}
