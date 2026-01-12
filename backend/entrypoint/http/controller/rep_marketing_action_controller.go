package controller

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	usecase "solis/backend/application/usecase/cooperative/repmarketingaction"
	domainErrors "solis/backend/core/domain/errors"
	"solis/backend/entrypoint/http/payload/request"
	"solis/backend/entrypoint/http/payload/response"
)

type RepMarketingActionController struct {
	createUseCase *usecase.CreateRepMarketingActionUseCase
	listUseCase   *usecase.ListRepMarketingActionsUseCase
	getUseCase    *usecase.GetRepMarketingActionUseCase
	updateUseCase *usecase.UpdateRepMarketingActionUseCase
	deleteUseCase *usecase.DeleteRepMarketingActionUseCase
	mapper        *response.RepMarketingActionPayloadMapper
}

func NewRepMarketingActionController(
	create *usecase.CreateRepMarketingActionUseCase,
	list *usecase.ListRepMarketingActionsUseCase,
	get *usecase.GetRepMarketingActionUseCase,
	update *usecase.UpdateRepMarketingActionUseCase,
	delete *usecase.DeleteRepMarketingActionUseCase,
) *RepMarketingActionController {
	return &RepMarketingActionController{
		createUseCase: create,
		listUseCase:   list,
		getUseCase:    get,
		updateUseCase: update,
		deleteUseCase: delete,
		mapper:        response.NewRepMarketingActionPayloadMapper(),
	}
}

// Create
func (c *RepMarketingActionController) Create(ctx *fiber.Ctx) error {
	var req request.CreateRepMarketingActionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case (validação está no use case)
	input := usecase.CreateRepMarketingActionInput{
		RepName:     req.RepName,
		Date:        req.Date,
		Description: req.Description,
	}

	action, err := c.createUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error creating rep marketing action: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Rep marketing action created successfully: %s", action.UUID())
	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.ToRepMarketingActionResponse(action))
}

// List
func (c *RepMarketingActionController) List(ctx *fiber.Ctx) error {
	var query request.ListRepMarketingActionsQuery
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
			"date":       true,
			"created_at": true,
		}
		if !validSortBy[*query.SortBy] {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Error: "Invalid sort_by value. Valid values: date, created_at",
			})
		}
	}

	// Executar use case
	input := usecase.ListRepMarketingActionsInput{
		RepName:   query.RepName,
		Month:     query.Month,
		Page:      query.Page,
		Limit:     query.Limit,
		SortBy:    query.SortBy,
		SortOrder: query.SortOrder,
	}

	actions, total, err := c.listUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error listing rep marketing actions: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Rep marketing actions listed successfully: page=%d, limit=%d, total=%d", query.Page, query.Limit, total)
	return ctx.JSON(c.mapper.ToRepMarketingActionsListResponse(actions, total, query.Page, query.Limit))
}

// GetByID
func (c *RepMarketingActionController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	action, err := c.getUseCase.Execute(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrRepMarketingActionNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Rep marketing action not found",
			})
		}
		log.Printf("Error getting rep marketing action %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToRepMarketingActionResponse(action))
}

// Update
func (c *RepMarketingActionController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req request.UpdateRepMarketingActionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case
	input := usecase.UpdateRepMarketingActionInput{
		ID:          id,
		RepName:     req.RepName,
		Date:        req.Date,
		Description: req.Description,
	}

	action, err := c.updateUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, domainErrors.ErrRepMarketingActionNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Rep marketing action not found",
			})
		}
		log.Printf("Error updating rep marketing action %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Rep marketing action updated successfully: %s", id)
	return ctx.JSON(c.mapper.ToRepMarketingActionResponse(action))
}

// Delete
func (c *RepMarketingActionController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	err := c.deleteUseCase.Execute(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrRepMarketingActionNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Rep marketing action not found",
			})
		}
		log.Printf("Error deleting rep marketing action %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Rep marketing action deleted successfully: %s", id)
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
