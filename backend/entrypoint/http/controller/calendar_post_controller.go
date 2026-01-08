package controller

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	calendarusecase "github.com/seu-usuario/solis-backend/application/usecase/calendar"
	"github.com/seu-usuario/solis-backend/entrypoint/http/middleware"
	calendarrequest "github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	calendarresponse "github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// CalendarPostController manipula requisições HTTP para CalendarPosts
type CalendarPostController struct {
	createCalendarPostUseCase            *calendarusecase.CreateCalendarPostUseCase
	getCalendarPostUseCase               *calendarusecase.GetCalendarPostUseCase
	updateCalendarPostUseCase            *calendarusecase.UpdateCalendarPostUseCase
	updateCalendarPostStatusUseCase      *calendarusecase.UpdateCalendarPostStatusUseCase
	confirmCalendarPostPublishingUseCase *calendarusecase.ConfirmCalendarPostPublishingUseCase
	deleteCalendarPostUseCase            *calendarusecase.DeleteCalendarPostUseCase
	listCalendarPostsUseCase             *calendarusecase.ListCalendarPostsUseCase
	mapper                               *CalendarPostMapper
}

// NewCalendarPostController cria um novo CalendarPostController
func NewCalendarPostController(
	createCalendarPostUseCase *calendarusecase.CreateCalendarPostUseCase,
	getCalendarPostUseCase *calendarusecase.GetCalendarPostUseCase,
	updateCalendarPostUseCase *calendarusecase.UpdateCalendarPostUseCase,
	updateCalendarPostStatusUseCase *calendarusecase.UpdateCalendarPostStatusUseCase,
	confirmCalendarPostPublishingUseCase *calendarusecase.ConfirmCalendarPostPublishingUseCase,
	deleteCalendarPostUseCase *calendarusecase.DeleteCalendarPostUseCase,
	listCalendarPostsUseCase *calendarusecase.ListCalendarPostsUseCase,
) *CalendarPostController {
	return &CalendarPostController{
		createCalendarPostUseCase:            createCalendarPostUseCase,
		getCalendarPostUseCase:               getCalendarPostUseCase,
		updateCalendarPostUseCase:            updateCalendarPostUseCase,
		updateCalendarPostStatusUseCase:      updateCalendarPostStatusUseCase,
		confirmCalendarPostPublishingUseCase: confirmCalendarPostPublishingUseCase,
		deleteCalendarPostUseCase:            deleteCalendarPostUseCase,
		listCalendarPostsUseCase:             listCalendarPostsUseCase,
		mapper:                               NewCalendarPostMapper(),
	}
}

// Create cria um novo post no calendário
func (c *CalendarPostController) Create(ctx *fiber.Ctx) error {
	var req calendarrequest.CreateCalendarPostRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(calendarresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Parse assigneeID
	assigneeID, err := uuid.Parse(req.AssigneeID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(calendarresponse.ErrorResponse{
			Error: "Invalid assignee ID",
		})
	}

	// Create input for use case
	input := calendarusecase.CreateCalendarPostInput{
		Title:       req.Title,
		Description: req.Description,
		Date:        req.Date,
		Time:        req.Time,
		Caption:     req.Caption,
		Category:    req.Category,
		Type:        req.Type,
		AssigneeID:  &assigneeID,
		Platforms:   req.Platforms,
		ImageURL:    req.ImageURL,
	}

	post, err := c.createCalendarPostUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error creating calendar post: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(calendarresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Calendar post created successfully: %s", post.ID())
	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.ToCalendarPostResponse(post))
}

// Get retorna um post por ID
func (c *CalendarPostController) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(calendarresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	post, err := c.getCalendarPostUseCase.Execute(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(calendarresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToCalendarPostResponse(post))
}

// Update atualiza um post existente
func (c *CalendarPostController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(calendarresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req calendarrequest.UpdateCalendarPostRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(calendarresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Create input for use case
	input := calendarusecase.UpdateCalendarPostInput{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Caption:     req.Caption,
		ImageURL:    req.ImageURL,
	}

	post, err := c.updateCalendarPostUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error updating calendar post %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(calendarresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Calendar post updated successfully: %s", id)
	return ctx.JSON(c.mapper.ToCalendarPostResponse(post))
}

// UpdateStatus atualiza o status de um post
func (c *CalendarPostController) UpdateStatus(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(calendarresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req calendarrequest.UpdateCalendarPostStatusRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(calendarresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Extract user from auth context
	userEmail, err := middleware.GetUserEmail(ctx)
	if err != nil {
		log.Printf("Error getting user from context: %v", err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(calendarresponse.ErrorResponse{
			Error: "User not authenticated",
		})
	}

	// Create input for use case
	input := calendarusecase.UpdateStatusInput{
		PostID:    id,
		NewStatus: req.Status,
		User:      userEmail,
		Comment:   nil,
	}

	post, err := c.updateCalendarPostStatusUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error updating status for calendar post %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(calendarresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Calendar post status updated successfully: %s to %s", id, req.Status)
	return ctx.JSON(c.mapper.ToCalendarPostResponse(post))
}

// ConfirmPublishing confirma a publicação de um post
func (c *CalendarPostController) ConfirmPublishing(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(calendarresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	var req calendarrequest.ConfirmCalendarPostPublishingRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(calendarresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Extract user from auth context
	userEmail, err := middleware.GetUserEmail(ctx)
	if err != nil {
		log.Printf("Error getting user from context: %v", err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(calendarresponse.ErrorResponse{
			Error: "User not authenticated",
		})
	}

	// Create input for use case
	input := calendarusecase.ConfirmPublishingInput{
		PostID:    id,
		Platforms: req.PublishedPlatforms,
		User:      userEmail,
	}

	post, err := c.confirmCalendarPostPublishingUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error confirming publishing for calendar post %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(calendarresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Calendar post publishing confirmed successfully: %s", id)
	return ctx.JSON(c.mapper.ToCalendarPostResponse(post))
}

// Delete deleta um post
func (c *CalendarPostController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	// Validar formato do UUID
	if _, err := uuid.Parse(id); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(calendarresponse.ErrorResponse{
			Error: "Invalid ID format",
		})
	}

	err := c.deleteCalendarPostUseCase.Execute(ctx.Context(), id)
	if err != nil {
		log.Printf("Error deleting calendar post %s: %v", id, err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(calendarresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Calendar post deleted successfully: %s", id)
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// List retorna uma lista de posts com paginação
func (c *CalendarPostController) List(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Parse query parameters
	var category, postType, status, startDate, endDate, platform, sortBy, sortOrder *string
	if cat := ctx.Query("category"); cat != "" {
		category = &cat
	}
	if typ := ctx.Query("type"); typ != "" {
		postType = &typ
	}
	if stat := ctx.Query("status"); stat != "" {
		status = &stat
	}
	if sd := ctx.Query("start_date"); sd != "" {
		startDate = &sd
	}
	if ed := ctx.Query("end_date"); ed != "" {
		endDate = &ed
	}
	if plat := ctx.Query("platform"); plat != "" {
		platform = &plat
	}
	if sb := ctx.Query("sort_by"); sb != "" {
		sortBy = &sb
	}
	if so := ctx.Query("sort_order"); so != "" {
		sortOrder = &so
	}

	// Extract userID from auth context (optional for filtering)
	var assigneeID *uuid.UUID
	userID, err := middleware.GetUserID(ctx)
	if err == nil {
		assigneeID = &userID
	}

	// Create input for use case
	input := calendarusecase.ListCalendarPostsInput{
		Category:   category,
		Type:       postType,
		Status:     status,
		AssigneeID: assigneeID,
		StartDate:  startDate,
		EndDate:    endDate,
		Platform:   platform,
		Page:       page,
		Limit:      limit,
		SortBy:     sortBy,
		SortOrder:  sortOrder,
	}

	posts, total, err := c.listCalendarPostsUseCase.Execute(ctx.Context(), input)
	if err != nil {
		log.Printf("Error listing calendar posts: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(calendarresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("Calendar posts listed successfully: page=%d, limit=%d, total=%d", page, limit, total)
	return ctx.JSON(c.mapper.ToCalendarPostsListResponse(posts, total, page, limit))
}
