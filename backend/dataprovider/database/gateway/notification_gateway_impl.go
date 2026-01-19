package gateway

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	notificationerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type notificationGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.NotificationMapper
}

// NewNotificationGateway cria uma nova implementação de NotificationGateway
func NewNotificationGateway(db *gorm.DB) gateway.NotificationGateway {
	return &notificationGatewayImpl{
		db:     db,
		mapper: &mapper.NotificationMapper{},
	}
}

// Create cria uma nova notificação
func (g *notificationGatewayImpl) Create(notification *entity.Notification) error {
	notificationModel := g.mapper.ToModel(notification)

	if err := g.db.Create(notificationModel).Error; err != nil {
		return err
	}

	return nil
}

// Update atualiza uma notificação existente
func (g *notificationGatewayImpl) Update(notification *entity.Notification) error {
	notificationModel := g.mapper.ToModel(notification)

	result := g.db.Where("uuid = ?", notificationModel.UUID).Save(notificationModel)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &notificationerrors.NotificationNotFoundError{ID: notification.ID().String()}
	}

	return nil
}

// Delete deleta uma notificação
func (g *notificationGatewayImpl) Delete(id uuid.UUID) error {
	result := g.db.Delete(&model.Notification{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &notificationerrors.NotificationNotFoundError{ID: id.String()}
	}

	return nil
}

// FindByID busca uma notificação pelo ID
func (g *notificationGatewayImpl) FindByID(id uuid.UUID) (*entity.Notification, error) {
	var notificationModel model.Notification

	err := g.db.
		Preload("User").
		Preload("Task").
		Where("uuid = ?", id).
		First(&notificationModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &notificationerrors.NotificationNotFoundError{ID: id.String()}
		}
		return nil, err
	}

	return g.mapper.ToEntity(&notificationModel), nil
}

// FindAll busca todas as notificações com paginação
func (g *notificationGatewayImpl) FindAll(pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Notification, int64, error) {
	var notificationModels []model.Notification
	var total int64

	// Get total count
	if err := g.db.Model(&model.Notification{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Notification{}).
		Preload("User").
		Preload("Task")

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&notificationModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	notificationModelPointers := make([]*model.Notification, len(notificationModels))
	for i := range notificationModels {
		notificationModelPointers[i] = &notificationModels[i]
	}

	entities := g.mapper.ToEntityList(notificationModelPointers)
	return entities, total, nil
}

// FindByUserID busca notificações por user ID
func (g *notificationGatewayImpl) FindByUserID(userID uuid.UUID, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Notification, int64, error) {
	var notificationModels []model.Notification
	var total int64

	// Get total count
	if err := g.db.Model(&model.Notification{}).Where("user_uuid = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Notification{}).
		Preload("User").
		Preload("Task").
		Where("user_uuid = ?", userID)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&notificationModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	notificationModelPointers := make([]*model.Notification, len(notificationModels))
	for i := range notificationModels {
		notificationModelPointers[i] = &notificationModels[i]
	}

	entities := g.mapper.ToEntityList(notificationModelPointers)
	return entities, total, nil
}

// FindUnreadByUserID busca notificações não lidas por user ID
func (g *notificationGatewayImpl) FindUnreadByUserID(userID uuid.UUID, pagination *valueobject.Pagination, sortOrders []*valueobject.SortOrder) ([]*entity.Notification, int64, error) {
	var notificationModels []model.Notification
	var total int64

	// Get total count
	if err := g.db.Model(&model.Notification{}).Where("user_uuid = ? AND read = ?", userID, false).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Build query
	query := g.db.Model(&model.Notification{}).
		Preload("User").
		Preload("Task").
		Where("user_uuid = ? AND read = ?", userID, false)

	// Apply sorting
	for _, sortOrder := range sortOrders {
		query = query.Order(sortOrder.ToSQLString())
	}

	// Get paginated results
	err := query.
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&notificationModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of pointers
	notificationModelPointers := make([]*model.Notification, len(notificationModels))
	for i := range notificationModels {
		notificationModelPointers[i] = &notificationModels[i]
	}

	entities := g.mapper.ToEntityList(notificationModelPointers)
	return entities, total, nil
}

// MarkAsRead marca uma notificação como lida
func (g *notificationGatewayImpl) MarkAsRead(id uuid.UUID) error {
	now := time.Now()
	result := g.db.Model(&model.Notification{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"read":    true,
			"read_at": &now,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return &notificationerrors.NotificationNotFoundError{ID: id.String()}
	}

	return nil
}

// MarkAllAsReadByUserID marca todas as notificações de um usuário como lidas
func (g *notificationGatewayImpl) MarkAllAsReadByUserID(userID uuid.UUID) error {
	now := time.Now()
	result := g.db.Model(&model.Notification{}).
		Where("user_uuid = ? AND read = ?", userID, false).
		Updates(map[string]interface{}{
			"read":    true,
			"read_at": &now,
		})

	return result.Error
}

// DeleteByTaskID deleta todas as notificações de uma tarefa
func (g *notificationGatewayImpl) DeleteByTaskID(taskID uuid.UUID) error {
	result := g.db.Delete(&model.Notification{}, "task_uuid = ?", taskID)
	return result.Error
}
