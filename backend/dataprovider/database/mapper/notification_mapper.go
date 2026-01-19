package mapper

import (
	"github.com/seu-usuario/solis-backend/core/domain/constants"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

// NotificationMapper converte entre Notification entity e Notification model
type NotificationMapper struct{}

// NewNotificationMapper cria um novo NotificationMapper
func NewNotificationMapper() *NotificationMapper {
	return &NotificationMapper{}
}

// ToEntity converte um Notification model para Notification entity
func (m *NotificationMapper) ToEntity(model *model.Notification) *entity.Notification {
	return entity.ReconstructNotification(
		model.UUID,
		model.UserUUID,
		model.TaskUUID,
		constants.NotificationType(model.Type),
		model.Title,
		model.Message,
		model.Read,
		model.Archived,
		model.Timestamp,
	)
}

// ToModel converte uma Notification entity para Notification model
func (m *NotificationMapper) ToModel(entity *entity.Notification) *model.Notification {
	notificationType := entity.Type()
	return &model.Notification{
		UUID:      entity.ID(),
		UserUUID:  entity.UserUUID(),
		TaskUUID:  entity.TaskUUID(),
		Type:      string(notificationType),
		Title:     entity.Title(),
		Message:   entity.Message(),
		Read:      entity.Read(),
		Archived:  entity.Archived(),
		Timestamp: entity.Timestamp(),
	}
}

// ToEntityList converte uma lista de Notification models para Notification entities
func (m *NotificationMapper) ToEntityList(models []*model.Notification) []*entity.Notification {
	entities := make([]*entity.Notification, 0, len(models))
	for _, model := range models {
		entities = append(entities, m.ToEntity(model))
	}
	return entities
}
