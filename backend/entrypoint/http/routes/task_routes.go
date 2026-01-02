package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seu-usuario/solis-backend/entrypoint/http/controller"
)

// SetupTaskRoutes configura as rotas de Tasks
func SetupTaskRoutes(router fiber.Router, taskController *controller.TaskController) {
	tasks := router.Group("/tasks")

	// CRUD de Tasks
	tasks.Post("/", taskController.Create)
	tasks.Get("/:id", taskController.Get)
	tasks.Put("/:id", taskController.Update)
	tasks.Delete("/:id", taskController.Delete)
	tasks.Get("/", taskController.List)
}

// SetupSubtaskRoutes configura as rotas de Subtasks
func SetupSubtaskRoutes(router fiber.Router, subtaskController *controller.SubtaskController) {
	subtasks := router.Group("/subtasks")

	// CRUD de Subtasks
	subtasks.Post("/", subtaskController.Create)
	subtasks.Get("/:id", subtaskController.Get)
	subtasks.Put("/:id", subtaskController.Update)
	subtasks.Delete("/:id", subtaskController.Delete)
	subtasks.Get("/", subtaskController.List)
}

// SetupCommentRoutes configura as rotas de Comments
func SetupCommentRoutes(router fiber.Router, commentController *controller.CommentController) {
	comments := router.Group("/comments")

	// CRUD de Comments
	comments.Post("/", commentController.Create)
	comments.Get("/:id", commentController.Get)
	comments.Put("/:id", commentController.Update)
	comments.Delete("/:id", commentController.Delete)
	comments.Get("/", commentController.List)
}

// SetupNotificationRoutes configura as rotas de Notifications
func SetupNotificationRoutes(router fiber.Router, notificationController *controller.NotificationController) {
	notifications := router.Group("/notifications")

	// CRUD de Notifications
	notifications.Post("/", notificationController.Create)
	notifications.Get("/:id", notificationController.Get)
	notifications.Put("/:id", notificationController.Update)
	notifications.Delete("/:id", notificationController.Delete)
	notifications.Get("/", notificationController.List)

	// Rotas especiais de Notifications
	notifications.Post("/:id/mark-as-read", notificationController.MarkAsRead)
	notifications.Post("/mark-all-as-read", notificationController.MarkAllAsRead)
	notifications.Delete("/by-task/:task_id", notificationController.DeleteByTaskID)
}
