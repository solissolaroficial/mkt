package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
	"github.com/seu-usuario/solis-backend/entrypoint/http/middleware"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

func SetupCredentialRoutes(
	router fiber.Router,
	controller *controller.ProgramCredentialController,
) {
	credentials := router.Group("/settings/credentials")

	// Apply rate limiting (100 requests per minute)
	credentials.Use(limiter.New(limiter.Config{
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

	// Read routes (any authenticated user)
	credentials.Get("/", controller.List)
	credentials.Get("/:id", controller.Get)

	// Admin routes (requires marketing role)
	admin := credentials.Group("")
	admin.Use(middleware.RequireMarketing())

	admin.Post("/", controller.Create)
	admin.Put("/:id", controller.Update)
	admin.Delete("/:id", controller.Delete)
}
