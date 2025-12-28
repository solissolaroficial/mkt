package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/seu-usuario/solis-backend/application/usecase/auth"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// AuthController handles HTTP requests for authentication
type AuthController struct {
	loginUseCase *auth.LoginUseCase
	mapper       *AuthMapper
}

// NewAuthController creates a new AuthController instance
func NewAuthController(loginUseCase *auth.LoginUseCase) *AuthController {
	return &AuthController{
		loginUseCase: loginUseCase,
		mapper:       NewAuthMapper(),
	}
}

// Login handles user authentication
// @Summary Login user
// @Description Authenticate a user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param loginRequest body request.LoginRequest true "Login credentials"
// @Success 200 {object} response.LoginResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /auth/login [post]
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	// Parse request body
	var req request.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Execute use case
	input := auth.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	}

	output, err := c.loginUseCase.Execute(context.Background(), input)
	if err != nil {
		// Handle domain errors
		switch err {
		case errors.ErrInvalidCredentials:
			return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
				Error: "Invalid credentials",
			})
		case errors.ErrUserInactive:
			return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
				Error: "User account is inactive",
			})
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
				Error: "Internal server error",
			})
		}
	}

	// Convert to response
	loginResponse := c.mapper.ToLoginResponse(output)

	return ctx.Status(fiber.StatusOK).JSON(loginResponse)
}
