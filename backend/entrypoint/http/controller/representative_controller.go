package controller

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	representatives "github.com/seu-usuario/solis-backend/application/usecase/representatives"
	representativeerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/mapper"
	representativerequest "github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	representativeresponse "github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

type RepresentativeController struct {
	createUseCase         *representatives.CreateRepresentativeUseCase
	getUseCase            *representatives.GetRepresentativeUseCase
	updateUseCase         *representatives.UpdateRepresentativeUseCase
	deleteUseCase         *representatives.DeleteRepresentativeUseCase
	listUseCase           *representatives.ListRepresentativesUseCase
	getStatsUseCase       *representatives.GetRepresentativeStatsUseCase
	getProfileUseCase     *representatives.GetRepresentativeProfileUseCase
	getAllProfilesUseCase *representatives.GetAllRepresentativeProfilesUseCase
	mapper                *mapper.RepresentativePayloadMapper
	gateway               gateway.RepresentativeGateway
}

func NewRepresentativeController(
	create *representatives.CreateRepresentativeUseCase,
	get *representatives.GetRepresentativeUseCase,
	update *representatives.UpdateRepresentativeUseCase,
	delete *representatives.DeleteRepresentativeUseCase,
	list *representatives.ListRepresentativesUseCase,
	getStats *representatives.GetRepresentativeStatsUseCase,
	getProfile *representatives.GetRepresentativeProfileUseCase,
	getAllProfiles *representatives.GetAllRepresentativeProfilesUseCase,
	gateway gateway.RepresentativeGateway,
) *RepresentativeController {
	return &RepresentativeController{
		createUseCase:         create,
		getUseCase:            get,
		updateUseCase:         update,
		deleteUseCase:         delete,
		listUseCase:           list,
		getStatsUseCase:       getStats,
		getProfileUseCase:     getProfile,
		getAllProfilesUseCase: getAllProfiles,
		mapper:                mapper.NewRepresentativePayloadMapper(),
		gateway:               gateway,
	}
}

// Create creates a new Representative
func (c *RepresentativeController) Create(ctx *fiber.Ctx) error {
	var req representativerequest.CreateRepresentativeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(representativeresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Execute use case
	input := c.mapper.CreateRequestToInput(&req)
	output, err := c.createUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, representativeerrors.ErrRepresentativeAlreadyExists) {
			return ctx.Status(fiber.StatusConflict).JSON(representativeresponse.ErrorResponse{
				Error: "Representative with this code already exists",
			})
		}
		log.Printf("Error creating representative: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(representativeresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Representative created successfully: %s", output.UUID.String())
	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.CreateOutputToResponse(output))
}

// GetByID finds a Representative by ID
func (c *RepresentativeController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validate UUID format
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(representativeresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	input, err := c.mapper.GetRequestToInput(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(representativeresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	output, err := c.getUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, representativeerrors.ErrRepresentativeNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(representativeresponse.ErrorResponse{
				Error: "Representative not found",
			})
		}
		log.Printf("Error getting representative %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(representativeresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.GetOutputToResponse(output))
}

// Update updates an existing Representative
func (c *RepresentativeController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validate UUID format
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(representativeresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req representativerequest.UpdateRepresentativeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(representativeresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	input, err := c.mapper.UpdateRequestToInput(id, &req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(representativeresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	output, err := c.updateUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, representativeerrors.ErrRepresentativeNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(representativeresponse.ErrorResponse{
				Error: "Representative not found",
			})
		}
		log.Printf("Error updating representative %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(representativeresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Representative updated successfully: %s", id)
	return ctx.JSON(c.mapper.UpdateOutputToResponse(output))
}

// Delete deletes a Representative (soft delete)
func (c *RepresentativeController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validate UUID format
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(representativeresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	input, err := c.mapper.DeleteRequestToInput(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(representativeresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	err = c.deleteUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, representativeerrors.ErrRepresentativeNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(representativeresponse.ErrorResponse{
				Error: "Representative not found",
			})
		}
		log.Printf("Error deleting representative %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(representativeresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Representative deleted successfully: %s", id)
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// List lists Representatives based on criteria
func (c *RepresentativeController) List(ctx *fiber.Ctx) error {
	var req representativerequest.ListRepresentativesRequest
	if err := ctx.QueryParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(representativeresponse.ErrorResponse{
			Error: "Invalid query parameters",
		})
	}

	// Execute use case
	input := c.mapper.ListRequestToInput(&req)
	output, err := c.listUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error listing representatives: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(representativeresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Representatives listed successfully: count=%d", len(output.Data))
	return ctx.JSON(c.mapper.ListOutputToResponse(output))
}

// GetTableData returns table data for Representatives
func (c *RepresentativeController) GetTableData(ctx *fiber.Ctx) error {
	var req representativerequest.ListRepresentativesRequest
	if err := ctx.QueryParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(representativeresponse.ErrorResponse{
			Error: "Invalid query parameters",
		})
	}

	// Execute use case
	input := c.mapper.ListRequestToInput(&req)
	output, err := c.listUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error getting representative table data: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(representativeresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	// Convert to table data format
	tableData := representativeresponse.ListRepresentativesResponse{
		Data:       c.mapper.ListOutputToResponse(output).Data,
		Total:      output.Total,
		Page:       output.Page,
		PageSize:   output.PageSize,
		TotalPages: output.TotalPages,
	}

	log.Printf("Representative table data retrieved successfully: count=%d", len(tableData.Data))
	return ctx.JSON(tableData)
}

// GetStats returns statistics for a Representative
func (c *RepresentativeController) GetStats(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validate UUID format
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(representativeresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	input, err := c.mapper.GetStatsRequestToInput(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(representativeresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	output, err := c.getStatsUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, representativeerrors.ErrRepresentativeNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(representativeresponse.ErrorResponse{
				Error: "Representative not found",
			})
		}
		log.Printf("Error getting representative stats %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(representativeresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.StatsOutputToResponse(output))
}

// GetProfile returns a representative profile by name
func (c *RepresentativeController) GetProfile(ctx *fiber.Ctx) error {
	repName := ctx.Params("name")

	if repName == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(representativeresponse.ErrorResponse{
			Error: "Representative name is required",
		})
	}

	input := representatives.GetRepresentativeProfileInput{
		Name: repName,
	}

	output, err := c.getProfileUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if err == representativeerrors.ErrRepresentativeNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(representativeresponse.ErrorResponse{
				Error: "Representative not found",
			})
		}
		log.Printf("Error getting representative profile %s: %v", repName, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(representativeresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ProfileOutputToResponse(output))
}

// GetAllProfiles returns all representative profiles
func (c *RepresentativeController) GetAllProfiles(ctx *fiber.Ctx) error {
	output, err := c.getAllProfilesUseCase.Execute(ctx.Context())
	if err != nil {
		log.Printf("Error getting all representative profiles: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(representativeresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.AllProfilesOutputToResponse(output))
}
