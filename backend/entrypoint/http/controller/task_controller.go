package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/application/usecase/tasks"
	taskrequest "github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	taskresponse "github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// TaskController manipula requisições HTTP para Tasks
type TaskController struct {
	createTaskUseCase   *tasks.CreateTaskUseCase
	updateTaskUseCase   *tasks.UpdateTaskUseCase
	deleteTaskUseCase   *tasks.DeleteTaskUseCase
	getTaskUseCase      *tasks.GetTaskUseCase
	listTasksUseCase    *tasks.ListTasksUseCase
	reorderTasksUseCase *tasks.ReorderTasksUseCase
	mapper              *TaskMapper
}

// NewTaskController cria um novo TaskController
func NewTaskController(
	createTaskUseCase *tasks.CreateTaskUseCase,
	updateTaskUseCase *tasks.UpdateTaskUseCase,
	deleteTaskUseCase *tasks.DeleteTaskUseCase,
	getTaskUseCase *tasks.GetTaskUseCase,
	listTasksUseCase *tasks.ListTasksUseCase,
	reorderTasksUseCase *tasks.ReorderTasksUseCase,
) *TaskController {
	return &TaskController{
		createTaskUseCase:   createTaskUseCase,
		updateTaskUseCase:   updateTaskUseCase,
		deleteTaskUseCase:   deleteTaskUseCase,
		getTaskUseCase:      getTaskUseCase,
		listTasksUseCase:    listTasksUseCase,
		reorderTasksUseCase: reorderTasksUseCase,
		mapper:              NewTaskMapper(),
	}
}

// Create cria uma nova tarefa
// @Summary Criar tarefa
// @Description Cria uma nova tarefa
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body taskrequest.CreateTaskRequest true "Dados da tarefa"
// @Success 201 {object} taskresponse.TaskResponse
// @Failure 400 {object} taskresponse.ErrorResponse
// @Router /tasks [post]
func (c *TaskController) Create(ctx *fiber.Ctx) error {
	var req taskrequest.CreateTaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	task, err := c.mapper.ToTask(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	err = c.createTaskUseCase.Execute(task)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.ToTaskResponse(task))
}

// Update atualiza uma tarefa existente
// @Summary Atualizar tarefa
// @Description Atualiza uma tarefa existente
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "ID da tarefa"
// @Param request body taskrequest.UpdateTaskRequest true "Dados da tarefa"
// @Success 200 {object} taskresponse.TaskResponse
// @Failure 400 {object} taskresponse.ErrorResponse
// @Failure 404 {object} taskresponse.ErrorResponse
// @Router /tasks/{id} [put]
func (c *TaskController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var req taskrequest.UpdateTaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Convert flows to []string
	flows := make([]string, len(req.Flows))
	for i, flow := range req.Flows {
		flows[i] = string(flow)
	}

	// Convert assigneeID to *string
	var assigneeID *string
	if req.AssigneeID != nil {
		assigneeID = req.AssigneeID
	}

	// Convert dueDate to *string
	var dueDate *string
	if req.DueDate != nil {
		dueDate = req.DueDate
	}

	// Convert startDate to *string
	var startDate *string
	if req.StartDate != nil {
		startDate = req.StartDate
	}

	// Convert to *string
	var category, priority, status *string
	if req.Category != nil {
		cat := string(*req.Category)
		category = &cat
	}
	if req.Priority != nil {
		pri := string(*req.Priority)
		priority = &pri
	}
	if req.Status != nil {
		stat := string(*req.Status)
		status = &stat
	}

	// Convert description to *string
	var description *string
	if req.Description != "" {
		description = &req.Description
	}

	updatedTask, err := c.updateTaskUseCase.Execute(
		id,
		req.Title,
		description,
		category,
		priority,
		status,
		assigneeID,
		req.Archived,
		flows,
		startDate,
		dueDate,
	)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToTaskResponse(updatedTask))
}

// Delete deleta uma tarefa
// @Summary Deletar tarefa
// @Description Deleta uma tarefa existente
// @Tags tasks
// @Produce json
// @Param id path string true "ID da tarefa"
// @Success 204
// @Failure 404 {object} taskresponse.ErrorResponse
// @Router /tasks/{id} [delete]
func (c *TaskController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.deleteTaskUseCase.Execute(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// Get retorna uma tarefa por ID
// @Summary Obter tarefa
// @Description Retorna uma tarefa por ID
// @Tags tasks
// @Produce json
// @Param id path string true "ID da tarefa"
// @Success 200 {object} taskresponse.TaskResponse
// @Failure 404 {object} taskresponse.ErrorResponse
// @Router /tasks/{id} [get]
func (c *TaskController) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	task, err := c.getTaskUseCase.Execute(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToTaskResponse(task))
}

// List retorna uma lista de tarefas com paginação
// @Summary Listar tarefas
// @Description Retorna uma lista de tarefas com paginação
// @Tags tasks
// @Produce json
// @Param page query int false "Número da página" default(1)
// @Param limit query int false "Itens por página" default(10)
// @Success 200 {object} taskresponse.TasksListResponse
// @Router /tasks [get]
func (c *TaskController) List(ctx *fiber.Ctx) error {
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	tasks, total, err := c.listTasksUseCase.Execute(
		nil, // statuses
		nil, // categories
		nil, // priorities
		nil, // assigneeID
		nil, // flow
		nil, // archived
		nil, // flows
		nil, // dateFrom
		nil, // dateTo
		page,
		limit,
		"created_at", // sortBy
		"desc",       // sortOrder
	)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToTasksListResponse(tasks, total, page, limit))
}

// Reorder reordena múltiplas tarefas
// @Summary Reordenar tarefas
// @Description Reordena múltiplas tarefas na ordem especificada
// @Tags tasks
// @Accept json
// @Produce json
// @Param request body taskrequest.ReorderTasksRequest true "IDs das tarefas na ordem desejada"
// @Success 204
// @Failure 400 {object} taskresponse.ErrorResponse
// @Router /tasks/reorder [post]
func (c *TaskController) Reorder(ctx *fiber.Ctx) error {
	var req taskrequest.ReorderTasksRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	// Parse task IDs
	taskIDs := make([]uuid.UUID, len(req.TaskIDs))
	for i, idStr := range req.TaskIDs {
		id, err := uuid.Parse(idStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
				Error: "Invalid task ID",
			})
		}
		taskIDs[i] = id
	}

	err := c.reorderTasksUseCase.Execute(taskIDs)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
