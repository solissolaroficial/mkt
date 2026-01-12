package routes

import (
	"time"

	"solis/backend/entrypoint/http/controller"
	"solis/backend/entrypoint/http/middleware"
	"solis/backend/entrypoint/http/payload/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func SetupShowroomItemRoutes(
	router fiber.Router,
	controller *controller.ShowroomItemController,
) {
	showroomItem := router.Group("/cooperative/showroom-items")

	// Aplicar rate limiting (100 requisições por minuto)
	showroomItem.Use(limiter.New(limiter.Config{
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
	showroomItem.Get("/", controller.List)
	showroomItem.Get("/:id", controller.GetByID)

	// Rotas administrativas (requer role admin ou marketing)
	admin := showroomItem.Group("")
	admin.Use(middleware.RequireMarketing())

	admin.Post("/", controller.Create)
	admin.Put("/:id", controller.Update)
	admin.Delete("/:id", controller.Delete)
}
