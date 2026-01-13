package controller

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	usecase "github.com/seu-usuario/solis-backend/application/usecase/gifts"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/mapper"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

type GiftItemController struct {
	createUseCase *usecase.CreateGiftItemUseCase
	listUseCase   *usecase.ListGiftItemsUseCase
	getUseCase    *usecase.GetGiftItemUseCase
	updateUseCase *usecase.UpdateGiftItemUseCase
	deleteUseCase *usecase.DeleteGiftItemUseCase
	mapper        *mapper.GiftItemPayloadMapper
}

func NewGiftItemController(
	create *usecase.CreateGiftItemUseCase,
	list *usecase.ListGiftItemsUseCase,
	get *usecase.GetGiftItemUseCase,
	update *usecase.UpdateGiftItemUseCase,
	delete *usecase.DeleteGiftItemUseCase,
) *GiftItemController {
	return &GiftItemController{
		createUseCase: create,
		listUseCase:   list,
		getUseCase:    get,
		updateUseCase: update,
		deleteUseCase: delete,
		mapper:        mapper.NewGiftItemPayloadMapper(),
	}
}

// Create
func (c *GiftItemController) Create(ctx *fiber.Ctx) error {
	var req request.CreateGiftItemRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case
	input := c.mapper.CreateRequestToInput(&req)
	output, err := c.createUseCase.Execute(input)
	if err != nil {
		log.Printf("Error creating gift item: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Gift item created successfully: %s", output.ID)
	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.OutputToResponse(output))
}

// List
func (c *GiftItemController) List(ctx *fiber.Ctx) error {
	var query request.ListGiftItemsQuery
	if err := ctx.QueryParser(&query); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid query parameters",
		})
	}

	// Executar use case
	input := c.mapper.ListQueryToInput(&query)
	outputs, err := c.listUseCase.Execute(input)
	if err != nil {
		log.Printf("Error listing gift items: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	responses := c.mapper.ListOutputsToResponses(outputs)

	// Determinar página e limite para o meta
	page := 1
	limit := len(responses)
	if query.Page != nil {
		page = *query.Page
	}
	if query.Limit != nil {
		limit = *query.Limit
	}

	log.Printf("Gift items listed successfully: count=%d", len(responses))
	return ctx.JSON(response.GiftItemListData{
		Items: responses,
		Meta: response.MetaResponse{
			Total: int64(len(responses)),
			Page:  page,
			Limit: limit,
		},
	})
}

// GetByID
func (c *GiftItemController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	input, err := c.mapper.GetRequestToInput(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	output, err := c.getUseCase.Execute(input)
	if err != nil {
		if errors.Is(err, domainerrors.ErrGiftItemNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Gift item not found",
			})
		}
		log.Printf("Error getting gift item %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.GetOutputToResponse(output))
}

// Update
func (c *GiftItemController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req request.UpdateGiftItemRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	input, err := c.mapper.UpdateRequestToInput(id, &req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	output, err := c.updateUseCase.Execute(input)
	if err != nil {
		if errors.Is(err, domainerrors.ErrGiftItemNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Gift item not found",
			})
		}
		log.Printf("Error updating gift item %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Gift item updated successfully: %s", id)
	return ctx.JSON(c.mapper.UpdateOutputToResponse(output))
}

// Delete
func (c *GiftItemController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	input, err := c.mapper.DeleteRequestToInput(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	err = c.deleteUseCase.Execute(input)
	if err != nil {
		if errors.Is(err, domainerrors.ErrGiftItemNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Gift item not found",
			})
		}
		if errors.Is(err, domainerrors.ErrGiftItemHasTransactions) {
			return ctx.Status(fiber.StatusConflict).JSON(response.ErrorResponse{
				Error: "Gift item has associated transactions and cannot be deleted",
			})
		}
		log.Printf("Error deleting gift item %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Gift item deleted successfully: %s", id)
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
