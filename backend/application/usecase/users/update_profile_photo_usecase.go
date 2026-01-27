package users

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// UpdateProfilePhotoInput contains data needed to update a user's profile photo
type UpdateProfilePhotoInput struct {
	UserID uuid.UUID
	File   *multipart.FileHeader
}

// UpdateProfilePhoto handles business logic for updating a user's profile photo
type UpdateProfilePhoto struct {
	userGateway    gateway.UserGateway
	storageGateway gateway.StorageGateway
	maxFileSize    int64
}

// NewUpdateProfilePhoto creates a new UpdateProfilePhoto use case
func NewUpdateProfilePhoto(
	userGateway gateway.UserGateway,
	storageGateway gateway.StorageGateway,
) *UpdateProfilePhoto {
	const maxFileSize = 5 * 1024 * 1024 // 5MB
	return &UpdateProfilePhoto{
		userGateway:    userGateway,
		storageGateway: storageGateway,
		maxFileSize:    maxFileSize,
	}
}

// UpdateProfilePhotoOutput contains result of updating a user's profile photo
type UpdateProfilePhotoOutput struct {
	Key string // Returns the S3 key, not the URL
}

// Execute updates a user's profile photo
func (uc *UpdateProfilePhoto) Execute(ctx context.Context, input UpdateProfilePhotoInput) (*UpdateProfilePhotoOutput, error) {
	// 1. Validar se storage está configurado
	if uc.storageGateway == nil {
		return nil, errors.ErrStorageNotConfigured
	}

	// 2. Validar tamanho do arquivo
	if input.File.Size > uc.maxFileSize {
		return nil, errors.ErrFileTooLarge
	}

	// 3. Validar extensão do arquivo
	ext := strings.ToLower(filepath.Ext(input.File.Filename))
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}

	isAllowed := false
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return nil, errors.ErrInvalidFileType
	}

	// 4. Abrir arquivo
	file, err := input.File.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// 5. Criar chave única para S3
	// Formato: profiles/{userID}/photo.{ext}
	fileKey := fmt.Sprintf("profiles/%s/photo%s", input.UserID, ext)

	// 6. Upload para S3
	_, err = uc.storageGateway.UploadFile(ctx, file, fileKey, input.File.Header.Get("Content-Type"))
	if err != nil {
		return nil, fmt.Errorf("failed to upload profile photo: %w", err)
	}

	// 7. Atualizar usuário com key da foto
	if err := uc.userGateway.UpdateProfilePhotoKey(ctx, input.UserID, fileKey); err != nil {
		// Rollback: deletar arquivo do S3 em caso de erro
		_ = uc.storageGateway.DeleteFile(ctx, fileKey)
		return nil, fmt.Errorf("failed to update profile photo key: %w", err)
	}

	return &UpdateProfilePhotoOutput{
		Key: fileKey,
	}, nil
}
