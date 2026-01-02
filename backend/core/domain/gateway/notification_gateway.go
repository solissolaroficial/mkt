package gateway

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"

	"github.com/google/uuid"
)

// NotificationGateway define a interface para persistência de notificações
type NotificationGateway interface {
	// Create cria uma nova notificação
	Create(notification *entity.Notification) error

	// Update atualiza uma notificação existente
	Update(notification *entity.Notification) error

	// Delete deleta uma notificação
	Delete(id uuid.UUID) error

	// FindByID busca uma notificação pelo ID
	FindByID(id uuid.UUID) (*entity.Notification, error)

	// FindAll busca todas as notificações com paginação
	FindAll(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Notification, int64, error)

	// FindByUserID busca notificações por user ID
	FindByUserID(userID uuid.UUID, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Notification, int64, error)

	// FindUnreadByUserID busca notificações não lidas por user ID
	FindUnreadByUserID(userID uuid.UUID, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Notification, int64, error)

	// MarkAsRead marca uma notificação como lida
	MarkAsRead(id uuid.UUID) error

	// MarkAllAsReadByUserID marca todas as notificações de um usuário como lidas
	MarkAllAsReadByUserID(userID uuid.UUID) error

	// DeleteByTaskID deleta todas as notificações de uma tarefa
	DeleteByTaskID(taskID uuid.UUID) error
}
