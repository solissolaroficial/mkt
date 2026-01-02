package tasks

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// DeleteNotificationUseCase deleta uma notificação existente
type DeleteNotificationUseCase struct {
	notificationGateway gateway.NotificationGateway
}

// NewDeleteNotificationUseCase cria um novo DeleteNotificationUseCase
func NewDeleteNotificationUseCase(notificationGateway gateway.NotificationGateway) *DeleteNotificationUseCase {
	return &DeleteNotificationUseCase{
		notificationGateway: notificationGateway,
	}
}

// Execute deleta uma notificação
func (uc *DeleteNotificationUseCase) Execute(id string) error {
	// Parse ID
	notificationID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid notification ID: %w", err)
	}

	// Delete notification
	return uc.notificationGateway.Delete(notificationID)
}
