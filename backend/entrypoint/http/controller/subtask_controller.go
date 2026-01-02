package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/seu-usuario/solis-backend/application/usecase/tasks"
	taskrequest "github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	taskresponse "github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// SubtaskController manipula requisições HTTP para Subtasks
type SubtaskController struct {
	createSubtaskUseCase *tasks.CreateSubtaskUseCase
	updateSubtaskUseCase *tasks.UpdateSubtaskUseCase
	deleteSubtaskUseCase *tasks.DeleteSubtaskUseCase
	getSubtaskUseCase    *tasks.GetSubtaskUseCase
	listSubtasksUseCase  *tasks.ListSubtasksUseCase
	mapper               *TaskMapper
}

// NewSubtaskController cria um novo SubtaskController
func NewSubtaskController(
	createSubtaskUseCase *tasks.CreateSubtaskUseCase,
	updateSubtaskUseCase *tasks.UpdateSubtaskUseCase,
	deleteSubtaskUseCase *tasks.DeleteSubtaskUseCase,
	getSubtaskUseCase *tasks.GetSubtaskUseCase,
	listSubtasksUseCase *tasks.ListSubtasksUseCase,
) *SubtaskController {
	return &SubtaskController{
		createSubtaskUseCase: createSubtaskUseCase,
		updateSubtaskUseCase: updateSubtaskUseCase,
		deleteSubtaskUseCase: deleteSubtaskUseCase,
		getSubtaskUseCase:    getSubtaskUseCase,
		listSubtasksUseCase:  listSubtasksUseCase,
		mapper:               NewTaskMapper(),
	}
}

// Create cria uma nova subtarefa
// @Summary Criar subtarefa
// @Description Cria uma nova subtarefa
// @Tags subtasks
// @Accept json
// @Produce json
// @Param request body taskrequest.CreateSubtaskRequest true "Dados da subtarefa"
// @Success 201 {object} taskresponse.SubtaskResponse
// @Failure 400 {object} taskresponse.ErrorResponse
// @Router /subtasks [post]
func (c *SubtaskController) Create(ctx *fiber.Ctx) error {
	var req taskrequest.CreateSubtaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	subtask, err := c.mapper.ToSubtask(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	err = c.createSubtaskUseCase.Execute(subtask)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.ToSubtaskResponse(subtask))
}

// Update atualiza uma subtarefa existente
// @Summary Atualizar subtarefa
// @Description Atualiza uma subtarefa existente
// @Tags subtasks
// @Accept json
// @Produce json
// @Param id path string true "ID da subtarefa"
// @Param request body taskrequest.UpdateSubtaskRequest true "Dados da subtarefa"
// @Success 200 {object} taskresponse.SubtaskResponse
// @Failure 400 {object} taskresponse.ErrorResponse
// @Failure 404 {object} taskresponse.ErrorResponse
// @Router /subtasks/{id} [put]
func (c *SubtaskController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var req taskrequest.UpdateSubtaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Convert assigneeID to *string
	var assigneeID *string
	if req.AssigneeID != nil {
		assigneeID = req.AssigneeID
	}

	// Convert Status to completed *bool
	var completed *bool
	if req.Status != nil {
		statusStr := string(*req.Status)
		completedVal := statusStr == "completed"
		completed = &completedVal
	}

	// Convert dueDate to *string
	var dueDate *string
	if req.DueDate != nil {
		dueDate = req.DueDate
	}

	updatedSubtask, err := c.updateSubtaskUseCase.Execute(
		id,
		*req.Title,
		completed,
		assigneeID,
		dueDate,
	)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToSubtaskResponse(updatedSubtask))
}

// Delete deleta uma subtarefa
// @Summary Deletar subtarefa
// @Description Deleta uma subtarefa existente
// @Tags subtasks
// @Produce json
// @Param id path string true "ID da subtarefa"
// @Success 204
// @Failure 404 {object} taskresponse.ErrorResponse
// @Router /subtasks/{id} [delete]
func (c *SubtaskController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.deleteSubtaskUseCase.Execute(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// Get retorna uma subtarefa por ID
// @Summary Obter subtarefa
// @Description Retorna uma subtarefa por ID
// @Tags subtasks
// @Produce json
// @Param id path string true "ID da subtarefa"
// @Success 200 {object} taskresponse.SubtaskResponse
// @Failure 404 {object} taskresponse.ErrorResponse
// @Router /subtasks/{id} [get]
func (c *SubtaskController) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	subtask, err := c.getSubtaskUseCase.Execute(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToSubtaskResponse(subtask))
}

// List retorna uma lista de subtarefas com paginação
// @Summary Listar subtarefas
// @Description Retorna uma lista de subtarefas com paginação
// @Tags subtasks
// @Produce json
// @Param page query int false "Número da página" default(1)
// @Param limit query int false "Itens por página" default(10)
// @Success 200 {object} taskresponse.SubtasksListResponse
// @Router /subtasks [get]
func (c *SubtaskController) List(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	subtasks, total, err := c.listSubtasksUseCase.Execute(page, limit)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToSubtasksListResponse(subtasks, total, page, limit))
}
