package tasks

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// DeleteNotificationsByTaskIDUseCase deleta todas as notificações de uma tarefa
type DeleteNotificationsByTaskIDUseCase struct {
	notificationGateway gateway.NotificationGateway
}

// NewDeleteNotificationsByTaskIDUseCase cria um novo DeleteNotificationsByTaskIDUseCase
func NewDeleteNotificationsByTaskIDUseCase(notificationGateway gateway.NotificationGateway) *DeleteNotificationsByTaskIDUseCase {
	return &DeleteNotificationsByTaskIDUseCase{
		notificationGateway: notificationGateway,
	}
}

// Execute deleta todas as notificações de uma tarefa
func (uc *DeleteNotificationsByTaskIDUseCase) Execute(taskID string) error {
	// Parse task ID
	taskUUID, err := uuid.Parse(taskID)
	if err != nil {
		return err
	}

	// Delete notifications by task ID
	return uc.notificationGateway.DeleteByTaskID(taskUUID)
}
