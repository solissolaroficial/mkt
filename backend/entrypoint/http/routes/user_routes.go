package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
)

// SetupUserRoutes configures all user-related routes
func SetupUserRoutes(router fiber.Router, userController *controller.UserController) {
	users := router.Group("/users")

	// List all active users (protected by auth middleware)
	users.Get("/", userController.ListUsers)

	// List all active users without pagination (protected by auth middleware)
	users.Get("/all", userController.ListAllUsers)
}

// SetupSettingsRoutes configura as rotas de configurações do usuário
func SetupSettingsRoutes(router fiber.Router, userController *controller.UserController) {
	settings := router.Group("/settings")

	// Get user profile (protected by auth middleware)
	settings.Get("/profile", userController.GetProfile)

	// Change user password (protected by auth middleware)
	settings.Put("/password", userController.ChangePassword)

	// Update user profile (protected by auth middleware)
	settings.Put("/profile", userController.UpdateProfile)
}
