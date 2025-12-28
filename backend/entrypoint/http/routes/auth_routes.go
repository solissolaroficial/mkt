package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
)

// SetupAuthRoutes registra as rotas de autenticação (públicas)
func SetupAuthRoutes(router fiber.Router, authController *controller.AuthController) {
	auth := router.Group("/auth")

	// POST /api/auth/login
	auth.Post("/login", authController.Login)
}
