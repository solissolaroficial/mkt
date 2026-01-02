package mapper

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

// SubtaskMapper converte entre Subtask entity e Subtask model
type SubtaskMapper struct{}

// NewSubtaskMapper cria um novo SubtaskMapper
func NewSubtaskMapper() *SubtaskMapper {
	return &SubtaskMapper{}
}

// ToEntity converte um Subtask model para Subtask entity
func (m *SubtaskMapper) ToEntity(model *model.Subtask) *entity.Subtask {
	return entity.ReconstructSubtask(
		model.UUID,
		model.TaskID,
		model.Title,
		model.Completed,
		model.AssigneeID,
		model.DueDate,
		model.CreatedAt,
		model.UpdatedAt,
		nil, // deletedAt não é usado na entity (soft delete no model)
	)
}

// ToModel converte uma Subtask entity para Subtask model
func (m *SubtaskMapper) ToModel(entity *entity.Subtask) *model.Subtask {
	return &model.Subtask{
		UUID:       entity.ID(),
		TaskID:     entity.TaskID(),
		Title:      entity.Title(),
		Completed:  entity.Completed(),
		AssigneeID: entity.AssigneeID(),
		DueDate:    entity.DueDate(),
		CreatedAt:  entity.CreatedAt(),
		UpdatedAt:  entity.UpdatedAt(),
	}
}

// ToEntityList converte uma lista de Subtask models para Subtask entities
func (m *SubtaskMapper) ToEntityList(models []*model.Subtask) []*entity.Subtask {
	entities := make([]*entity.Subtask, 0, len(models))
	for _, model := range models {
		entities = append(entities, m.ToEntity(model))
	}
	return entities
}
