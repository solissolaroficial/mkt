package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
	"github.com/seu-usuario/solis-backend/entrypoint/http/middleware"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

func SetupGiftRoutes(
	router fiber.Router,
	giftItemController *controller.GiftItemController,
	giftTransactionController *controller.GiftTransactionController,
) {
	gifts := router.Group("/gifts")

	// Aplicar rate limiting (100 requisições por minuto)
	gifts.Use(limiter.New(limiter.Config{
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

	// Rotas de Gift Items
	items := gifts.Group("/items")

	// Rotas de leitura (qualquer usuário autenticado)
	items.Get("/", giftItemController.List)
	items.Get("/:id", giftItemController.GetByID)

	// Rotas administrativas (requer role admin ou marketing)
	itemsAdmin := items.Group("")
	itemsAdmin.Use(middleware.RequireMarketing())

	itemsAdmin.Post("/", giftItemController.Create)
	itemsAdmin.Put("/:id", giftItemController.Update)
	itemsAdmin.Delete("/:id", giftItemController.Delete)

	// Rotas de Gift Transactions
	transactions := gifts.Group("/transactions")

	// Rotas de leitura (qualquer usuário autenticado)
	transactions.Get("/", giftTransactionController.List)
	transactions.Get("/:id", giftTransactionController.GetByID)

	// Rotas administrativas (requer role admin ou marketing)
	transactionsAdmin := transactions.Group("")
	transactionsAdmin.Use(middleware.RequireMarketing())

	transactionsAdmin.Post("/", giftTransactionController.Create)
	transactionsAdmin.Put("/:id", giftTransactionController.Update)
	transactionsAdmin.Delete("/:id", giftTransactionController.Delete)
}
