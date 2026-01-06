package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// RequireRole cria um middleware que verifica se o usuário possui uma das roles permitidas
// Roles válidas: admin, marketing, commercial
func RequireRole(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Obter role do usuário do contexto (injetado pelo AuthMiddleware)
		userRole, ok := c.Locals("userRole").(string)
		if !ok || userRole == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not authenticated",
			})
		}

		// Verificar se a role do usuário está na lista de roles permitidas
		for _, allowedRole := range roles {
			if userRole == allowedRole {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Insufficient permissions",
		})
	}
}

// RequireAdmin cria um middleware que exige que o usuário seja admin
func RequireAdmin() fiber.Handler {
	return RequireRole("admin")
}

// RequireMarketing cria um middleware que exige que o usuário seja admin ou marketing
func RequireMarketing() fiber.Handler {
	return RequireRole("admin", "marketing")
}
