package mapper

import (
	"time"

	"gorm.io/gorm"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type ShowroomItemMapper struct{}

func NewShowroomItemMapper() *ShowroomItemMapper {
	return &ShowroomItemMapper{}
}

// ModelToEntity converte Model para Entity
func (m *ShowroomItemMapper) ModelToEntity(model *model.ShowroomItemModel) (*entity.ShowroomItem, error) {
	// Converter deletedAt
	var deletedAt *time.Time
	if model.DeletedAt.Valid {
		deletedAt = &model.DeletedAt.Time
	}

	return entity.ReconstructShowroomItem(
		model.UUID,
		model.PDV,
		model.City,
		model.Contact,
		model.RepName,
		model.DeliveryForecast,
		model.WorkshopDate,
		model.Delivered,
		model.CreatedAt,
		model.UpdatedAt,
		deletedAt,
	), nil
}

// ModelsToEntities converte slice de Model para slice de Entity
func (m *ShowroomItemMapper) ModelsToEntities(models []*model.ShowroomItemModel) ([]*entity.ShowroomItem, error) {
	items := make([]*entity.ShowroomItem, len(models))
	for i, model := range models {
		item, err := m.ModelToEntity(model)
		if err != nil {
			return nil, err
		}
		items[i] = item
	}
	return items, nil
}

// EntityToModel converte Entity para Model
func (m *ShowroomItemMapper) EntityToModel(item *entity.ShowroomItem) *model.ShowroomItemModel {
	// Converter deletedAt
	var deletedAt gorm.DeletedAt
	if item.DeletedAt() != nil {
		deletedAt.Time = *item.DeletedAt()
		deletedAt.Valid = true
	}

	return &model.ShowroomItemModel{
		UUID:             item.ID(),
		PDV:              item.PDV(),
		City:             item.City(),
		Contact:          item.Contact(),
		RepName:          item.RepName(),
		DeliveryForecast: item.DeliveryForecast(),
		WorkshopDate:     item.WorkshopDate(),
		Delivered:        item.Delivered(),
		CreatedAt:        item.CreatedAt(),
		UpdatedAt:        item.UpdatedAt(),
		DeletedAt:        deletedAt,
	}
}
