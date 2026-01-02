package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/seu-usuario/solis-backend/application/usecase/tasks"
	taskrequest "github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	taskresponse "github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// NotificationController manipula requisições HTTP para Notifications
type NotificationController struct {
	createNotificationUseCase          *tasks.CreateNotificationUseCase
	updateNotificationUseCase          *tasks.UpdateNotificationUseCase
	deleteNotificationUseCase          *tasks.DeleteNotificationUseCase
	getNotificationUseCase             *tasks.GetNotificationUseCase
	listNotificationsUseCase           *tasks.ListNotificationsUseCase
	markAsReadNotificationUseCase      *tasks.MarkAsReadNotificationUseCase
	markAllAsReadNotificationsUseCase  *tasks.MarkAllAsReadNotificationsUseCase
	deleteNotificationsByTaskIDUseCase *tasks.DeleteNotificationsByTaskIDUseCase
	mapper                             *TaskMapper
}

// NewNotificationController cria um novo NotificationController
func NewNotificationController(
	createNotificationUseCase *tasks.CreateNotificationUseCase,
	updateNotificationUseCase *tasks.UpdateNotificationUseCase,
	deleteNotificationUseCase *tasks.DeleteNotificationUseCase,
	getNotificationUseCase *tasks.GetNotificationUseCase,
	listNotificationsUseCase *tasks.ListNotificationsUseCase,
	markAsReadNotificationUseCase *tasks.MarkAsReadNotificationUseCase,
	markAllAsReadNotificationsUseCase *tasks.MarkAllAsReadNotificationsUseCase,
	deleteNotificationsByTaskIDUseCase *tasks.DeleteNotificationsByTaskIDUseCase,
) *NotificationController {
	return &NotificationController{
		createNotificationUseCase:          createNotificationUseCase,
		updateNotificationUseCase:          updateNotificationUseCase,
		deleteNotificationUseCase:          deleteNotificationUseCase,
		getNotificationUseCase:             getNotificationUseCase,
		listNotificationsUseCase:           listNotificationsUseCase,
		markAsReadNotificationUseCase:      markAsReadNotificationUseCase,
		markAllAsReadNotificationsUseCase:  markAllAsReadNotificationsUseCase,
		deleteNotificationsByTaskIDUseCase: deleteNotificationsByTaskIDUseCase,
		mapper:                             NewTaskMapper(),
	}
}

// Create cria uma nova notificação
// @Summary Criar notificação
// @Description Cria uma nova notificação
// @Tags notifications
// @Accept json
// @Produce json
// @Param request body taskrequest.CreateNotificationRequest true "Dados da notificação"
// @Success 201 {object} taskresponse.NotificationResponse
// @Failure 400 {object} taskresponse.ErrorResponse
// @Router /notifications [post]
func (c *NotificationController) Create(ctx *fiber.Ctx) error {
	var req taskrequest.CreateNotificationRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	notification, err := c.createNotificationUseCase.Execute(
		req.UserID,
		req.Title,
		req.Message,
		req.NotificationType,
		req.TaskID,
	)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(c.mapper.ToNotificationResponse(notification))
}

// Update atualiza uma notificação existente
// @Summary Atualizar notificação
// @Description Atualiza uma notificação existente
// @Tags notifications
// @Accept json
// @Produce json
// @Param id path string true "ID da notificação"
// @Param request body taskrequest.UpdateNotificationRequest true "Dados da notificação"
// @Success 200 {object} taskresponse.NotificationResponse
// @Failure 400 {object} taskresponse.ErrorResponse
// @Failure 404 {object} taskresponse.ErrorResponse
// @Router /notifications/{id} [put]
func (c *NotificationController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var req taskrequest.UpdateNotificationRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
			Error: "Invalid request body",
		})
	}

	err := c.updateNotificationUseCase.Execute(id, req.Title, req.Message)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	// Get the updated notification
	notification, err := c.getNotificationUseCase.Execute(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToNotificationResponse(notification))
}

// Delete deleta uma notificação
// @Summary Deletar notificação
// @Description Deleta uma notificação existente
// @Tags notifications
// @Produce json
// @Param id path string true "ID da notificação"
// @Success 204
// @Failure 404 {object} taskresponse.ErrorResponse
// @Router /notifications/{id} [delete]
func (c *NotificationController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.deleteNotificationUseCase.Execute(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusNoContent).Send(nil)
}

// Get retorna uma notificação por ID
// @Summary Obter notificação
// @Description Retorna uma notificação por ID
// @Tags notifications
// @Produce json
// @Param id path string true "ID da notificação"
// @Success 200 {object} taskresponse.NotificationResponse
// @Failure 404 {object} taskresponse.ErrorResponse
// @Router /notifications/{id} [get]
func (c *NotificationController) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	notification, err := c.getNotificationUseCase.Execute(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToNotificationResponse(notification))
}

// List retorna uma lista de notificações com paginação
// @Summary Listar notificações
// @Description Retorna uma lista de notificações com paginação
// @Tags notifications
// @Produce json
// @Param user_id query string true "ID do usuário"
// @Param page query int false "Número da página" default(1)
// @Param limit query int false "Itens por página" default(10)
// @Success 200 {object} taskresponse.NotificationsListResponse
// @Router /notifications [get]
func (c *NotificationController) List(ctx *fiber.Ctx) error {
	userID := ctx.Query("user_id")
	if userID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
			Error: "user_id is required",
		})
	}

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	notifications, total, err := c.listNotificationsUseCase.Execute(userID, page, limit)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToNotificationsListResponse(notifications, total, page, limit))
}

// MarkAsRead marca uma notificação como lida
// @Summary Marcar notificação como lida
// @Description Marca uma notificação como lida
// @Tags notifications
// @Produce json
// @Param id path string true "ID da notificação"
// @Success 200 {object} taskresponse.NotificationResponse
// @Failure 404 {object} taskresponse.ErrorResponse
// @Router /notifications/{id}/mark-as-read [post]
func (c *NotificationController) MarkAsRead(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	notification, err := c.markAsReadNotificationUseCase.Execute(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToNotificationResponse(notification))
}

// MarkAllAsRead marca todas as notificações de um usuário como lidas
// @Summary Marcar todas as notificações como lidas
// @Description Marca todas as notificações de um usuário como lidas
// @Tags notifications
// @Produce json
// @Param user_id query string true "ID do usuário"
// @Success 200 {object} taskresponse.NotificationsListResponse
// @Router /notifications/mark-all-as-read [post]
func (c *NotificationController) MarkAllAsRead(ctx *fiber.Ctx) error {
	userID := ctx.Query("user_id")
	if userID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(taskresponse.ErrorResponse{
			Error: "user_id is required",
		})
	}

	notifications, err := c.markAllAsReadNotificationsUseCase.Execute(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(c.mapper.ToNotificationsListResponse(notifications, int64(len(notifications)), 1, len(notifications)))
}

// DeleteByTaskID deleta todas as notificações de uma tarefa
// @Summary Deletar notificações por tarefa
// @Description Deleta todas as notificações de uma tarefa
// @Tags notifications
// @Produce json
// @Param task_id path string true "ID da tarefa"
// @Success 200 {object} taskresponse.ErrorResponse
// @Router /notifications/by-task/{task_id} [delete]
func (c *NotificationController) DeleteByTaskID(ctx *fiber.Ctx) error {
	taskID := ctx.Params("task_id")
	err := c.deleteNotificationsByTaskIDUseCase.Execute(taskID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(taskresponse.ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.JSON(taskresponse.ErrorResponse{
		Error: "All notifications for task deleted",
	})
}
