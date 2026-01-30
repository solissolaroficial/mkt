package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
	"github.com/seu-usuario/solis-backend/entrypoint/http/middleware"
)

// SetupKpiRoutes registra as rotas de KPIs (protegidas)
func SetupKpiRoutes(router fiber.Router, kpiController *controller.KpiController) {
	kpis := router.Group("/kpis")

	// CRUD de KPIs
	kpis.Get("/", kpiController.List)
	kpis.Post("/", kpiController.Create)
	kpis.Get("/:id", kpiController.GetByID)
	kpis.Delete("/:id", middleware.RequireAdmin(), kpiController.Delete)

	// Buscar KPIs por slugs
	kpis.Post("/by-slugs", kpiController.GetBySlugs)

	// Monthly data
	kpis.Put("/:kpiId/monthly-data", kpiController.UpdateMonthlyData)
	kpis.Delete("/:kpiId/monthly-data/:monthlyDataId", middleware.RequireAdmin(), kpiController.DeleteMonthlyData)
}
