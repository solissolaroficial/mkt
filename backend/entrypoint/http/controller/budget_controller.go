package controller

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	budget "github.com/seu-usuario/solis-backend/application/usecase/budget"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/mapper"
	budgetrequest "github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	budgetresponse "github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

type BudgetController struct {
	createUseCase      *budget.CreateBudgetItemUseCase
	listUseCase        *budget.ListBudgetItemsUseCase
	getUseCase         *budget.GetBudgetItemUseCase
	updateUseCase      *budget.UpdateBudgetItemUseCase
	deleteUseCase      *budget.DeleteBudgetItemUseCase
	batchCreateUseCase *budget.BatchCreateBudgetItemsUseCase
	getSummaryUseCase  *budget.GetBudgetSummaryUseCase
	mapper             *mapper.BudgetPayloadMapper
	gateway            gateway.BudgetGateway
}

func NewBudgetController(
	create *budget.CreateBudgetItemUseCase,
	list *budget.ListBudgetItemsUseCase,
	get *budget.GetBudgetItemUseCase,
	update *budget.UpdateBudgetItemUseCase,
	delete *budget.DeleteBudgetItemUseCase,
	batchCreate *budget.BatchCreateBudgetItemsUseCase,
	getSummary *budget.GetBudgetSummaryUseCase,
	gateway gateway.BudgetGateway,
) *BudgetController {
	return &BudgetController{
		createUseCase:      create,
		listUseCase:        list,
		getUseCase:         get,
		updateUseCase:      update,
		deleteUseCase:      delete,
		batchCreateUseCase: batchCreate,
		getSummaryUseCase:  getSummary,
		mapper:             mapper.NewBudgetPayloadMapper(),
		gateway:            gateway,
	}
}

// Create cria um novo BudgetItem
func (c *BudgetController) Create(ctx *fiber.Ctx) error {
	var req budgetrequest.CreateBudgetItemRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(budgetresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case
	input := c.mapper.CreateRequestToInput(&req)
	output, err := c.createUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, domainerrors.ErrBudgetAlreadyExists) {
			return ctx.Status(fiber.StatusConflict).JSON(budgetresponse.ErrorResponse{
				Error: "Budget item with same code already exists",
			})
		}
		log.Printf("Error creating budget item: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(budgetresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Budget item created successfully: %s", output.ID())
	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.OutputToResponse(output))
}

// List lista BudgetItems baseados em critérios
func (c *BudgetController) List(ctx *fiber.Ctx) error {
	var query budgetrequest.ListBudgetItemsQuery
	if err := ctx.QueryParser(&query); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(budgetresponse.ErrorResponse{
			Error: "Invalid query parameters",
		})
	}

	// Executar use case
	criteria := query.ToCriteria()
	outputs, err := c.listUseCase.Execute(ctx.Context(), criteria)
	if err != nil {
		log.Printf("Error listing budget items: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(budgetresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	responses := c.mapper.ListOutputsToResponses(outputs)

	log.Printf("Budget items listed successfully: count=%d", len(responses))
	return ctx.JSON(responses)
}

// GetByID busca um BudgetItem pelo ID
func (c *BudgetController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(budgetresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	output, err := c.getUseCase.Execute(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, domainerrors.ErrBudgetNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(budgetresponse.ErrorResponse{
				Error: "Budget item not found",
			})
		}
		log.Printf("Error getting budget item %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(budgetresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.OutputToResponse(output))
}

// Update atualiza um BudgetItem existente
func (c *BudgetController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(budgetresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req budgetrequest.UpdateBudgetItemRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(budgetresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	_, input := c.mapper.UpdateRequestToInput(id, &req)
	output, err := c.updateUseCase.Execute(ctx.Context(), id, input)
	if err != nil {
		if errors.Is(err, domainerrors.ErrBudgetNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(budgetresponse.ErrorResponse{
				Error: "Budget item not found",
			})
		}
		log.Printf("Error updating budget item %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(budgetresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Budget item updated successfully: %s", id)
	return ctx.JSON(c.mapper.OutputToResponse(output))
}

// Delete remove um BudgetItem (soft delete)
func (c *BudgetController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(budgetresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	err := c.deleteUseCase.Execute(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, domainerrors.ErrBudgetNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(budgetresponse.ErrorResponse{
				Error: "Budget item not found",
			})
		}
		log.Printf("Error deleting budget item %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(budgetresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Budget item deleted successfully: %s", id)
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// BatchCreate cria múltiplos BudgetItems em lote
func (c *BudgetController) BatchCreate(ctx *fiber.Ctx) error {
	var req budgetrequest.BatchCreateBudgetItemsRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(budgetresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case
	input := c.mapper.BatchCreateRequestToInput(&req)
	outputs, err := c.batchCreateUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error batch creating budget items: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(budgetresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	responses := c.mapper.ListOutputsToResponses(outputs)

	log.Printf("Budget items batch created successfully: count=%d", len(responses))
	return ctx.Status(fiber.StatusCreated).JSON(responses)
}

// GetSummary retorna resumo agregado de orçamento
func (c *BudgetController) GetSummary(ctx *fiber.Ctx) error {
	var query budgetrequest.ListBudgetItemsQuery
	if err := ctx.QueryParser(&query); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(budgetresponse.ErrorResponse{
			Error: "Invalid query parameters",
		})
	}

	// Executar use case
	criteria := query.ToCriteria()
	summaries, err := c.getSummaryUseCase.Execute(ctx.Context(), criteria)
	if err != nil {
		log.Printf("Error getting budget summary: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(budgetresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	responses := c.mapper.SummaryOutputsToResponses(summaries)

	log.Printf("Budget summary retrieved successfully: count=%d", len(responses))
	return ctx.JSON(responses)
}

// GetDistinctYears retorna os anos disponíveis
func (c *BudgetController) GetDistinctYears(ctx *fiber.Ctx) error {
	years, err := c.gateway.GetDistinctYears(ctx.Context())
	if err != nil {
		log.Printf("Error getting distinct years: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(budgetresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Distinct years retrieved successfully: count=%d", len(years))
	return ctx.JSON(struct {
		Years []int `json:"years"`
	}{
		Years: years,
	})
}
