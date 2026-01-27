package controller

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/application/usecase/kpis"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/helper"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// KpiController handles HTTP requests for KPI operations
type KpiController struct {
	createKpiUseCase         *kpis.CreateKpiUseCase
	getKpiUseCase            *kpis.GetKpiUseCase
	listKpisUseCase          *kpis.ListKpisUseCase
	getKpisBySlugsUseCase    *kpis.GetKpisBySlugsUseCase
	updateKpiUseCase         *kpis.UpdateKpiUseCase
	deleteKpiUseCase         *kpis.DeleteKpiUseCase
	updateMonthlyDataUseCase *kpis.UpdateMonthlyDataUseCase
	mapper                   *KpiMapper
}

// NewKpiController creates a new KpiController instance
func NewKpiController(
	createKpiUseCase *kpis.CreateKpiUseCase,
	getKpiUseCase *kpis.GetKpiUseCase,
	listKpisUseCase *kpis.ListKpisUseCase,
	getKpisBySlugsUseCase *kpis.GetKpisBySlugsUseCase,
	updateKpiUseCase *kpis.UpdateKpiUseCase,
	deleteKpiUseCase *kpis.DeleteKpiUseCase,
	updateMonthlyDataUseCase *kpis.UpdateMonthlyDataUseCase,
) *KpiController {
	return &KpiController{
		createKpiUseCase:         createKpiUseCase,
		getKpiUseCase:            getKpiUseCase,
		listKpisUseCase:          listKpisUseCase,
		getKpisBySlugsUseCase:    getKpisBySlugsUseCase,
		updateKpiUseCase:         updateKpiUseCase,
		deleteKpiUseCase:         deleteKpiUseCase,
		updateMonthlyDataUseCase: updateMonthlyDataUseCase,
		mapper:                   NewKpiMapper(),
	}
}

// Create handles KPI creation
// @Summary Create a new KPI
// @Description Create a new KPI category with the provided data
// @Tags kpis
// @Accept json
// @Produce json
// @Param kpiRequest body request.CreateKpiRequest true "KPI data"
// @Success 201 {object} response.KpiResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /kpis [post]
func (c *KpiController) Create(ctx *fiber.Ctx) error {
	// Parse request body
	var req request.CreateKpiRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Execute use case
	input := kpis.CreateKpiInput{
		Title: req.Title,
		Color: req.Color,
		Unit:  req.Unit,
	}

	kpi, err := c.createKpiUseCase.Execute(context.Background(), input)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	// Convert to response - now includes monthly data
	kpiResponse := c.mapper.ToKpiResponseWithMonthlyData(kpi, kpi.MonthlyDatas())

	return ctx.Status(fiber.StatusCreated).JSON(kpiResponse)
}

// GetByID handles KPI retrieval by ID
// @Summary Get KPI by ID
// @Description Retrieve a KPI category by its ID
// @Tags kpis
// @Accept json
// @Produce json
// @Param id path string true "KPI ID"
// @Success 200 {object} response.KpiResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /kpis/{id} [get]
func (c *KpiController) GetByID(ctx *fiber.Ctx) error {
	// Extract ID from params
	id := ctx.Params("id")

	kpi, err := c.getKpiUseCase.Execute(context.Background(), id)
	if err != nil {
		if err == errors.ErrKpiNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "KPI not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "Internal server error",
		})
	}

	// Convert to response - now includes monthly data
	kpiResponse := c.mapper.ToKpiResponseWithMonthlyData(kpi, kpi.MonthlyDatas())

	return ctx.Status(fiber.StatusOK).JSON(kpiResponse)
}

// List handles KPI listing with pagination and filters
// @Summary List KPIs
// @Description Retrieve a paginated list of KPI categories with optional month/year filters
// @Tags kpis
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Param month query string false "Month filter (JAN, FEV, MAR, etc.)"
// @Param year query int false "Year filter"
// @Success 200 {object} response.KpiListResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /kpis [get]
func (c *KpiController) List(ctx *fiber.Ctx) error {
	// Parse query params using BaseQueryParams
	var queryParams request.BaseQueryParams
	if err := ctx.QueryParser(&queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid query parameters",
		})
	}

	// Validate query params
	if err := queryParams.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	// Create pagination value object
	pagination := valueobject.NewPagination(queryParams.GetPage(), queryParams.GetLimit())

	// Execute use case with filters (month/year not used in usecase, filtering done in response)
	output, err := c.listKpisUseCase.Execute(context.Background(), pagination)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "Internal server error",
		})
	}

	// Convert entities to responses with monthly data
	kpiResponses := make([]response.KpiResponse, len(output.Kpis))
	for i, kpi := range output.Kpis {
		// Filter monthly data based on month/year if provided
		filteredMonthlyData := helper.FilterMonthlyData(kpi.MonthlyDatas(), queryParams.GetMonth(), queryParams.GetYear())
		kpiResponses[i] = *c.mapper.ToKpiResponseWithMonthlyData(kpi, filteredMonthlyData)
	}

	// Calculate pagination info
	totalPages := int(output.Total) / queryParams.GetLimit()
	if int(output.Total)%queryParams.GetLimit() > 0 {
		totalPages++
	}

	// Create list response
	listResponse := response.KpiListResponse{
		Data: kpiResponses,
		Pagination: response.PaginationResponse{
			Page:       queryParams.GetPage(),
			PageSize:   queryParams.GetLimit(),
			Total:      int64(output.Total),
			TotalPages: totalPages,
		},
	}

	return ctx.Status(fiber.StatusOK).JSON(listResponse)
}

// GetBySlugs handles KPI retrieval by a list of slugs
// @Summary Get KPIs by slugs
// @Description Retrieve KPI categories by a list of slugs with optional month/year filters.
// Note: Pagination is always set to Page=1, PageSize=len(results), Total=len(results), TotalPages=1
// since this endpoint returns a filtered result set from a specific list of slugs.
// @Tags kpis
// @Accept json
// @Produce json
// @Param slugs body request.GetKpisBySlugsRequest true "List of slugs"
// @Param month query string false "Month filter (JAN, FEV, MAR, etc.)"
// @Param year query int false "Year filter"
// @Success 200 {object} response.KpiListResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /kpis/by-slugs [post]
func (c *KpiController) GetBySlugs(ctx *fiber.Ctx) error {
	// Parse request body
	var req request.GetKpisBySlugsRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Parse query params for month/year filters
	var queryParams request.BaseQueryParams
	if err := ctx.QueryParser(&queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid query parameters",
		})
	}

	// Validate query params
	if err := queryParams.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	// Execute use case
	input := kpis.GetKpisBySlugsInput{
		Slugs: req.Slugs,
	}

	output, err := c.getKpisBySlugsUseCase.Execute(context.Background(), input)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "Internal server error",
		})
	}

	// Convert entities to responses with monthly data, applying filters
	kpiResponses := make([]response.KpiResponse, len(output.Kpis))
	for i, kpi := range output.Kpis {
		// Filter monthly data based on month/year if provided
		filteredMonthlyData := helper.FilterMonthlyData(kpi.MonthlyDatas(), queryParams.GetMonth(), queryParams.GetYear())

		// Check if it's full year (month = ---)
		if queryParams.GetMonth() != nil && *queryParams.GetMonth() == "---" {
			// Calculate annual sum
			var totalRealized float64 = 0
			var totalMeta float64 = 0

			for _, data := range filteredMonthlyData {
				if data.Realized() != nil {
					totalRealized += *data.Realized()
				}
				if data.Meta() != nil {
					totalMeta += *data.Meta()
				}
			}

			// Create consolidated item with month = ---
			consolidatedMonthlyData, err := entity.ReconstructMonthlyData(
				uuid.New(),
				kpi.ID(),
				*queryParams.GetYear(), // Use filtered year (dereference pointer)
				"---",                  // Consolidated month
				&totalRealized,         // Annual sum
				&totalMeta,             // Sum of all metas
				nil,                    // Empty breakdown
				time.Now(),
				time.Now(),
			)
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
					Error: "Failed to create consolidated monthly data",
				})
			}

			consolidatedData := []*entity.MonthlyData{consolidatedMonthlyData}
			kpiResponses[i] = *c.mapper.ToKpiResponseWithMonthlyData(kpi, consolidatedData)
		} else {
			// Use current logic: use filtered monthly data
			kpiResponses[i] = *c.mapper.ToKpiResponseWithMonthlyData(kpi, filteredMonthlyData)
		}
	}

	// Create list response
	listResponse := response.KpiListResponse{
		Data: kpiResponses,
		Pagination: response.PaginationResponse{
			Page:       1,
			PageSize:   len(kpiResponses),
			Total:      int64(len(kpiResponses)),
			TotalPages: 1,
		},
	}

	return ctx.Status(fiber.StatusOK).JSON(listResponse)
}

// Update handles KPI update
// @Summary Update KPI
// @Description Update an existing KPI category
// @Tags kpis
// @Accept json
// @Produce json
// @Param id path string true "KPI ID"
// @Param kpiRequest body request.UpdateKpiRequest true "KPI update data"
// @Success 200 {object} response.KpiResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /kpis/{id} [put]
func (c *KpiController) Update(ctx *fiber.Ctx) error {
	// Extract ID from params
	id := ctx.Params("id")

	// Parse request body
	var req request.UpdateKpiRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Execute use case
	input := kpis.UpdateKpiInput{
		Title: req.Title,
		Color: req.Color,
		Unit:  req.Unit,
	}

	kpi, err := c.updateKpiUseCase.Execute(context.Background(), id, input)
	if err != nil {
		if err == errors.ErrKpiNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "KPI not found",
			})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	// Convert to response - now includes monthly data
	kpiResponse := c.mapper.ToKpiResponseWithMonthlyData(kpi, kpi.MonthlyDatas())

	return ctx.Status(fiber.StatusOK).JSON(kpiResponse)
}

// Delete handles KPI deletion
// @Summary Delete KPI
// @Description Delete an existing KPI category
// @Tags kpis
// @Accept json
// @Produce json
// @Param id path string true "KPI ID"
// @Success 204
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /kpis/{id} [delete]
func (c *KpiController) Delete(ctx *fiber.Ctx) error {
	// Extract ID from params
	id := ctx.Params("id")

	err := c.deleteKpiUseCase.Execute(context.Background(), id)
	if err != nil {
		if err == errors.ErrKpiNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "KPI not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "Internal server error",
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// UpdateMonthlyData handles monthly data update
// @Summary Update monthly data
// @Description Update monthly data for a specific KPI
// @Tags kpis
// @Accept json
// @Produce json
// @Param kpiId path string true "KPI ID"
// @Param monthlyDataRequest body request.UpdateMonthlyDataRequest true "Monthly data update"
// @Success 200 {object} response.MonthlyDataResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /kpis/{kpiId}/monthly-data [put]
func (c *KpiController) UpdateMonthlyData(ctx *fiber.Ctx) error {
	// Extract ID from params
	kpiId := ctx.Params("kpiId")

	// Parse request body
	var req request.UpdateMonthlyDataRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Extract user ID from JWT token
	userID := ""
	if ctx.Locals("userID") != nil {
		userID = ctx.Locals("userID").(string)
	}

	// Execute use case
	input := kpis.UpdateMonthlyDataInput{
		KpiID:     kpiId,
		Year:      req.Year,
		Month:     req.Month,
		Realized:  req.Realized,
		Meta:      req.Meta,
		Breakdown: req.Breakdown,
		UserID:    userID,
		Context:   req.Context,
	}

	monthlyData, err := c.updateMonthlyDataUseCase.Execute(context.Background(), input)
	if err != nil {
		if err == errors.ErrMonthDataNotFound || err == errors.ErrKpiNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Monthly data or KPI not found",
			})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: err.Error(),
		})
	}

	// Convert to response
	monthlyDataResponse := c.mapper.ToMonthlyDataResponse(monthlyData)

	return ctx.Status(fiber.StatusOK).JSON(monthlyDataResponse)
}
