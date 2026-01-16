package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
)

// SetupBudgetRoutes configura as rotas para Budget
func SetupBudgetRoutes(router fiber.Router, budgetController *controller.BudgetController) {
	// Rotas de Budget
	budget := router.Group("/budget")

	// CRUD básico
	budget.Post("/", budgetController.Create)
	budget.Get("/", budgetController.List)

	// Operações adicionais (rotas específicas devem vir ANTES de /:id)
	budget.Post("/batch", budgetController.BatchCreate)
	budget.Get("/summary", budgetController.GetSummary)
	budget.Get("/years", budgetController.GetDistinctYears)

	// Rotas com parâmetro (devem vir DEPOIS das rotas específicas)
	budget.Get("/:id", budgetController.GetByID)
	budget.Put("/:id", budgetController.Update)
	budget.Delete("/:id", budgetController.Delete)
}
