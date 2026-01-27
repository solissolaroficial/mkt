package users

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// RemoveProfilePhoto remove a foto de perfil de um usuário
type RemoveProfilePhoto struct {
	userGateway    gateway.UserGateway
	storageGateway gateway.StorageGateway
}

// NewRemoveProfilePhoto cria um novo use case de remoção de foto de perfil
func NewRemoveProfilePhoto(
	userGateway gateway.UserGateway,
	storageGateway gateway.StorageGateway,
) *RemoveProfilePhoto {
	return &RemoveProfilePhoto{
		userGateway:    userGateway,
		storageGateway: storageGateway,
	}
}

// RemoveProfilePhotoInput contém os dados necessários para remover a foto de perfil
type RemoveProfilePhotoInput struct {
	UserID uuid.UUID
}

// RemoveProfilePhotoOutput contém o resultado da remoção da foto de perfil
type RemoveProfilePhotoOutput struct {
	Success bool
}

// Execute remove a foto de perfil de um usuário
func (uc *RemoveProfilePhoto) Execute(ctx context.Context, input RemoveProfilePhotoInput) (*RemoveProfilePhotoOutput, error) {
	// 1. Validar se storage está configurado
	if uc.storageGateway == nil {
		return nil, errors.ErrStorageNotConfigured
	}

	// 2. Buscar usuário para obter a key da foto atual
	user, err := uc.userGateway.FindByID(ctx, input.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	// 3. Se não houver foto, não fazer nada
	profilePhotoKey := user.ProfilePhotoKey()
	if profilePhotoKey == "" {
		return &RemoveProfilePhotoOutput{Success: true}, nil
	}

	// 4. Deletar arquivo do S3
	if err := uc.storageGateway.DeleteFile(ctx, profilePhotoKey); err != nil {
		return nil, fmt.Errorf("failed to delete profile photo from S3: %w", err)
	}

	// 5. Atualizar usuário removendo a key da foto
	if err := uc.userGateway.UpdateProfilePhotoKey(ctx, input.UserID, ""); err != nil {
		return nil, fmt.Errorf("failed to remove profile photo key: %w", err)
	}

	return &RemoveProfilePhotoOutput{Success: true}, nil
}
