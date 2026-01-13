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

type GiftTransactionController struct {
	createUseCase *usecase.CreateGiftTransactionUseCase
	listUseCase   *usecase.ListGiftTransactionsUseCase
	getUseCase    *usecase.GetGiftTransactionUseCase
	updateUseCase *usecase.UpdateGiftTransactionUseCase
	deleteUseCase *usecase.DeleteGiftTransactionUseCase
	mapper        *mapper.GiftTransactionPayloadMapper
}

func NewGiftTransactionController(
	create *usecase.CreateGiftTransactionUseCase,
	list *usecase.ListGiftTransactionsUseCase,
	get *usecase.GetGiftTransactionUseCase,
	update *usecase.UpdateGiftTransactionUseCase,
	delete *usecase.DeleteGiftTransactionUseCase,
) *GiftTransactionController {
	return &GiftTransactionController{
		createUseCase: create,
		listUseCase:   list,
		getUseCase:    get,
		updateUseCase: update,
		deleteUseCase: delete,
		mapper:        mapper.NewGiftTransactionPayloadMapper(),
	}
}

// Create
func (c *GiftTransactionController) Create(ctx *fiber.Ctx) error {
	var req request.CreateGiftTransactionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case
	input := c.mapper.CreateRequestToInput(&req)
	output, err := c.createUseCase.Execute(input)
	if err != nil {
		log.Printf("Error creating gift transaction: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Gift transaction created successfully: %s", output.ID)
	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.OutputToResponse(output))
}

// List
func (c *GiftTransactionController) List(ctx *fiber.Ctx) error {
	var query request.ListGiftTransactionsQuery
	if err := ctx.QueryParser(&query); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid query parameters",
		})
	}

	// Executar use case
	input := c.mapper.ListQueryToInput(&query)
	outputs, err := c.listUseCase.Execute(input)
	if err != nil {
		log.Printf("Error listing gift transactions: %v", err)
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

	log.Printf("Gift transactions listed successfully: count=%d", len(responses))
	return ctx.JSON(response.GiftTransactionListData{
		Transactions: responses,
		Meta: response.MetaResponse{
			Total: int64(len(responses)),
			Page:  page,
			Limit: limit,
		},
	})
}

// GetByID
func (c *GiftTransactionController) GetByID(ctx *fiber.Ctx) error {
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
		if errors.Is(err, domainerrors.ErrGiftTransactionNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Gift transaction not found",
			})
		}
		log.Printf("Error getting gift transaction %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.GetOutputToResponse(output))
}

// Update
func (c *GiftTransactionController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req request.UpdateGiftTransactionRequest
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
		if errors.Is(err, domainerrors.ErrGiftTransactionNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Gift transaction not found",
			})
		}
		if errors.Is(err, domainerrors.ErrInvalidFieldForTransactionType) {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Error: "Campos incompatíveis com o tipo de transação (price para 'out' ou representative para 'in')",
			})
		}
		log.Printf("Error updating gift transaction %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Gift transaction updated successfully: %s", id)
	return ctx.JSON(c.mapper.UpdateOutputToResponse(output))
}

// Delete
func (c *GiftTransactionController) Delete(ctx *fiber.Ctx) error {
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
		if errors.Is(err, domainerrors.ErrGiftTransactionNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Gift transaction not found",
			})
		}
		log.Printf("Error deleting gift transaction %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Gift transaction deleted successfully: %s", id)
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
