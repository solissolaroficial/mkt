package controller

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	flows "github.com/seu-usuario/solis-backend/application/usecase/flows"
	flowerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// FlowController handles HTTP requests for flows
type FlowController struct {
	createFlowUseCase   *flows.CreateFlow
	updateFlowUseCase   *flows.UpdateFlow
	deleteFlowUseCase   *flows.DeleteFlow
	getFlowUseCase      *flows.GetFlow
	listFlowsUseCase    *flows.ListFlows
	reorderFlowsUseCase *flows.ReorderFlows
}

// NewFlowController creates a new FlowController instance
func NewFlowController(
	createFlowUseCase *flows.CreateFlow,
	updateFlowUseCase *flows.UpdateFlow,
	deleteFlowUseCase *flows.DeleteFlow,
	getFlowUseCase *flows.GetFlow,
	listFlowsUseCase *flows.ListFlows,
	reorderFlowsUseCase *flows.ReorderFlows,
) *FlowController {
	return &FlowController{
		createFlowUseCase:   createFlowUseCase,
		updateFlowUseCase:   updateFlowUseCase,
		deleteFlowUseCase:   deleteFlowUseCase,
		getFlowUseCase:      getFlowUseCase,
		listFlowsUseCase:    listFlowsUseCase,
		reorderFlowsUseCase: reorderFlowsUseCase,
	}
}

// CreateFlow cria um novo fluxo
// POST /api/flows
func (c *FlowController) CreateFlow(ctx *fiber.Ctx) error {
	log.Printf("CreateFlow called")
	// 1. Parse request body
	var req request.CreateFlowRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Printf("CreateFlow: Error parsing request body: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	log.Printf("CreateFlow: Request body parsed: %+v", req)

	// 2. Executar use case
	flow, err := c.createFlowUseCase.Execute(
		ctx.Context(),
		req.Name,
		req.Description,
		req.Color,
		req.SortOrder,
	)
	if err != nil {
		log.Printf("CreateFlow: Error creating flow: %v", err)
		// Tratar erros de domínio específicos
		switch {
		case err == flowerrors.ErrFlowEmptyName:
			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Error: "Flow name is required",
			})
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
				Error: "Failed to create flow",
			})
		}
	}

	log.Printf("CreateFlow: Flow created successfully: ID=%s, Name=%s", flow.ID(), flow.Name())

	// 3. Converter para response
	flowResponse := response.FlowResponse{
		UUID:        flow.ID().String(),
		Name:        flow.Name(),
		Description: flow.Description(),
		Color:       flow.Color(),
		SortOrder:   flow.SortOrder(),
		CreatedAt:   flow.CreatedAt().Format(time.RFC3339),
		UpdatedAt:   flow.UpdatedAt().Format(time.RFC3339),
	}

	// 4. Retornar resposta de sucesso
	return ctx.Status(fiber.StatusCreated).JSON(flowResponse)
}

// GetFlow busca um fluxo por ID
// GET /api/flows/:id
func (c *FlowController) GetFlow(ctx *fiber.Ctx) error {
	log.Printf("GetFlow called with ID: %s", ctx.Params("id"))
	// 1. Obter ID da URL
	flowID := ctx.Params("id")
	if flowID == "" {
		log.Printf("GetFlow: Flow ID is required")
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Flow ID is required",
		})
	}

	// 2. Executar use case
	flow, err := c.getFlowUseCase.Execute(ctx.Context(), flowID)
	if err != nil {
		log.Printf("GetFlow: Error getting flow: %v", err)
		// Tratar erros de domínio específicos
		switch {
		case err == flowerrors.ErrFlowNotFound:
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Flow not found",
			})
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
				Error: "Failed to get flow",
			})
		}
	}

	log.Printf("GetFlow: Flow retrieved successfully: ID=%s, Name=%s", flow.ID(), flow.Name())

	// 3. Converter para response
	flowResponse := response.FlowResponse{
		UUID:        flow.ID().String(),
		Name:        flow.Name(),
		Description: flow.Description(),
		Color:       flow.Color(),
		SortOrder:   flow.SortOrder(),
		CreatedAt:   flow.CreatedAt().Format(time.RFC3339),
		UpdatedAt:   flow.UpdatedAt().Format(time.RFC3339),
	}

	// 4. Retornar resposta
	return ctx.Status(fiber.StatusOK).JSON(flowResponse)
}

// ListFlows lista todos os fluxos
// GET /api/flows
func (c *FlowController) ListFlows(ctx *fiber.Ctx) error {
	log.Printf("ListFlows called")
	// 1. Executar use case
	flows, err := c.listFlowsUseCase.Execute(ctx.Context())
	if err != nil {
		log.Printf("ListFlows: Error listing flows: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "Failed to list flows",
		})
	}

	log.Printf("ListFlows: Flows retrieved successfully: %d flows", len(flows))

	// 2. Converter para response
	flowResponses := make([]response.FlowResponse, len(flows))
	for i, flow := range flows {
		flowResponses[i] = response.FlowResponse{
			UUID:        flow.ID().String(),
			Name:        flow.Name(),
			Description: flow.Description(),
			Color:       flow.Color(),
			SortOrder:   flow.SortOrder(),
			CreatedAt:   flow.CreatedAt().Format(time.RFC3339),
			UpdatedAt:   flow.UpdatedAt().Format(time.RFC3339),
		}
	}

	// 3. Retornar resposta com paginação
	return ctx.Status(fiber.StatusOK).JSON(response.FlowListResponse{
		Data:       flowResponses,
		Total:      int64(len(flowResponses)),
		Page:       1,
		Limit:      100,
		TotalPages: 1,
	})
}

// UpdateFlow atualiza um fluxo
// PUT /api/flows/:id
func (c *FlowController) UpdateFlow(ctx *fiber.Ctx) error {
	log.Printf("UpdateFlow called with ID: %s", ctx.Params("id"))
	// 1. Obter ID da URL
	flowID := ctx.Params("id")
	if flowID == "" {
		log.Printf("UpdateFlow: Flow ID is required")
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Flow ID is required",
		})
	}

	// 2. Parse request body
	var req request.UpdateFlowRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Printf("UpdateFlow: Error parsing request body: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	log.Printf("UpdateFlow: Request body parsed: %+v", req)

	// 3. Executar use case
	var sortOrder int
	if req.SortOrder != nil {
		sortOrder = *req.SortOrder
	}

	// Handle name field (now optional)
	var name string
	if req.Name != nil {
		name = *req.Name
	}

	flow, err := c.updateFlowUseCase.Execute(
		ctx.Context(),
		flowID,
		name,
		req.Description,
		req.Color,
		sortOrder,
	)
	if err != nil {
		log.Printf("UpdateFlow: Error updating flow: %v", err)
		// Tratar erros de domínio específicos
		switch {
		case err == flowerrors.ErrFlowNotFound:
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Flow not found",
			})
		case err == flowerrors.ErrFlowEmptyName:
			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Error: "Flow name is required",
			})
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
				Error: "Failed to update flow",
			})
		}
	}

	log.Printf("UpdateFlow: Flow updated successfully: ID=%s, Name=%s", flow.ID(), flow.Name())

	// 4. Converter para response
	flowResponse := response.FlowResponse{
		UUID:        flow.ID().String(),
		Name:        flow.Name(),
		Description: flow.Description(),
		Color:       flow.Color(),
		SortOrder:   flow.SortOrder(),
		CreatedAt:   flow.CreatedAt().Format(time.RFC3339),
		UpdatedAt:   flow.UpdatedAt().Format(time.RFC3339),
	}

	// 5. Retornar resposta de sucesso
	return ctx.Status(fiber.StatusOK).JSON(flowResponse)
}

// DeleteFlow remove um fluxo
// DELETE /api/flows/:id
func (c *FlowController) DeleteFlow(ctx *fiber.Ctx) error {
	log.Printf("DeleteFlow called with ID: %s", ctx.Params("id"))
	// 1. Obter ID da URL
	flowID := ctx.Params("id")
	if flowID == "" {
		log.Printf("DeleteFlow: Flow ID is required")
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Flow ID is required",
		})
	}

	// 2. Executar use case
	err := c.deleteFlowUseCase.Execute(ctx.Context(), flowID)
	if err != nil {
		log.Printf("DeleteFlow: Error deleting flow: %v", err)
		// Tratar erros de domínio específicos
		switch {
		case err == flowerrors.ErrFlowNotFound:
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Flow not found",
			})
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
				Error: "Failed to delete flow",
			})
		}
	}

	log.Printf("DeleteFlow: Flow deleted successfully: ID=%s", flowID)

	// 3. Retornar resposta de sucesso
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Success: true,
		Message: "Flow deleted successfully",
	})
}

// ReorderFlows reordena os fluxos
// POST /api/flows/reorder
func (c *FlowController) ReorderFlows(ctx *fiber.Ctx) error {
	log.Printf("ReorderFlows called")
	// 1. Parse request body
	var req request.ReorderFlowsRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Printf("ReorderFlows: Error parsing request body: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	log.Printf("ReorderFlows: Request body parsed: %+v", req)

	// 2. Executar use case
	err := c.reorderFlowsUseCase.Execute(ctx.Context(), req.FlowIDs)
	if err != nil {
		log.Printf("ReorderFlows: Error reordering flows: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "Failed to reorder flows",
		})
	}

	log.Printf("ReorderFlows: Flows reordered successfully: %d flows", len(req.FlowIDs))

	// 3. Retornar resposta de sucesso
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Success: true,
		Message: "Flows reordered successfully",
	})
}
