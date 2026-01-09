package controller

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	usecase "github.com/seu-usuario/solis-backend/application/usecase/social"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

type SocialController struct {
	createUseCase *usecase.CreateSocialBenchmarkingUseCase
	listUseCase   *usecase.ListSocialBenchmarkingsUseCase
	getUseCase    *usecase.GetSocialBenchmarkingUseCase
	updateUseCase *usecase.UpdateSocialBenchmarkingUseCase
	deleteUseCase *usecase.DeleteSocialBenchmarkingUseCase
	mapper        *response.SocialPayloadMapper
}

func NewSocialController(
	create *usecase.CreateSocialBenchmarkingUseCase,
	list *usecase.ListSocialBenchmarkingsUseCase,
	get *usecase.GetSocialBenchmarkingUseCase,
	update *usecase.UpdateSocialBenchmarkingUseCase,
	delete *usecase.DeleteSocialBenchmarkingUseCase,
) *SocialController {
	return &SocialController{
		createUseCase: create,
		listUseCase:   list,
		getUseCase:    get,
		updateUseCase: update,
		deleteUseCase: delete,
		mapper:        response.NewSocialPayloadMapper(),
	}
}

// Create
func (c *SocialController) Create(ctx *fiber.Ctx) error {
	var req request.CreateSocialBenchmarkingRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case (validação está no use case)
	input := usecase.CreateSocialBenchmarkingInput{
		BrandName:   req.BrandName,
		AvgLikes:    req.AvgLikes,
		AvgComments: req.AvgComments,
		Followers:   req.Followers,
	}

	benchmarking, err := c.createUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error creating social benchmarking: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Social benchmarking created successfully: %s", benchmarking.ID())
	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.ToSocialBenchmarkingResponse(benchmarking))
}

// List
func (c *SocialController) List(ctx *fiber.Ctx) error {
	var query request.ListSocialBenchmarkingsQuery
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
			"engagement_rate": true,
			"avg_likes":       true,
			"avg_comments":    true,
			"created_at":      true,
		}
		if !validSortBy[*query.SortBy] {
			return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Error: "Invalid sort_by value. Valid values: engagement_rate, avg_likes, avg_comments, created_at",
			})
		}
	}

	// Executar use case
	input := usecase.ListSocialBenchmarkingsInput{
		BrandName: query.BrandName,
		Active:    query.Active,
		StartDate: query.StartDate,
		EndDate:   query.EndDate,
		Page:      query.Page,
		Limit:     query.Limit,
		SortBy:    query.SortBy,
		SortOrder: query.SortOrder,
	}

	benchmarkings, total, err := c.listUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error listing social benchmarkings: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Social benchmarkings listed successfully: page=%d, limit=%d, total=%d", query.Page, query.Limit, total)
	return ctx.JSON(c.mapper.ToSocialBenchmarkingsListResponse(benchmarkings, total, query.Page, query.Limit))
}

// GetByID
func (c *SocialController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	benchmarking, err := c.getUseCase.Execute(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrSocialBenchmarkingNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Social benchmarking not found",
			})
		}
		log.Printf("Error getting social benchmarking %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToSocialBenchmarkingResponse(benchmarking))
}

// Update
func (c *SocialController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req request.UpdateSocialBenchmarkingRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Executar use case
	input := usecase.UpdateSocialBenchmarkingInput{
		ID:          id,
		BrandName:   req.BrandName,
		AvgLikes:    req.AvgLikes,
		AvgComments: req.AvgComments,
		Followers:   req.Followers,
	}

	benchmarking, err := c.updateUseCase.Execute(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, domainErrors.ErrSocialBenchmarkingNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Social benchmarking not found",
			})
		}
		log.Printf("Error updating social benchmarking %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Social benchmarking updated successfully: %s", id)
	return ctx.JSON(c.mapper.ToSocialBenchmarkingResponse(benchmarking))
}

// Delete
func (c *SocialController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	err := c.deleteUseCase.Execute(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, domainErrors.ErrSocialBenchmarkingNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Social benchmarking not found",
			})
		}
		log.Printf("Error deleting social benchmarking %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Social benchmarking deleted successfully: %s", id)
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
