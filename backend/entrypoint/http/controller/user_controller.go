package controller

import (
	"context"
	stderrors "errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/application/usecase/users"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// UserController handles HTTP requests for users
type UserController struct {
	listUsersUseCase          *users.ListUsersUseCase
	changePasswordUseCase     *users.ChangePasswordUseCase
	updateProfileUseCase      *users.UpdateProfile
	getProfileUseCase         *users.GetProfile
	updateProfilePhotoUseCase *users.UpdateProfilePhoto
	removeProfilePhotoUseCase *users.RemoveProfilePhoto
	storageGateway            gateway.StorageGateway
	mapper                    *UserMapper
}

// NewUserController creates a new UserController instance
func NewUserController(
	listUsersUseCase *users.ListUsersUseCase,
	changePasswordUseCase *users.ChangePasswordUseCase,
	updateProfileUseCase *users.UpdateProfile,
	getProfileUseCase *users.GetProfile,
	updateProfilePhotoUseCase *users.UpdateProfilePhoto,
	removeProfilePhotoUseCase *users.RemoveProfilePhoto,
	storageGateway gateway.StorageGateway,
) *UserController {
	return &UserController{
		listUsersUseCase:          listUsersUseCase,
		changePasswordUseCase:     changePasswordUseCase,
		updateProfileUseCase:      updateProfileUseCase,
		getProfileUseCase:         getProfileUseCase,
		updateProfilePhotoUseCase: updateProfilePhotoUseCase,
		removeProfilePhotoUseCase: removeProfilePhotoUseCase,
		storageGateway:            storageGateway,
		mapper:                    &UserMapper{},
	}
}

// getProfilePhotoURL retorna a URL completa da foto de perfil (DEPRECATED - use presigned URL)
func (c *UserController) getProfilePhotoURL(profilePhotoKey string) string {
	if profilePhotoKey == "" {
		return ""
	}
	return c.storageGateway.GetFileURL(profilePhotoKey)
}

// getProfilePhotoPresignedURL retorna uma presigned URL da foto de perfil
func (c *UserController) getProfilePhotoPresignedURL(ctx context.Context, profilePhotoKey string) (string, error) {
	if profilePhotoKey == "" {
		return "", stderrors.New("no profile photo key provided")
	}
	if c.storageGateway == nil {
		return "", stderrors.New("storage is not configured")
	}

	// Gerar presigned URL com expiração de 1 hora
	url, err := c.storageGateway.GetPresignedURL(ctx, profilePhotoKey, 1*time.Hour)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return url, nil
}

// ListUsers returns all active users with pagination
// GET /api/users?page=1&limit=20
func (c *UserController) ListUsers(ctx *fiber.Ctx) error {
	// Parse pagination parameters
	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 20)

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20
	}

	users, err := c.listUsersUseCase.Execute(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "Failed to list users",
		})
	}

	// Apply pagination
	total := len(users)
	totalPages := (total + limit - 1) / limit
	start := (page - 1) * limit
	end := start + limit

	if start >= total {
		// Return empty result for out of range pages
		return ctx.JSON(response.PaginationResponse{
			Page:       page,
			PageSize:   limit,
			Total:      int64(total),
			TotalPages: totalPages,
			Data:       []response.PublicUserResponse{},
		})
	}

	if end > total {
		end = total
	}

	paginatedUsers := users[start:end]
	userResponses := c.mapper.ToPublicResponseList(paginatedUsers)

	// Gerar presigned URLs para usuários com foto
	for i := range userResponses {
		if userResponses[i].ProfilePhotoKey != "" {
			url, err := c.getProfilePhotoPresignedURL(ctx.Context(), userResponses[i].ProfilePhotoKey)
			if err == nil {
				userResponses[i].ProfilePhotoURL = url
			}
		}
	}

	return ctx.JSON(response.PaginationResponse{
		Page:       page,
		PageSize:   limit,
		Total:      int64(total),
		TotalPages: totalPages,
		Data:       userResponses,
	})
}

// ListAllUsers returns all active users without pagination
// GET /api/users/all
func (c *UserController) ListAllUsers(ctx *fiber.Ctx) error {
	users, err := c.listUsersUseCase.Execute(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "Failed to list users",
		})
	}

	userResponses := c.mapper.ToPublicResponseList(users)

	// Gerar presigned URLs para usuários com foto
	for i := range userResponses {
		if userResponses[i].ProfilePhotoKey != "" {
			url, err := c.getProfilePhotoPresignedURL(ctx.Context(), userResponses[i].ProfilePhotoKey)
			if err == nil {
				userResponses[i].ProfilePhotoURL = url
			}
		}
	}

	return ctx.JSON(userResponses)
}

// ChangePassword altera a senha do usuário autenticado
// PUT /api/settings/password
func (c *UserController) ChangePassword(ctx *fiber.Ctx) error {
	// 1. Obter ID do usuário do contexto (injetado pelo middleware de autenticação)
	userIDValue := ctx.Locals("userID")
	if userIDValue == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "User not authenticated",
		})
	}

	// O middleware armazena userID como string, então precisamos converter para UUID
	userIDStr, ok := userIDValue.(string)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "Invalid user ID in context",
		})
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "Invalid user ID format",
		})
	}

	// 2. Parse request body
	var req request.ChangePasswordRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// 3. Executar use case
	input := users.ChangePasswordInput{
		UserID:          userID,
		CurrentPassword: req.CurrentPassword,
		NewPassword:     req.NewPassword,
	}

	if err := c.changePasswordUseCase.Execute(ctx.Context(), input); err != nil {
		// Tratar erros de domínio específicos usando errors.Is
		switch {
		case stderrors.Is(err, errors.ErrUserNotFound):
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "User not found",
			})
		case stderrors.Is(err, errors.ErrCurrentPasswordMismatch):
			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Error: "Current password is incorrect",
			})
		case stderrors.Is(err, errors.ErrPasswordTooWeak):
			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Error: "Password does not meet strength requirements",
			})
		case stderrors.Is(err, errors.ErrPasswordSameAsCurrent):
			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Error: "New password must be different from current password",
			})
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
				Error: "Failed to change password",
			})
		}
	}

	// 4. Retornar resposta de sucesso
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessChangePasswordResponse())
}

// UpdateProfile atualiza os dados do perfil do usuário autenticado
// PUT /api/settings/profile
func (c *UserController) UpdateProfile(ctx *fiber.Ctx) error {
	// 1. Obter ID do usuário do contexto (injetado pelo middleware de autenticação)
	userIDValue := ctx.Locals("userID")
	if userIDValue == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "User not authenticated",
		})
	}

	// O middleware armazena userID como string, então precisamos converter para UUID
	userIDStr, ok := userIDValue.(string)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "Invalid user ID in context",
		})
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "Invalid user ID format",
		})
	}

	// 2. Parse request body
	var req request.UpdateProfileRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// 3. Executar use case
	input := users.UpdateProfileInput{
		UserID: userID,
		Name:   req.Name,
		Email:  req.Email,
		Role:   req.Role,
	}

	user, err := c.updateProfileUseCase.Execute(ctx.Context(), input)
	if err != nil {
		// Tratar erros de domínio específicos usando errors.Is
		switch {
		case stderrors.Is(err, errors.ErrUserNotFound):
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "User not found",
			})
		case stderrors.Is(err, errors.ErrUserEmailExists):
			return ctx.Status(fiber.StatusConflict).JSON(response.ErrorResponse{
				Error: "Email already exists",
			})
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
				Error: "Failed to update profile",
			})
		}
	}

	// 4. Converter para response
	userResponse := response.UserResponse{
		ID:              user.ID().String(),
		Name:            user.Name(),
		Email:           user.Email(),
		Role:            user.Role(),
		Active:          user.IsActive(),
		ProfilePhotoKey: user.ProfilePhotoKey(),
		ProfilePhotoURL: "", // Não retornar URL pública - usar presigned URL
		CreatedAt:       user.CreatedAt(),
		UpdatedAt:       user.UpdatedAt(),
	}

	// 5. Retornar resposta de sucesso
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessUpdateProfileResponse(userResponse))
}

// GetProfile retorna os dados do perfil do usuário autenticado
// GET /api/settings/profile
func (c *UserController) GetProfile(ctx *fiber.Ctx) error {
	// 1. Obter ID do usuário do contexto (injetado pelo middleware de autenticação)
	userIDValue := ctx.Locals("userID")
	if userIDValue == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "User not authenticated",
		})
	}

	// O middleware armazena userID como string, então precisamos converter para UUID
	userIDStr, ok := userIDValue.(string)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "Invalid user ID in context",
		})
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "Invalid user ID format",
		})
	}

	// 2. Buscar usuário usando o use case
	user, err := c.getProfileUseCase.Execute(ctx.Context(), userID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
			Error: "User not found",
		})
	}

	// 3. Converter para response
	userResponse := response.UserResponse{
		ID:              user.ID().String(),
		Name:            user.Name(),
		Email:           user.Email(),
		Role:            user.Role(),
		Active:          user.IsActive(),
		ProfilePhotoKey: user.ProfilePhotoKey(),
		ProfilePhotoURL: "", // Não retornar URL pública - usar presigned URL
		CreatedAt:       user.CreatedAt(),
		UpdatedAt:       user.UpdatedAt(),
	}

	// Gerar presigned URL se houver foto
	if userResponse.ProfilePhotoKey != "" {
		presignedURL, err := c.getProfilePhotoPresignedURL(ctx.Context(), userResponse.ProfilePhotoKey)
		if err == nil {
			userResponse.ProfilePhotoURL = presignedURL
		}
	}

	// 4. Retornar resposta
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessGetProfileResponse(userResponse))
}

// UploadProfilePhoto faz upload da foto de perfil do usuário autenticado
// POST /api/settings/profile-photo
func (c *UserController) UploadProfilePhoto(ctx *fiber.Ctx) error {
	// 1. Obter ID do usuário do contexto (injetado pelo middleware de autenticação)
	userIDValue := ctx.Locals("userID")
	if userIDValue == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "User not authenticated",
		})
	}

	userIDStr, ok := userIDValue.(string)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "Invalid user ID in context",
		})
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "Invalid user ID format",
		})
	}

	// 2. Obter arquivo do form
	file, err := ctx.FormFile("photo")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "No photo provided",
		})
	}

	// 3. Executar use case
	input := users.UpdateProfilePhotoInput{
		UserID: userID,
		File:   file,
	}

	output, err := c.updateProfilePhotoUseCase.Execute(ctx.Context(), input)
	if err != nil {
		// Tratar erros de domínio específicos
		switch {
		case stderrors.Is(err, errors.ErrUserNotFound):
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "User not found",
			})
		case stderrors.Is(err, errors.ErrStorageNotConfigured):
			return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
				Error: "Storage is not configured",
			})
		case stderrors.Is(err, errors.ErrFileTooLarge):
			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Error: "File too large (max 5MB)",
			})
		case stderrors.Is(err, errors.ErrInvalidFileType):
			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Error: "Invalid file type (only jpg, jpeg, png, gif allowed)",
			})
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
				Error: "Failed to upload profile photo",
			})
		}
	}

	// 4. Retornar resposta de sucesso (retorna a key, não a URL)
	return ctx.Status(fiber.StatusOK).JSON(response.UploadProfilePhotoResponse{
		Key: output.Key,
	})
}

// RemoveProfilePhoto remove a foto de perfil do usuário autenticado
// DELETE /api/settings/profile-photo
func (c *UserController) RemoveProfilePhoto(ctx *fiber.Ctx) error {
	// 1. Obter ID do usuário do contexto (injetado pelo middleware de autenticação)
	userIDValue := ctx.Locals("userID")
	if userIDValue == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "User not authenticated",
		})
	}

	userIDStr, ok := userIDValue.(string)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "Invalid user ID in context",
		})
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "Invalid user ID format",
		})
	}

	// 2. Executar use case
	input := users.RemoveProfilePhotoInput{
		UserID: userID,
	}

	output, err := c.removeProfilePhotoUseCase.Execute(ctx.Context(), input)
	if err != nil {
		// Tratar erros de domínio específicos
		switch {
		case stderrors.Is(err, errors.ErrUserNotFound):
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "User not found",
			})
		case stderrors.Is(err, errors.ErrStorageNotConfigured):
			return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
				Error: "Storage is not configured",
			})
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
				Error: "Failed to remove profile photo",
			})
		}
	}

	// 3. Retornar resposta de sucesso
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Success: output.Success,
		Message: "Profile photo removed successfully",
	})
}

// GetProfilePhotoPresignedURL retorna uma presigned URL para acessar a foto de perfil
// GET /api/settings/profile-photo/url
func (c *UserController) GetProfilePhotoPresignedURL(ctx *fiber.Ctx) error {
	// 1. Obter ID do usuário do contexto (injetado pelo middleware de autenticação)
	userIDValue := ctx.Locals("userID")
	if userIDValue == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "User not authenticated",
		})
	}

	userIDStr, ok := userIDValue.(string)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "Invalid user ID in context",
		})
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "Invalid user ID format",
		})
	}

	// 2. Buscar usuário usando o use case
	user, err := c.getProfileUseCase.Execute(ctx.Context(), userID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
			Error: "User not found",
		})
	}

	// 3. Verificar se o usuário tem foto de perfil
	profilePhotoKey := user.ProfilePhotoKey()
	if profilePhotoKey == "" {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
			Error: "User has no profile photo",
		})
	}

	// 4. Gerar presigned URL
	presignedURL, err := c.getProfilePhotoPresignedURL(ctx.Context(), profilePhotoKey)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "Failed to generate presigned URL",
		})
	}

	// 5. Retornar presigned URL
	return ctx.Status(fiber.StatusOK).JSON(response.PresignedURLResponse{
		URL: presignedURL,
	})
}
