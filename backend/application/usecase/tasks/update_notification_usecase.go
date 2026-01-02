package tasks

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// UpdateNotificationUseCase atualiza uma notificação existente
type UpdateNotificationUseCase struct {
	notificationGateway gateway.NotificationGateway
}

// NewUpdateNotificationUseCase cria um novo UpdateNotificationUseCase
func NewUpdateNotificationUseCase(notificationGateway gateway.NotificationGateway) *UpdateNotificationUseCase {
	return &UpdateNotificationUseCase{
		notificationGateway: notificationGateway,
	}
}

// Execute atualiza uma notificação
func (uc *UpdateNotificationUseCase) Execute(id string, title string, message string) error {
	// Parse ID
	notificationID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid notification ID: %w", err)
	}

	// Find existing notification
	notification, err := uc.notificationGateway.FindByID(notificationID)
	if err != nil {
		return err
	}

	// Update notification via Reconstruct
	updatedNotification := entity.ReconstructNotification(
		notification.ID(),
		notification.UserID(),
		notification.TaskID(),
		notification.Type(),
		title,
		message,
		notification.Read(),
		notification.Archived(),
		notification.Timestamp(),
	)

	// Save notification
	return uc.notificationGateway.Update(updatedNotification)
}
