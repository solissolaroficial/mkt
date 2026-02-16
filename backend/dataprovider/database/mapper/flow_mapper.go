package mapper

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

// FlowMapper converte entre Flow entity e Flow model
type FlowMapper struct{}

// NewFlowMapper cria um novo FlowMapper
func NewFlowMapper() *FlowMapper {
	return &FlowMapper{}
}

// ToEntity converte um Flow model para Flow entity
func (m *FlowMapper) ToEntity(model *model.Flow) *entity.Flow {
	var deletedAt *time.Time
	if model.DeletedAt.Valid {
		deletedAt = &model.DeletedAt.Time
	}
	return entity.ReconstructFlow(
		model.UUID,
		model.Name,
		model.Description,
		model.Color,
		model.SortOrder,
		model.CreatedAt,
		model.UpdatedAt,
		deletedAt,
	)
}

// ToModel converte uma Flow entity para Flow model
func (m *FlowMapper) ToModel(entity *entity.Flow) *model.Flow {
	return &model.Flow{
		UUID:        entity.ID(),
		Name:        entity.Name(),
		Description: entity.Description(),
		Color:       entity.Color(),
		SortOrder:   entity.SortOrder(),
		CreatedAt:   entity.CreatedAt(),
		UpdatedAt:   entity.UpdatedAt(),
		DeletedAt:   gorm.DeletedAt{}, // Zero value for GORM
	}
}

// ToEntityList converte uma lista de Flow models para Flow entities
func (m *FlowMapper) ToEntityList(models []*model.Flow) []*entity.Flow {
	entities := make([]*entity.Flow, 0, len(models))
	for _, model := range models {
		entities = append(entities, m.ToEntity(model))
	}
	return entities
}
