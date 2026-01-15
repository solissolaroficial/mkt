package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
)

// SetupRepresentativeMonthlyGoalRoutes sets up routes for representative monthly goals
func SetupRepresentativeMonthlyGoalRoutes(router fiber.Router, controller *controller.RepresentativeMonthlyGoalController) {
	api := router.Group("/representative-monthly-goals")

	// CRUD operations
	api.Post("/", controller.Create)
	api.Get("/:id", controller.GetByID)
	api.Put("/:id", controller.Update)
	api.Delete("/:id", controller.Delete)
	api.Get("/", controller.List)

	// Table data endpoint for transposed view
	api.Get("/table/data", controller.GetTableData)
}
