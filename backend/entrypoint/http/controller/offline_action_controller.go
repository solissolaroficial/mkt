package controller

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	usecase "github.com/seu-usuario/solis-backend/application/usecase/cooperative/offlineaction"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

type OfflineActionController struct {
	createUseCase *usecase.CreateOfflineActionUseCase
	listUseCase   *usecase.ListOfflineActionsUseCase
	getUseCase    *usecase.GetOfflineActionUseCase
	updateUseCase *usecase.UpdateOfflineActionUseCase
	deleteUseCase *usecase.DeleteOfflineActionUseCase
	mapper        *response.OfflineActionPayloadMapper
}

func NewOfflineActionController(
	create *usecase.CreateOfflineActionUseCase,
	list *usecase.ListOfflineActionsUseCase,
	get *usecase.GetOfflineActionUseCase,
	update *usecase.UpdateOfflineActionUseCase,
	delete *usecase.DeleteOfflineActionUseCase,
) *OfflineActionController {
	return &OfflineActionController{
		createUseCase: create,
		listUseCase:   list,
		getUseCase:    get,
		updateUseCase: update,
		deleteUseCase: delete,
		mapper:        response.NewOfflineActionPayloadMapper(),
	}
}

// Create
func (c *OfflineActionController) Create(ctx *fiber.Ctx) error {
	var req request.CreateOfflineActionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case (validação está no use case)
	input := usecase.CreateOfflineActionInput{
		RequestedAmount:    req.RequestedAmount,
		ActionDate:         req.ActionDate,
		Category:           req.Category,
		PDV:                req.PDV,
		RepresentativeUUID: req.RepresentativeUUID,
		Observation:        req.Observation,
	}

	action, err := c.createUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error creating offline action: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Offline action created successfully: %s", action.ID())
	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.ToOfflineActionResponse(action))
}

// List
func (c *OfflineActionController) List(ctx *fiber.Ctx) error {
	var query request.ListOfflineActionsQuery
	if err := ctx.QueryParser(&query); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid query parameters",
		})
	}

	// Valores padrão para paginação
	if query.Page < 1 {
		query.Page = 1
	}
	if query.Limit < 1 || query.Limit > 100 {
		query.Limit = 10
	}

	// Validar sort_by se fornecido
	if query.SortBy != nil {
		validSortBy := map[string]bool{
			"action_date": true,
			"created_at":  true,
		}
		if !validSortBy[*query.SortBy] {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Error: "Invalid sort_by value. Valid values: action_date, created_at",
			})
		}
	}

	// Executar use case
	input := usecase.ListOfflineActionsInput{
		Category:           query.Category,
		RepresentativeUUID: query.RepresentativeUUID,
		Month:              query.Month,
		StartDate:          query.StartDate,
		EndDate:            query.EndDate,
		Status:             query.Status,
		Page:               query.Page,
		Limit:              query.Limit,
		SortBy:             query.SortBy,
		SortOrder:          query.SortOrder,
	}

	actions, total, err := c.listUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error listing offline actions: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Offline actions listed successfully: page=%d, limit=%d, total=%d", query.Page, query.Limit, total)
	return ctx.JSON(c.mapper.ToOfflineActionsListResponse(actions, total, query.Page, query.Limit))
}

// GetByID
func (c *OfflineActionController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	action, err := c.getUseCase.Execute(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrOfflineActionNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Offline action not found",
			})
		}
		log.Printf("Error getting offline action %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToOfflineActionResponse(action))
}

// Update
func (c *OfflineActionController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req request.UpdateOfflineActionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case
	input := usecase.UpdateOfflineActionInput{
		ID:                 id,
		ApprovedAmount:     req.ApprovedAmount,
		OrderNumber:        req.OrderNumber,
		DepartureDate:      req.DepartureDate,
		DeliveryForecast:   req.DeliveryForecast,
		DeliveryDate:       req.DeliveryDate,
		City:               req.City,
		UF:                 req.UF,
		Scored:             req.Scored,
		Status:             req.Status,
		Observation:        req.Observation,
		PDV:                req.PDV,
		RepresentativeUUID: req.RepresentativeUUID,
	}

	action, err := c.updateUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, domainErrors.ErrOfflineActionNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Offline action not found",
			})
		}
		log.Printf("Error updating offline action %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Offline action updated successfully: %s", id)
	return ctx.JSON(c.mapper.ToOfflineActionResponse(action))
}

// Delete
func (c *OfflineActionController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	err := c.deleteUseCase.Execute(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrOfflineActionNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Offline action not found",
			})
		}
		log.Printf("Error deleting offline action %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Offline action deleted successfully: %s", id)
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
