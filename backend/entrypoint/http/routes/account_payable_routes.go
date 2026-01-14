package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
	"github.com/seu-usuario/solis-backend/entrypoint/http/middleware"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

func SetupAccountPayableRoutes(
	router fiber.Router,
	accountPayableController *controller.AccountPayableController,
) {
	financial := router.Group("/financial")

	// Aplicar rate limiting (100 requisições por minuto)
	financial.Use(limiter.New(limiter.Config{
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

	// Rotas de Accounts Payable
	accountsPayable := financial.Group("/accounts-payable")

	// Rotas de leitura (qualquer usuário autenticado)
	accountsPayable.Get("/", accountPayableController.List)
	accountsPayable.Get("/:id", accountPayableController.GetByID)

	// Rotas administrativas (requer role admin ou marketing)
	accountsPayableAdmin := accountsPayable.Group("")
	accountsPayableAdmin.Use(middleware.RequireMarketing())

	accountsPayableAdmin.Post("/", accountPayableController.Create)
	accountsPayableAdmin.Put("/:id", accountPayableController.Update)
	accountsPayableAdmin.Delete("/:id", accountPayableController.Delete)
	accountsPayableAdmin.Patch("/:id/nf", accountPayableController.ToggleNF)
	accountsPayableAdmin.Patch("/:id/boleto", accountPayableController.ToggleBoleto)
	accountsPayableAdmin.Post("/:id/send-to-finance", accountPayableController.SendToFinance)
}
