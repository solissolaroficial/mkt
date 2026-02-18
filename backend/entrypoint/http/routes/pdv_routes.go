package routes

import (
	"time"

	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
	"github.com/seu-usuario/solis-backend/entrypoint/http/middleware"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// userBasedKeyGenerator generates a rate limit key based on user ID or falls back to IP
func userBasedKeyGenerator(c *fiber.Ctx) string {
	userID := c.Locals("userID")
	if userID != nil {
		if id, ok := userID.(string); ok {
			return id
		}
	}
	return c.IP()
}

// SetupPdvRoutes configura as rotas do PDV
func SetupPdvRoutes(
	router fiber.Router,
	controller *controller.PdvController,
) {
	pdvPosts := router.Group("/pdv/posts")
	recurrentPdvs := router.Group("/pdv/recurrent")

	// Aplicar middleware de autorização (marketing)
	pdvPosts.Use(middleware.RequireMarketing())
	recurrentPdvs.Use(middleware.RequireMarketing())

	// Rotas especiais de PDV Posts (mais específicas antes das genéricas)
	// Rate limiting para operações de escrita (20 req/min)
	pdvPosts.Patch("/:id/status",
		limiter.New(limiter.Config{
			Max:          20,
			Expiration:   1 * time.Minute,
			KeyGenerator: userBasedKeyGenerator,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(response.ErrorResponse{
					Error: "Too many requests, please try again later",
				})
			},
		}),
		controller.UpdatePdvPostStatus,
	)

	// CRUD de PDV Posts (rotas genéricas)
	// Rate limiting para operações de escrita (20 req/min)
	pdvPosts.Post("/",
		limiter.New(limiter.Config{
			Max:          20,
			Expiration:   1 * time.Minute,
			KeyGenerator: userBasedKeyGenerator,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(response.ErrorResponse{
					Error: "Too many requests, please try again later",
				})
			},
		}),
		controller.CreatePdvPost,
	)

	pdvPosts.Put("/:id",
		limiter.New(limiter.Config{
			Max:          20,
			Expiration:   1 * time.Minute,
			KeyGenerator: userBasedKeyGenerator,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(response.ErrorResponse{
					Error: "Too many requests, please try again later",
				})
			},
		}),
		controller.UpdatePdvPost,
	)

	pdvPosts.Delete("/:id",
		limiter.New(limiter.Config{
			Max:        20,
			Expiration: 1 * time.Minute,
			KeyGenerator: func(c *fiber.Ctx) string {
				userID := c.Locals("userID")
				if userID != nil {
					if id, ok := userID.(string); ok {
						return id
					}
				}
				return c.IP()
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(response.ErrorResponse{
					Error: "Too many requests, please try again later",
				})
			},
		}),
		controller.DeletePdvPost,
	)

	// Rate limiting para operações de leitura (100 req/min)
	// NOTA: GET /:id (específico) antes de GET / (genérico)
	pdvPosts.Get("/:id",
		limiter.New(limiter.Config{
			Max:          100,
			Expiration:   1 * time.Minute,
			KeyGenerator: userBasedKeyGenerator,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(response.ErrorResponse{
					Error: "Too many requests, please try again later",
				})
			},
		}),
		controller.GetPdvPost,
	)

	pdvPosts.Get("/",
		limiter.New(limiter.Config{
			Max:          100,
			Expiration:   1 * time.Minute,
			KeyGenerator: userBasedKeyGenerator,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(response.ErrorResponse{
					Error: "Too many requests, please try again later",
				})
			},
		}),
		controller.ListPdvPosts,
	)

	// CRUD de PDVs Recorrentes
	// Rate limiting para operações de escrita (20 req/min)
	recurrentPdvs.Post("/",
		limiter.New(limiter.Config{
			Max:          20,
			Expiration:   1 * time.Minute,
			KeyGenerator: userBasedKeyGenerator,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(response.ErrorResponse{
					Error: "Too many requests, please try again later",
				})
			},
		}),
		controller.CreateRecurrentPdv,
	)

	recurrentPdvs.Put("/:id",
		limiter.New(limiter.Config{
			Max:        20,
			Expiration: 1 * time.Minute,
			KeyGenerator: func(c *fiber.Ctx) string {
				userID := c.Locals("userID")
				if userID != nil {
					if id, ok := userID.(string); ok {
						return id
					}
				}
				return c.IP()
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(response.ErrorResponse{
					Error: "Too many requests, please try again later",
				})
			},
		}),
		controller.UpdateRecurrentPdv,
	)

	recurrentPdvs.Delete("/:id",
		limiter.New(limiter.Config{
			Max:          20,
			Expiration:   1 * time.Minute,
			KeyGenerator: userBasedKeyGenerator,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(response.ErrorResponse{
					Error: "Too many requests, please try again later",
				})
			},
		}),
		controller.DeleteRecurrentPdv,
	)

	// Rate limiting para operações de leitura (100 req/min)
	recurrentPdvs.Get("/:id",
		limiter.New(limiter.Config{
			Max:          100,
			Expiration:   1 * time.Minute,
			KeyGenerator: userBasedKeyGenerator,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(response.ErrorResponse{
					Error: "Too many requests, please try again later",
				})
			},
		}),
		controller.GetRecurrentPdv,
	)

	recurrentPdvs.Get("/",
		limiter.New(limiter.Config{
			Max:          100,
			Expiration:   1 * time.Minute,
			KeyGenerator: userBasedKeyGenerator,
			LimitReached: func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(response.ErrorResponse{
					Error: "Too many requests, please try again later",
				})
			},
		}),
		controller.ListRecurrentPdvs,
	)
}
