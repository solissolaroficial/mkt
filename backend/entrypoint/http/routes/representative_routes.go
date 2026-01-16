package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
)

// SetupRepresentativeRoutes configura as rotas para Representatives
func SetupRepresentativeRoutes(router fiber.Router, representativeController *controller.RepresentativeController) {
	api := router.Group("/v1")

	// Rotas de Representatives
	representatives := api.Group("/representatives")

	// CRUD básico
	representatives.Post("/", representativeController.Create)
	representatives.Get("/", representativeController.List)

	// Operações adicionais (rotas específicas devem vir ANTES de /:id)
	representatives.Get("/table", representativeController.GetTableData)
	representatives.Get("/profiles", representativeController.GetAllProfiles)
	representatives.Get("/profile/:name", representativeController.GetProfile)

	// Rotas com parâmetro (devem vir DEPOIS das rotas específicas)
	representatives.Get("/:id", representativeController.GetByID)
	representatives.Get("/:id/stats", representativeController.GetStats)
	representatives.Put("/:id", representativeController.Update)
	representatives.Delete("/:id", representativeController.Delete)
}
