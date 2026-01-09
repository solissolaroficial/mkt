package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
	"github.com/seu-usuario/solis-backend/entrypoint/http/middleware"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

func SetupSocialRoutes(
	router fiber.Router,
	controller *controller.SocialController,
) {
	social := router.Group("/social/benchmarking")

	// Aplicar rate limiting (100 requisições por minuto)
	social.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(response.ErrorResponse{
				Error: "Too many requests, please try again later",
			})
		},
	}))

	// Rotas de leitura (qualquer usuário autenticado)
	social.Get("/", controller.List)
	social.Get("/:id", controller.GetByID)

	// Rotas administrativas (requer role admin ou marketing)
	admin := social.Group("")
	admin.Use(middleware.RequireMarketing())

	admin.Post("/", controller.Create)
	admin.Put("/:id", controller.Update)
	admin.Delete("/:id", controller.Delete)
}
