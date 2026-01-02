package tasks

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/constants"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// CreateNotificationUseCase cria uma nova notificação
type CreateNotificationUseCase struct {
	notificationGateway gateway.NotificationGateway
}

// NewCreateNotificationUseCase cria um novo CreateNotificationUseCase
func NewCreateNotificationUseCase(notificationGateway gateway.NotificationGateway) *CreateNotificationUseCase {
	return &CreateNotificationUseCase{
		notificationGateway: notificationGateway,
	}
}

// Execute cria uma nova notificação
func (uc *CreateNotificationUseCase) Execute(
	userID string,
	title string,
	message string,
	notificationType string,
	taskID string,
) (*entity.Notification, error) {
	// Parse user ID
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	// Parse task ID (optional)
	var taskUUID *uuid.UUID
	if taskID != "" {
		parsedTaskID, err := uuid.Parse(taskID)
		if err != nil {
			return nil, err
		}
		taskUUID = &parsedTaskID
	}

	// Convert notification type string to constants.NotificationType
	notifType := constants.NotificationType(notificationType)

	// Create notification
	notification, err := entity.NewNotification(
		userUUID,
		taskUUID,
		notifType,
		title,
		message,
	)
	if err != nil {
		return nil, err
	}

	// Save notification
	if err := uc.notificationGateway.Create(notification); err != nil {
		return nil, err
	}

	return notification, nil
}
