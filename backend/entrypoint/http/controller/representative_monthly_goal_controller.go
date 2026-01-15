package controller

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	representativemonthlygoal "github.com/seu-usuario/solis-backend/application/usecase/representativemonthlygoal"
	representativemonthlygoalerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/mapper"
	representativemonthlygoalrequest "github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

type RepresentativeMonthlyGoalController struct {
	createUseCase       *representativemonthlygoal.CreateRepresentativeMonthlyGoalUseCase
	getUseCase          *representativemonthlygoal.GetRepresentativeMonthlyGoalUseCase
	updateUseCase       *representativemonthlygoal.UpdateRepresentativeMonthlyGoalUseCase
	deleteUseCase       *representativemonthlygoal.DeleteRepresentativeMonthlyGoalUseCase
	listUseCase         *representativemonthlygoal.ListRepresentativeMonthlyGoalsUseCase
	getTableDataUseCase *representativemonthlygoal.GetRepresentativeGoalsTableDataUseCase
	mapper              *mapper.RepresentativeMonthlyGoalPayloadMapper
	gateway             gateway.RepresentativeMonthlyGoalGateway
}

func NewRepresentativeMonthlyGoalController(
	create *representativemonthlygoal.CreateRepresentativeMonthlyGoalUseCase,
	get *representativemonthlygoal.GetRepresentativeMonthlyGoalUseCase,
	update *representativemonthlygoal.UpdateRepresentativeMonthlyGoalUseCase,
	delete *representativemonthlygoal.DeleteRepresentativeMonthlyGoalUseCase,
	list *representativemonthlygoal.ListRepresentativeMonthlyGoalsUseCase,
	getTableData *representativemonthlygoal.GetRepresentativeGoalsTableDataUseCase,
	gateway gateway.RepresentativeMonthlyGoalGateway,
) *RepresentativeMonthlyGoalController {
	return &RepresentativeMonthlyGoalController{
		createUseCase:       create,
		getUseCase:          get,
		updateUseCase:       update,
		deleteUseCase:       delete,
		listUseCase:         list,
		getTableDataUseCase: getTableData,
		mapper:              mapper.NewRepresentativeMonthlyGoalPayloadMapper(),
		gateway:             gateway,
	}
}

// Create creates a new Representative Monthly Goal
func (c *RepresentativeMonthlyGoalController) Create(ctx *fiber.Ctx) error {
	var req representativemonthlygoalrequest.CreateRepresentativeMonthlyGoalRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Execute use case
	input := c.mapper.CreateRequestToInput(&req)
	output, err := c.createUseCase.Execute(ctx.Context(), *input)
	if err != nil {
		if errors.Is(err, representativemonthlygoalerrors.ErrRepresentativeAlreadyExists) {
			return ctx.Status(fiber.StatusConflict).JSON(response.ErrorResponse{
				Error: "Goal for this representative and month/year already exists",
			})
		}
		log.Printf("Error creating representative monthly goal: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Representative monthly goal created successfully: %s", output.ID.String())
	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.CreateOutputToResponse(output))
}

// GetByID finds a Representative Monthly Goal by ID
func (c *RepresentativeMonthlyGoalController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validate UUID format
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

	output, err := c.getUseCase.Execute(ctx.Context(), *input)
	if err != nil {
		if errors.Is(err, representativemonthlygoalerrors.ErrRepresentativeGoalNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Representative monthly goal not found",
			})
		}
		log.Printf("Error getting representative monthly goal %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.GetOutputToResponse(output))
}

// Update updates an existing Representative Monthly Goal
func (c *RepresentativeMonthlyGoalController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validate UUID format
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req representativemonthlygoalrequest.UpdateRepresentativeMonthlyGoalRequest
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

	output, err := c.updateUseCase.Execute(ctx.Context(), *input)
	if err != nil {
		if errors.Is(err, representativemonthlygoalerrors.ErrRepresentativeGoalNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Representative monthly goal not found",
			})
		}
		log.Printf("Error updating representative monthly goal %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Representative monthly goal updated successfully: %s", id)
	return ctx.JSON(c.mapper.UpdateOutputToResponse(output))
}

// Delete deletes a Representative Monthly Goal (soft delete)
func (c *RepresentativeMonthlyGoalController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validate UUID format
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

	err = c.deleteUseCase.Execute(ctx.Context(), *input)
	if err != nil {
		if errors.Is(err, representativemonthlygoalerrors.ErrRepresentativeGoalNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Representative monthly goal not found",
			})
		}
		log.Printf("Error deleting representative monthly goal %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Representative monthly goal deleted successfully: %s", id)
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// List lists Representative Monthly Goals based on criteria
func (c *RepresentativeMonthlyGoalController) List(ctx *fiber.Ctx) error {
	var req representativemonthlygoalrequest.ListRepresentativeMonthlyGoalsRequest
	if err := ctx.QueryParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid query parameters",
		})
	}

	// Execute use case
	input := c.mapper.ListRequestToInput(&req)
	output, err := c.listUseCase.Execute(ctx.Context(), *input)
	if err != nil {
		log.Printf("Error listing representative monthly goals: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Representative monthly goals listed successfully: count=%d", len(output.Data))
	return ctx.JSON(c.mapper.ListOutputToResponse(output))
}

// GetTableData returns table data for Representative Monthly Goals (transposed view)
func (c *RepresentativeMonthlyGoalController) GetTableData(ctx *fiber.Ctx) error {
	var req representativemonthlygoalrequest.GetRepresentativeGoalsTableDataRequest
	if err := ctx.QueryParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid query parameters",
		})
	}

	// Execute use case
	input := c.mapper.GetTableDataRequestToInput(&req)
	output, err := c.getTableDataUseCase.Execute(ctx.Context(), *input)
	if err != nil {
		log.Printf("Error getting representative monthly goal table data: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Representative monthly goal table data retrieved successfully")
	return ctx.JSON(c.mapper.GetTableDataOutputToResponse(output))
}
