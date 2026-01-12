package mapper

import (
	"time"

	"gorm.io/gorm"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type RepMarketingActionMapper struct{}

func NewRepMarketingActionMapper() *RepMarketingActionMapper {
	return &RepMarketingActionMapper{}
}

// ModelToEntity converte Model para Entity
func (m *RepMarketingActionMapper) ModelToEntity(model *model.RepMarketingActionModel) (*entity.RepMarketingAction, error) {
	// Converter deletedAt
	var deletedAt *time.Time
	if model.DeletedAt.Valid {
		deletedAt = &model.DeletedAt.Time
	}

	return entity.ReconstructRepMarketingAction(
		model.UUID,
		model.RepName,
		model.Date,
		model.Description,
		model.Month,
		model.CreatedAt,
		model.UpdatedAt,
		deletedAt,
	), nil
}

// ModelsToEntities converte slice de Model para slice de Entity
func (m *RepMarketingActionMapper) ModelsToEntities(models []*model.RepMarketingActionModel) ([]*entity.RepMarketingAction, error) {
	actions := make([]*entity.RepMarketingAction, len(models))
	for i, model := range models {
		action, err := m.ModelToEntity(model)
		if err != nil {
			return nil, err
		}
		actions[i] = action
	}
	return actions, nil
}

// EntityToModel converte Entity para Model
func (m *RepMarketingActionMapper) EntityToModel(action *entity.RepMarketingAction) *model.RepMarketingActionModel {
	// Converter deletedAt
	var deletedAt gorm.DeletedAt
	if action.DeletedAt() != nil {
		deletedAt.Time = *action.DeletedAt()
		deletedAt.Valid = true
	}

	return &model.RepMarketingActionModel{
		UUID:        action.ID(),
		RepName:     action.RepName(),
		Date:        action.Date(),
		Description: action.Description(),
		Month:       action.Month(),
		CreatedAt:   action.CreatedAt(),
		UpdatedAt:   action.UpdatedAt(),
		DeletedAt:   deletedAt,
	}
}
