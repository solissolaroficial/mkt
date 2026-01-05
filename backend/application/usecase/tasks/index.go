package tasks

import (
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// Export all task use cases
func NewUseCases(
	taskGateway gateway.TaskGateway,
	subtaskGateway gateway.SubtaskGateway,
	commentGateway gateway.CommentGateway,
	notificationGateway gateway.NotificationGateway,
	userGateway gateway.UserGateway,
) *UseCases {
	return &UseCases{
		CreateTask: NewCreateTaskUseCase(taskGateway, subtaskGateway),
		UpdateTask: NewUpdateTaskUseCase(taskGateway),
		DeleteTask: NewDeleteTaskUseCase(taskGateway),
		GetTask:    NewGetTaskUseCase(taskGateway),
		ListTasks:  NewListTasksUseCase(taskGateway),
		Subtask: SubtaskUseCases{
			Create: NewCreateSubtaskUseCase(subtaskGateway),
			Update: NewUpdateSubtaskUseCase(subtaskGateway),
			Delete: NewDeleteSubtaskUseCase(subtaskGateway),
			Get:    NewGetSubtaskUseCase(subtaskGateway),
			List:   NewListSubtasksUseCase(subtaskGateway),
		},
		Comment: CommentUseCases{
			Create: NewCreateCommentUseCase(commentGateway, notificationGateway, userGateway),
			Update: NewUpdateCommentUseCase(commentGateway),
			Delete: NewDeleteCommentUseCase(commentGateway),
			Get:    NewGetCommentUseCase(commentGateway),
			List:   NewListCommentsUseCase(commentGateway),
		},
		Notification: NotificationUseCases{
			Create:         NewCreateNotificationUseCase(notificationGateway),
			Update:         NewUpdateNotificationUseCase(notificationGateway),
			Delete:         NewDeleteNotificationUseCase(notificationGateway),
			Get:            NewGetNotificationUseCase(notificationGateway),
			List:           NewListNotificationsUseCase(notificationGateway),
			MarkAsRead:     NewMarkAsReadNotificationUseCase(notificationGateway),
			MarkAllAsRead:  NewMarkAllAsReadNotificationsUseCase(notificationGateway),
			DeleteByTaskID: NewDeleteNotificationsByTaskIDUseCase(notificationGateway),
		},
	}
}

// UseCases agrupa todos os use cases de tarefas
type UseCases struct {
	CreateTask   *CreateTaskUseCase
	UpdateTask   *UpdateTaskUseCase
	DeleteTask   *DeleteTaskUseCase
	GetTask      *GetTaskUseCase
	ListTasks    *ListTasksUseCase
	Subtask      SubtaskUseCases
	Comment      CommentUseCases
	Notification NotificationUseCases
}

// SubtaskUseCases agrupa os use cases de subtarefas
type SubtaskUseCases struct {
	Create *CreateSubtaskUseCase
	Update *UpdateSubtaskUseCase
	Delete *DeleteSubtaskUseCase
	Get    *GetSubtaskUseCase
	List   *ListSubtasksUseCase
}

// CommentUseCases agrupa os use cases de comentários
type CommentUseCases struct {
	Create *CreateCommentUseCase
	Update *UpdateCommentUseCase
	Delete *DeleteCommentUseCase
	Get    *GetCommentUseCase
	List   *ListCommentsUseCase
}

// NotificationUseCases agrupa os use cases de notificações
type NotificationUseCases struct {
	Create         *CreateNotificationUseCase
	Update         *UpdateNotificationUseCase
	Delete         *DeleteNotificationUseCase
	Get            *GetNotificationUseCase
	List           *ListNotificationsUseCase
	MarkAsRead     *MarkAsReadNotificationUseCase
	MarkAllAsRead  *MarkAllAsReadNotificationsUseCase
	DeleteByTaskID *DeleteNotificationsByTaskIDUseCase
}
