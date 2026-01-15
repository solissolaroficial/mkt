package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
)

// SetupBudgetRoutes configura as rotas para Budget
func SetupBudgetRoutes(router fiber.Router, budgetController *controller.BudgetController) {
	api := router.Group("/v1")

	// Rotas de Budget
	budget := api.Group("/budget")

	// CRUD básico
	budget.Post("/", budgetController.Create)
	budget.Get("/", budgetController.List)
	budget.Get("/:id", budgetController.GetByID)
	budget.Put("/:id", budgetController.Update)
	budget.Delete("/:id", budgetController.Delete)

	// Operações adicionais
	budget.Post("/batch", budgetController.BatchCreate)
	budget.Get("/summary", budgetController.GetSummary)
	budget.Get("/years", budgetController.GetDistinctYears)
}
