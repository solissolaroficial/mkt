package mapper

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type GiftItemMapper struct{}

func NewGiftItemMapper() *GiftItemMapper {
	return &GiftItemMapper{}
}

// ModelToEntity converte Model para Entity
func (m *GiftItemMapper) ModelToEntity(model *model.GiftItemModel) (*entity.GiftItem, error) {
	// Converter deletedAt
	var deletedAt *time.Time
	if model.DeletedAt.Valid {
		deletedAt = &model.DeletedAt.Time
	}

	return entity.ReconstructGiftItem(
		model.UUID,
		model.Name,
		model.Stock,
		model.Price,
		model.CreatedAt,
		model.UpdatedAt,
		deletedAt,
	), nil
}

// ModelsToEntities converte slice de Model para slice de Entity
func (m *GiftItemMapper) ModelsToEntities(models []*model.GiftItemModel) ([]*entity.GiftItem, error) {
	items := make([]*entity.GiftItem, len(models))
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
func (m *GiftItemMapper) EntityToModel(item *entity.GiftItem) *model.GiftItemModel {
	// Converter deletedAt
	var deletedAt gorm.DeletedAt
	if item.DeletedAt() != nil {
		deletedAt.Time = *item.DeletedAt()
		deletedAt.Valid = true
	}

	return &model.GiftItemModel{
		UUID:      item.ID(),
		Name:      item.Name().Value(),
		Stock:     item.Stock(),
		Price:     item.Price().Value(),
		CreatedAt: item.CreatedAt(),
		UpdatedAt: item.UpdatedAt(),
		DeletedAt: deletedAt,
	}
}
