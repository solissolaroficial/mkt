package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
)

// SetupFlowRoutes configura as rotas de fluxo
func SetupFlowRoutes(app fiber.Router, flowController *controller.FlowController) {
	flow := app.Group("/flows")

	// CRUD de fluxos
	flow.Post("", flowController.CreateFlow)
	flow.Get("", flowController.ListFlows)
	flow.Get("/:id", flowController.GetFlow)
	flow.Put("/:id", flowController.UpdateFlow)
	flow.Delete("/:id", flowController.DeleteFlow)

	// Reordenar fluxos
	flow.Post("/reorder", flowController.ReorderFlows)
}
