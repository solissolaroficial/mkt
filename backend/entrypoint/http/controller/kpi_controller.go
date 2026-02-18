package controller

import (
	"context"
	"slices"
	"sort"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/application/usecase/kpis"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/helper"
	"github.com/seu-usuario/solis-backend/entrypoint/http/middleware"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// KPIs that use the last monthly value instead of sum for annual calculation
var kpisLastValue = []string{"autoridade_na_internet_da"}

// KpiController handles HTTP requests for KPI operations
type KpiController struct {
	createKpiUseCase         *kpis.CreateKpiUseCase
	getKpiUseCase            *kpis.GetKpiUseCase
	listKpisUseCase          *kpis.ListKpisUseCase
	getKpisBySlugsUseCase    *kpis.GetKpisBySlugsUseCase
	deleteKpiUseCase         *kpis.DeleteKpiUseCase
	updateMonthlyDataUseCase *kpis.UpdateMonthlyDataUseCase
	deleteMonthlyDataUseCase *kpis.DeleteMonthlyData
	mapper                   *KpiMapper
}

// NewKpiController creates a new KpiController instance
func NewKpiController(
	createKpiUseCase *kpis.CreateKpiUseCase,
	getKpiUseCase *kpis.GetKpiUseCase,
	listKpisUseCase *kpis.ListKpisUseCase,
	getKpisBySlugsUseCase *kpis.GetKpisBySlugsUseCase,
	deleteKpiUseCase *kpis.DeleteKpiUseCase,
	updateMonthlyDataUseCase *kpis.UpdateMonthlyDataUseCase,
	deleteMonthlyDataUseCase *kpis.DeleteMonthlyData,
) *KpiController {
	return &KpiController{
		createKpiUseCase:         createKpiUseCase,
		getKpiUseCase:            getKpiUseCase,
		listKpisUseCase:          listKpisUseCase,
		getKpisBySlugsUseCase:    getKpisBySlugsUseCase,
		deleteKpiUseCase:         deleteKpiUseCase,
		updateMonthlyDataUseCase: updateMonthlyDataUseCase,
		deleteMonthlyDataUseCase: deleteMonthlyDataUseCase,
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

	// Extract user ID from context (set by AuthMiddleware)
	userID, err := middleware.GetUserID(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "User not authenticated",
		})
	}

	// Execute use case
	input := kpis.CreateKpiInput{
		Title:     req.Title,
		Color:     req.Color,
		Unit:      req.Unit,
		CreatedBy: &userID,
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
			// Check if this KPI uses last value instead of sum
			usesLastValue := slices.Contains(kpisLastValue, kpi.Slug())

			var totalRealized float64 = 0
			var totalMeta float64 = 0

			if usesLastValue && len(filteredMonthlyData) > 0 {
				// Se usa último valor, garantir que dados estão ordenados por mês
				monthOrder := map[string]int{"JAN": 1, "FEV": 2, "MAR": 3, "ABR": 4, "MAI": 5, "JUN": 6,
					"JUL": 7, "AGO": 8, "SET": 9, "OUT": 10, "NOV": 11, "DEZ": 12}
				sort.Slice(filteredMonthlyData, func(i, j int) bool {
					return monthOrder[filteredMonthlyData[i].Month()] < monthOrder[filteredMonthlyData[j].Month()]
				})

				// Use the last monthly value for KPIs that require it
				lastData := filteredMonthlyData[len(filteredMonthlyData)-1]
				if lastData.Realized() != nil {
					totalRealized = *lastData.Realized()
				}
				if lastData.Meta() != nil {
					totalMeta = *lastData.Meta()
				}
			} else {
				// Calculate annual sum normally
				for _, data := range filteredMonthlyData {
					if data.Realized() != nil {
						totalRealized += *data.Realized()
					}
					if data.Meta() != nil {
						totalMeta += *data.Meta()
					}
				}
			}

			// Create consolidated item with month = ---
			consolidatedMonthlyData, err := entity.ReconstructMonthlyData(
				uuid.New(),
				kpi.ID(),
				*queryParams.GetYear(), // Use filtered year (dereference pointer)
				"---",                  // Consolidated month
				&totalRealized,         // Annual value (sum or last value)
				&totalMeta,             // Sum of all metas
				[]byte{},               // Empty breakdown as []byte
				time.Now(),
				time.Now(),
				nil, // deletedAt = nil (not deleted)
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

// Delete handles KPI deletion
// @Summary Delete KPI
// @Description Delete an existing KPI category
// @Tags kpis
// @Accept json
// @Produce json
// @Param id path string true "KPI ID"
// @Success 204
// @Failure 404 {object} response.ErrorResponse
// @Failure 403 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /kpis/{id} [delete]
func (c *KpiController) Delete(ctx *fiber.Ctx) error {
	// Extract ID from params
	id := ctx.Params("id")

	// Extract user ID from JWT token (safe type assertion)
	userID := ""
	if userIDStr, ok := ctx.Locals("userID").(string); ok {
		userID = userIDStr
	}

	// Check if user is admin
	isAdmin := false
	if role, ok := ctx.Locals("userRole").(string); ok && role == "admin" {
		isAdmin = true
	}

	err := c.deleteKpiUseCase.Execute(context.Background(), id, userID, isAdmin)
	if err != nil {
		if err == errors.ErrKpiNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "KPI not found",
			})
		}
		return ctx.Status(fiber.StatusForbidden).JSON(response.ErrorResponse{
			Error: err.Error(),
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

	// Extract user ID from JWT token (safe type assertion)
	userID := ""
	if userIDStr, ok := ctx.Locals("userID").(string); ok {
		userID = userIDStr
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

// DeleteMonthlyData handles monthly data deletion
// @Summary Delete monthly data
// @Description Delete a monthly data record by its ID (admin only)
// @Tags kpis
// @Accept json
// @Produce json
// @Param kpiId path string true "KPI ID"
// @Param monthlyDataId path string true "Monthly data ID"
// @Success 204
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /kpis/{kpiId}/monthly-data/{monthlyDataId} [delete]
func (c *KpiController) DeleteMonthlyData(ctx *fiber.Ctx) error {
	// Extract IDs from params
	monthlyDataId := ctx.Params("monthlyDataId")

	// Execute use case
	err := c.deleteMonthlyDataUseCase.Execute(context.Background(), monthlyDataId)
	if err != nil {
		if err == errors.ErrMonthDataNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "Monthly data not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "Internal server error",
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// AddDailyEntry handles adding a daily entry to monthly data
// @Summary Add daily entry
// @Description Add a daily entry to monthly data for a specific KPI
// @Tags kpis
// @Accept json
// @Produce json
// @Param kpiId path string true "KPI ID"
// @Param dailyEntryRequest body request.AddDailyEntryRequest true "Daily entry data"
// @Success 200 {object} response.MonthlyDataResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /kpis/{kpiId}/daily-entry [post]
func (c *KpiController) AddDailyEntry(ctx *fiber.Ctx) error {
	// Extract ID from params
	kpiId := ctx.Params("kpiId")

	// Parse request body
	var req request.AddDailyEntryRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Extract user ID from JWT token (safe type assertion)
	userID := ""
	if userIDStr, ok := ctx.Locals("userID").(string); ok {
		userID = userIDStr
	}
	// If userID is empty, return unauthorized error
	if userID == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "User not authenticated",
		})
	}

	// Execute use case
	input := kpis.AddDailyEntryInput{
		KpiID:   kpiId,
		Year:    req.Year,
		Month:   req.Month,
		Date:    req.Date,
		Value:   req.Value,
		Context: req.Context,
		User:    userID,
	}

	monthlyData, err := c.updateMonthlyDataUseCase.AddDailyEntry(context.Background(), input)
	if err != nil {
		if err == errors.ErrKpiNotFound || err == errors.ErrMonthDataNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "KPI or monthly data not found",
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

// UpdateDailyEntry handles updating a daily entry in monthly data
// @Summary Update daily entry
// @Description Update an existing daily entry in monthly data for a specific KPI
// @Tags kpis
// @Accept json
// @Produce json
// @Param kpiId path string true "KPI ID"
// @Param dailyEntryRequest body request.UpdateDailyEntryRequest true "Daily entry data"
// @Success 200 {object} response.MonthlyDataResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /kpis/{kpiId}/daily-entry [put]
func (c *KpiController) UpdateDailyEntry(ctx *fiber.Ctx) error {
	// Extract ID from params
	kpiId := ctx.Params("kpiId")

	// Parse request body
	var req request.UpdateDailyEntryRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Extract user ID from JWT token (safe type assertion)
	userID := ""
	if userIDStr, ok := ctx.Locals("userID").(string); ok {
		userID = userIDStr
	}
	// If userID is empty, return unauthorized error
	if userID == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "User not authenticated",
		})
	}

	// Execute use case
	input := kpis.UpdateDailyEntryInput{
		KpiID:   kpiId,
		Year:    req.Year,
		Month:   req.Month,
		Date:    req.Date,
		Value:   req.Value,
		Context: req.Context,
		User:    userID,
	}

	monthlyData, err := c.updateMonthlyDataUseCase.UpdateDailyEntry(context.Background(), input)
	if err != nil {
		if err == errors.ErrKpiNotFound || err == errors.ErrMonthDataNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "KPI or monthly data not found",
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

// DeleteDailyEntry handles deleting a daily entry from monthly data
// @Summary Delete daily entry
// @Description Delete a daily entry from monthly data for a specific KPI
// @Tags kpis
// @Accept json
// @Produce json
// @Param kpiId path string true "KPI ID"
// @Param dailyEntryRequest body request.DeleteDailyEntryRequest true "Daily entry data"
// @Success 200 {object} response.MonthlyDataResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /kpis/{kpiId}/daily-entry [delete]
func (c *KpiController) DeleteDailyEntry(ctx *fiber.Ctx) error {
	// Extract ID from params
	kpiId := ctx.Params("kpiId")

	// Parse request body
	var req request.DeleteDailyEntryRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Extract user ID from JWT token (safe type assertion)
	userID := ""
	if userIDStr, ok := ctx.Locals("userID").(string); ok {
		userID = userIDStr
	}
	// If userID is empty, return unauthorized error
	if userID == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Error: "User not authenticated",
		})
	}

	// Execute use case
	// req.Year is already an int, no conversion needed
	yearInt := req.Year

	input := kpis.DeleteDailyEntryInput{
		KpiID:   kpiId,
		Year:    yearInt,
		Month:   req.Month,
		Date:    req.Date,
		User:    userID,
		Context: req.Context,
	}

	monthlyData, err := c.updateMonthlyDataUseCase.DeleteDailyEntry(context.Background(), input)
	if err != nil {
		if err == errors.ErrKpiNotFound || err == errors.ErrMonthDataNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "KPI or monthly data not found",
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

// GetDailyEntries handles getting all daily entries for a specific KPI, year and month
// @Summary Get daily entries
// @Description Get all daily entries for a specific KPI, year and month
// @Tags kpis
// @Accept json
// @Produce json
// @Param kpiId path string true "KPI ID"
// @Param month query string true "Month (JAN, FEV, etc.)"
// @Param year query int true "Year"
// @Success 200 {array} entity.DailyEntry
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /kpis/{kpiId}/daily-entries [get]
func (c *KpiController) GetDailyEntries(ctx *fiber.Ctx) error {
	// Extract ID from params
	kpiId := ctx.Params("kpiId")

	// Parse query params
	month := ctx.Query("month")
	yearStr := ctx.Query("year")

	if month == "" || yearStr == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "month and year are required query parameters",
		})
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Error: "invalid year parameter",
		})
	}

	// Execute use case
	input := kpis.GetDailyEntriesInput{
		KpiID: kpiId,
		Year:  year,
		Month: month,
	}

	dailyEntries, err := c.updateMonthlyDataUseCase.GetDailyEntries(context.Background(), input)
	if err != nil {
		if err == errors.ErrKpiNotFound || err == errors.ErrMonthDataNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
				Error: "KPI or monthly data not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Error: "Internal server error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(dailyEntries)
}
