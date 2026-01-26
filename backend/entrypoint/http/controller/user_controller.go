package controller

import (
	stderrors "errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/application/usecase/users"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// UserController handles HTTP requests for users
type UserController struct {
	listUsersUseCase      *users.ListUsersUseCase
	changePasswordUseCase *users.ChangePasswordUseCase
	updateProfileUseCase  *users.UpdateProfile
	getProfileUseCase     *users.GetProfile
	mapper                *UserMapper
}

// NewUserController creates a new UserController instance
func NewUserController(
	listUsersUseCase *users.ListUsersUseCase,
	changePasswordUseCase *users.ChangePasswordUseCase,
	updateProfileUseCase *users.UpdateProfile,
	getProfileUseCase *users.GetProfile,
) *UserController {
	return &UserController{
		listUsersUseCase:      listUsersUseCase,
		changePasswordUseCase: changePasswordUseCase,
		updateProfileUseCase:  updateProfileUseCase,
		getProfileUseCase:     getProfileUseCase,
		mapper:                &UserMapper{},
	}
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
		ID:        user.ID().String(),
		Name:      user.Name(),
		Email:     user.Email(),
		Role:      user.Role(),
		Active:    user.IsActive(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
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
		ID:        user.ID().String(),
		Name:      user.Name(),
		Email:     user.Email(),
		Role:      user.Role(),
		Active:    user.IsActive(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}

	// 4. Retornar resposta
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessGetProfileResponse(userResponse))
}
