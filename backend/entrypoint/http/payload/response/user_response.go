package response

import "time"

// UserResponse represents a user in HTTP responses
type UserResponse struct {
	ID              string    `json:"id"`
	Email           string    `json:"email"`
	Name            string    `json:"name"`
	Role            string    `json:"role"`
	Active          bool      `json:"active"`
	ProfilePhotoKey string    `json:"profile_photo_key,omitempty"`
	ProfilePhotoURL string    `json:"profile_photo_url,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// UploadProfilePhotoResponse representa a resposta de upload de foto de perfil
type UploadProfilePhotoResponse struct {
	Key string `json:"key"`
}

// PresignedURLResponse representa a resposta de presigned URL
type PresignedURLResponse struct {
	URL       string `json:"url"`
	ExpiresIn int64  `json:"expires_in"` // Tempo de expiração em segundos
}

// SuccessResponse representa uma resposta de sucesso genérica
type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
