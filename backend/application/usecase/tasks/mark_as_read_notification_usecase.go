package tasks

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// MarkAsReadNotificationUseCase marca uma notificação como lida
type MarkAsReadNotificationUseCase struct {
	notificationGateway gateway.NotificationGateway
}

// NewMarkAsReadNotificationUseCase cria um novo MarkAsReadNotificationUseCase
func NewMarkAsReadNotificationUseCase(notificationGateway gateway.NotificationGateway) *MarkAsReadNotificationUseCase {
	return &MarkAsReadNotificationUseCase{
		notificationGateway: notificationGateway,
	}
}

// Execute marca uma notificação como lida
func (uc *MarkAsReadNotificationUseCase) Execute(id string) (*entity.Notification, error) {
	// Parse ID
	notificationID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid notification ID: %w", err)
	}

	// Find existing notification
	notification, err := uc.notificationGateway.FindByID(notificationID)
	if err != nil {
		return nil, err
	}

	// Mark as read
	notification.MarkAsRead()

	// Save notification
	if err := uc.notificationGateway.Update(notification); err != nil {
		return nil, err
	}

	return notification, nil
}
