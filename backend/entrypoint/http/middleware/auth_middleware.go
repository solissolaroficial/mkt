package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/seu-usuario/solis-backend/core/domain/service"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// AuthMiddleware struct for JWT authentication
type AuthMiddleware struct {
	jwtService service.JWTService
}

// NewAuthMiddleware creates a new instance of AuthMiddleware
func NewAuthMiddleware(jwtService service.JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

// Handle returns a fiber.Handler for JWT authentication
func (m *AuthMiddleware) Handle() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 1. Extrair token do header Authorization
		authHeader := c.Get("Authorization")

		// 2. Validar formato
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
				Error: "Authorization header is required",
			})
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
				Error: "Invalid authorization header format. Expected 'Bearer <token>'",
			})
		}

		// 3. Extrair token (remover prefixo "Bearer ")
		token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		// 4. Validar token usando jwtService.ValidateToken(token, service.AccessToken)
		claims, err := m.jwtService.ValidateToken(token, service.AccessToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
				Error: "Invalid or expired token",
			})
		}

		// 5. Armazenar dados no context do Fiber
		c.Locals("userID", claims.UserID.String())
		c.Locals("userEmail", claims.Email)
		c.Locals("userRole", claims.Role)

		// 6. Continuar para próximo handler
		return c.Next()
	}
}

// GetUserID extracts user ID from fiber context
func GetUserID(c *fiber.Ctx) (uuid.UUID, error) {
	userIDStr, ok := c.Locals("userID").(string)
	if !ok {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "User ID not found in context")
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid user ID format")
	}

	return userID, nil
}

// GetUserEmail extracts user email from fiber context
func GetUserEmail(c *fiber.Ctx) (string, error) {
	email, ok := c.Locals("userEmail").(string)
	if !ok {
		return "", fiber.NewError(fiber.StatusUnauthorized, "User email not found in context")
	}

	return email, nil
}

// GetUserRole extracts user role from fiber context
func GetUserRole(c *fiber.Ctx) (string, error) {
	role, ok := c.Locals("userRole").(string)
	if !ok {
		return "", fiber.NewError(fiber.StatusUnauthorized, "User role not found in context")
	}

	return role, nil
}
