package tasks

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetNotificationUseCase busca uma notificação por ID
type GetNotificationUseCase struct {
	notificationGateway gateway.NotificationGateway
}

// NewGetNotificationUseCase cria um novo GetNotificationUseCase
func NewGetNotificationUseCase(notificationGateway gateway.NotificationGateway) *GetNotificationUseCase {
	return &GetNotificationUseCase{
		notificationGateway: notificationGateway,
	}
}

// Execute busca uma notificação por ID
func (uc *GetNotificationUseCase) Execute(id string) (*entity.Notification, error) {
	// Parse ID
	notificationID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid notification ID: %w", err)
	}

	// Find notification
	notification, err := uc.notificationGateway.FindByID(notificationID)
	if err != nil {
		return nil, err
	}

	return notification, nil
}
