package controller

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	usecase "github.com/seu-usuario/solis-backend/application/usecase/cooperative/showroomitem"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

type ShowroomItemController struct {
	createUseCase *usecase.CreateShowroomItemUseCase
	listUseCase   *usecase.ListShowroomItemsUseCase
	getUseCase    *usecase.GetShowroomItemUseCase
	updateUseCase *usecase.UpdateShowroomItemUseCase
	deleteUseCase *usecase.DeleteShowroomItemUseCase
	mapper        *response.ShowroomItemPayloadMapper
}

func NewShowroomItemController(
	create *usecase.CreateShowroomItemUseCase,
	list *usecase.ListShowroomItemsUseCase,
	get *usecase.GetShowroomItemUseCase,
	update *usecase.UpdateShowroomItemUseCase,
	delete *usecase.DeleteShowroomItemUseCase,
) *ShowroomItemController {
	return &ShowroomItemController{
		createUseCase: create,
		listUseCase:   list,
		getUseCase:    get,
		updateUseCase: update,
		deleteUseCase: delete,
		mapper:        response.NewShowroomItemPayloadMapper(),
	}
}

// Create
func (c *ShowroomItemController) Create(ctx *fiber.Ctx) error {
	var req request.CreateShowroomItemRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case (validação está no use case)
	input := usecase.CreateShowroomItemInput{
		PDV:                req.PDV,
		City:               req.City,
		Contact:            req.Contact,
		RepresentativeUUID: req.RepresentativeUUID,
		DeliveryForecast:   req.DeliveryForecast,
		WorkshopDate:       req.WorkshopDate,
	}

	item, err := c.createUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error creating showroom item: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Showroom item created successfully: %s", item.ID())
	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.ToShowroomItemResponse(item))
}

// List
func (c *ShowroomItemController) List(ctx *fiber.Ctx) error {
	var query request.ListShowroomItemsQuery
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
			"created_at":        true,
			"delivery_forecast": true,
		}
		if !validSortBy[*query.SortBy] {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Error: "Invalid sort_by value. Valid values: created_at, delivery_forecast",
			})
		}
	}

	// Executar use case
	input := usecase.ListShowroomItemsInput{
		RepresentativeUUID: query.RepresentativeUUID,
		Delivered:          query.Delivered,
		City:               query.City,
		Page:               query.Page,
		Limit:              query.Limit,
		SortBy:             query.SortBy,
		SortOrder:          query.SortOrder,
	}

	items, total, err := c.listUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error listing showroom items: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Showroom items listed successfully: page=%d, limit=%d, total=%d", query.Page, query.Limit, total)
	return ctx.JSON(c.mapper.ToShowroomItemsListResponse(items, total, query.Page, query.Limit))
}

// GetByID
func (c *ShowroomItemController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	item, err := c.getUseCase.Execute(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrShowroomItemNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Showroom item not found",
			})
		}
		log.Printf("Error getting showroom item %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToShowroomItemResponse(item))
}

// Update
func (c *ShowroomItemController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req request.UpdateShowroomItemRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case
	input := usecase.UpdateShowroomItemInput{
		ID:                 id,
		City:               req.City,
		Contact:            req.Contact,
		RepresentativeUUID: req.RepresentativeUUID,
		DeliveryForecast:   req.DeliveryForecast,
		WorkshopDate:       req.WorkshopDate,
		Delivered:          req.Delivered,
		PDV:                req.PDV,
	}

	item, err := c.updateUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, domainErrors.ErrShowroomItemNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Showroom item not found",
			})
		}
		log.Printf("Error updating showroom item %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Showroom item updated successfully: %s", id)
	return ctx.JSON(c.mapper.ToShowroomItemResponse(item))
}

// Delete
func (c *ShowroomItemController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	err := c.deleteUseCase.Execute(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrShowroomItemNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Showroom item not found",
			})
		}
		log.Printf("Error deleting showroom item %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Showroom item deleted successfully: %s", id)
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
