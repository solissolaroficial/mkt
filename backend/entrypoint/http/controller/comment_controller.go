package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/seu-usuario/solis-backend/application/usecase/tasks"
	taskrequest "github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	taskresponse "github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// CommentController manipula requisições HTTP para Comments
type CommentController struct {
	createCommentUseCase *tasks.CreateCommentUseCase
	updateCommentUseCase *tasks.UpdateCommentUseCase
	deleteCommentUseCase *tasks.DeleteCommentUseCase
	getCommentUseCase    *tasks.GetCommentUseCase
	listCommentsUseCase  *tasks.ListCommentsUseCase
	mapper               *TaskMapper
}

// NewCommentController cria um novo CommentController
func NewCommentController(
	createCommentUseCase *tasks.CreateCommentUseCase,
	updateCommentUseCase *tasks.UpdateCommentUseCase,
	deleteCommentUseCase *tasks.DeleteCommentUseCase,
	getCommentUseCase *tasks.GetCommentUseCase,
	listCommentsUseCase *tasks.ListCommentsUseCase,
) *CommentController {
	return &CommentController{
		createCommentUseCase: createCommentUseCase,
		updateCommentUseCase: updateCommentUseCase,
		deleteCommentUseCase: deleteCommentUseCase,
		getCommentUseCase:    getCommentUseCase,
		listCommentsUseCase:  listCommentsUseCase,
		mapper:               NewTaskMapper(),
	}
}

// Create cria um novo comentário
// @Summary Criar comentário
// @Description Cria um novo comentário
// @Tags comments
// @Accept json
// @Produce json
// @Param request body taskrequest.CreateCommentRequest true "Dados do comentário"
// @Success 201 {object} taskresponse.CommentResponse
// @Failure 400 {object} taskresponse.ErrorResponse
// @Router /comments [post]
func (c *CommentController) Create(ctx *fiber.Ctx) error {
	var req taskrequest.CreateCommentRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	comment, err := c.mapper.ToComment(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	err = c.createCommentUseCase.Execute(ctx.Context(), comment)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.ToCommentResponse(comment))
}

// Update atualiza um comentário existente
// @Summary Atualizar comentário
// @Description Atualiza um comentário existente
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "ID do comentário"
// @Param request body taskrequest.UpdateCommentRequest true "Dados do comentário"
// @Success 200 {object} taskresponse.CommentResponse
// @Failure 400 {object} taskresponse.ErrorResponse
// @Failure 404 {object} taskresponse.ErrorResponse
// @Router /comments/{id} [put]
func (c *CommentController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var req taskrequest.UpdateCommentRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	updateErr := c.updateCommentUseCase.Execute(id, req.Content)
	if updateErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: updateErr.Error(),
		})
	}

	// Get updated comment
	comment, getErr := c.getCommentUseCase.Execute(id)
	if getErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: getErr.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToCommentResponse(comment))
}

// Delete deleta um comentário
// @Summary Deletar comentário
// @Description Deleta um comentário existente
// @Tags comments
// @Produce json
// @Param id path string true "ID do comentário"
// @Success 204
// @Failure 404 {object} taskresponse.ErrorResponse
// @Router /comments/{id} [delete]
func (c *CommentController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.deleteCommentUseCase.Execute(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// Get retorna um comentário por ID
// @Summary Obter comentário
// @Description Retorna um comentário por ID
// @Tags comments
// @Produce json
// @Param id path string true "ID do comentário"
// @Success 200 {object} taskresponse.CommentResponse
// @Failure 404 {object} taskresponse.ErrorResponse
// @Router /comments/{id} [get]
func (c *CommentController) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	comment, getErr := c.getCommentUseCase.Execute(id)
	if getErr != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(taskresponse.ErrorResponse{
			Error: getErr.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToCommentResponse(comment))
}

// List retorna uma lista de comentários de uma tarefa
// @Summary Listar comentários
// @Description Retorna uma lista de comentários de uma tarefa
// @Tags comments
// @Produce json
// @Param task_id query string true "ID da tarefa"
// @Param page query int false "Número da página (padrão: 1)"
// @Param limit query int false "Itens por página (padrão: 10)"
// @Success 200 {object} taskresponse.CommentsListResponse
// @Router /comments [get]
func (c *CommentController) List(ctx *fiber.Ctx) error {
	taskID := ctx.Query("task_id")
	if taskID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
			Error: "task_id is required",
		})
	}

	// Extrair parâmetros de paginação da query com valores padrão
	page := 1
	limit := 10

	if pageStr := ctx.Query("page"); pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr := ctx.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	comments, total, listErr := c.listCommentsUseCase.ExecuteWithPagination(taskID, page, limit)
	if listErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: listErr.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToCommentsListResponse(comments, total, page, limit))
}
