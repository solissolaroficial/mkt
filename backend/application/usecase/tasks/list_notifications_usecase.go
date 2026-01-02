package tasks

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// ListNotificationsUseCase lista todas as notificações de um usuário
type ListNotificationsUseCase struct {
	notificationGateway gateway.NotificationGateway
}

// NewListNotificationsUseCase cria um novo ListNotificationsUseCase
func NewListNotificationsUseCase(notificationGateway gateway.NotificationGateway) *ListNotificationsUseCase {
	return &ListNotificationsUseCase{
		notificationGateway: notificationGateway,
	}
}

// Execute lista todas as notificações de um usuário
func (uc *ListNotificationsUseCase) Execute(userID string, page, limit int) ([]*entity.Notification, int64, error) {
	// Parse user ID
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, 0, err
	}

	// Create pagination
	pagination := valueobject.NewPagination(page, limit)

	// Default sort order (created at descending)
	sortOrder, err := valueobject.NewSortOrder("created_at", valueobject.SortDirectionDesc)
	if err != nil {
		return nil, 0, err
	}

	// Find notifications by user ID
	notifications, total, err := uc.notificationGateway.FindByUserID(userUUID, &pagination, []*valueobject.SortOrder{sortOrder})
	if err != nil {
		return nil, 0, err
	}

	return notifications, total, nil
}
