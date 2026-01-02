package tasks

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// MarkAllAsReadNotificationsUseCase marca todas as notificações de um usuário como lidas
type MarkAllAsReadNotificationsUseCase struct {
	notificationGateway gateway.NotificationGateway
}

// NewMarkAllAsReadNotificationsUseCase cria um novo MarkAllAsReadNotificationsUseCase
func NewMarkAllAsReadNotificationsUseCase(notificationGateway gateway.NotificationGateway) *MarkAllAsReadNotificationsUseCase {
	return &MarkAllAsReadNotificationsUseCase{
		notificationGateway: notificationGateway,
	}
}

// Execute marca todas as notificações de um usuário como lidas
func (uc *MarkAllAsReadNotificationsUseCase) Execute(userID string) ([]*entity.Notification, error) {
	// Parse user ID
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	// Default pagination (no limit, no offset)
	pagination := valueobject.NewPagination(0, 0)

	// Default sort order (created at descending)
	sortOrder, err := valueobject.NewSortOrder("created_at", valueobject.SortDirectionDesc)
	if err != nil {
		return nil, err
	}

	// Find all unread notifications
	notifications, _, err := uc.notificationGateway.FindUnreadByUserID(userUUID, &pagination, []*valueobject.SortOrder{sortOrder})
	if err != nil {
		return nil, err
	}

	// Mark all as read
	for _, notification := range notifications {
		notification.MarkAsRead()
		if err := uc.notificationGateway.Update(notification); err != nil {
			return nil, err
		}
	}

	return notifications, nil
}
