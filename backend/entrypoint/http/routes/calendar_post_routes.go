package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
	"github.com/seu-usuario/solis-backend/entrypoint/http/middleware"
)

// SetupCalendarPostRoutes configura as rotas de CalendarPosts
// Requer role admin ou marketing para todas as operações
// Aplica rate limiting de 100 requisições por minuto
func SetupCalendarPostRoutes(router fiber.Router, calendarPostController *controller.CalendarPostController) {
	calendarPosts := router.Group("/calendar-posts")

	// Aplicar rate limiting (100 requisições por minuto)
	calendarPosts.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Rate limiting por IP
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests, please try again later",
			})
		},
	}))

	// Aplicar middleware de autorização (admin ou marketing)
	calendarPosts.Use(middleware.RequireMarketing())

	// CRUD de CalendarPosts
	calendarPosts.Post("/", calendarPostController.Create)
	calendarPosts.Get("/:id", calendarPostController.Get)
	calendarPosts.Put("/:id", calendarPostController.Update)
	calendarPosts.Delete("/:id", calendarPostController.Delete)
	calendarPosts.Get("/", calendarPostController.List)

	// Rotas especiais de CalendarPosts
	calendarPosts.Put("/:id/status", calendarPostController.UpdateStatus)
	calendarPosts.Put("/:id/confirm-publishing", calendarPostController.ConfirmPublishing)
}
