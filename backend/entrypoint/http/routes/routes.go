package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
	"github.com/seu-usuario/solis-backend/entrypoint/http/middleware"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// Controllers agrupa todos os controllers da aplicação
type Controllers struct {
	AuthController                      *controller.AuthController
	KpiController                       *controller.KpiController
	TaskController                      *controller.TaskController
	SubtaskController                   *controller.SubtaskController
	CommentController                   *controller.CommentController
	NotificationController              *controller.NotificationController
	UserController                      *controller.UserController
	CalendarPostController              *controller.CalendarPostController
	PdvController                       *controller.PdvController
	SocialController                    *controller.SocialController
	SocialPostController                *controller.SocialPostController
	OfflineActionController             *controller.OfflineActionController
	ShowroomItemController              *controller.ShowroomItemController
	RepMarketingActionController        *controller.RepMarketingActionController
	GiftItemController                  *controller.GiftItemController
	GiftTransactionController           *controller.GiftTransactionController
	AccountPayableController            *controller.AccountPayableController
	BudgetController                    *controller.BudgetController
	RepresentativeController            *controller.RepresentativeController
	RepresentativeMonthlyGoalController *controller.RepresentativeMonthlyGoalController
	BrandController                     *controller.BrandController
	FlowController                      *controller.FlowController
}

// Middlewares agrupa todos os middlewares da aplicação
type Middlewares struct {
	AuthMiddleware *middleware.AuthMiddleware
	CorsMiddleware fiber.Handler
}

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes(app *fiber.App, controllers *Controllers, middlewares *Middlewares) {
	// 1. CORS global (sempre primeiro)
	app.Use(middlewares.CorsMiddleware)

	// 2. Health check (público)
	app.Get("/health", healthCheck)

	// 3. API group
	api := app.Group("/api")

	// 4. Rotas públicas
	SetupAuthRoutes(api, controllers.AuthController)

	// 5. Rotas protegidas (com auth middleware)
	protected := api.Group("", middlewares.AuthMiddleware.Handle())
	SetupKpiRoutes(protected, controllers.KpiController)
	SetupTaskRoutes(protected, controllers.TaskController)
	SetupSubtaskRoutes(protected, controllers.SubtaskController)
	SetupCommentRoutes(protected, controllers.CommentController)
	SetupNotificationRoutes(protected, controllers.NotificationController)
	SetupUserRoutes(protected, controllers.UserController)
	SetupSettingsRoutes(protected, controllers.UserController)
	SetupCalendarPostRoutes(protected, controllers.CalendarPostController)
	SetupPdvRoutes(protected, controllers.PdvController)
	SetupSocialRoutes(protected, controllers.SocialController)
	RegisterSocialPostRoutes(protected, controllers.SocialPostController)
	SetupOfflineActionRoutes(protected, controllers.OfflineActionController)
	SetupShowroomItemRoutes(protected, controllers.ShowroomItemController)
	SetupRepMarketingActionRoutes(protected, controllers.RepMarketingActionController)
	SetupGiftRoutes(protected, controllers.GiftItemController, controllers.GiftTransactionController)
	SetupAccountPayableRoutes(protected, controllers.AccountPayableController)
	SetupBudgetRoutes(protected, controllers.BudgetController)
	SetupRepresentativeRoutes(protected, controllers.RepresentativeController)
	SetupRepresentativeMonthlyGoalRoutes(protected, controllers.RepresentativeMonthlyGoalController)
	SetupBrandRoutes(protected, controllers.BrandController)
	SetupFlowRoutes(protected, controllers.FlowController)

	// 6. 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
			Error: "Route not found",
		})
	})
}

// healthCheck é um handler simples para verificar se a API está rodando
func healthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"service": "solis-mkt-api",
	})
}
