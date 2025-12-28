package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// NewCorsMiddleware returns a configured CORS middleware
func NewCorsMiddleware() fiber.Handler {
	// Ler origens permitidas de variável de ambiente
	allowOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if allowOrigins == "" {
		// Tentar obter do FRONTEND_URL como fallback
		frontendURL := os.Getenv("FRONTEND_URL")
		if frontendURL != "" {
			allowOrigins = frontendURL
		} else {
			// Apenas para desenvolvimento
			allowOrigins = "*"
		}
	}

	// Configurar o middleware CORS
	return cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length",
		MaxAge:           3600, // cache de preflight em segundos
	})
}
