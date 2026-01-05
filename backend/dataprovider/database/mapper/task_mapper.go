package mapper

import (
	"encoding/json"

	"github.com/seu-usuario/solis-backend/core/domain/constants"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

// TaskMapper converte entre Task entity e Task model
type TaskMapper struct{}

// NewTaskMapper cria um novo TaskMapper
func NewTaskMapper() *TaskMapper {
	return &TaskMapper{}
}

// ToEntity converte um Task model para Task entity
func (m *TaskMapper) ToEntity(model *model.Task) (*entity.Task, error) {
	// Unmarshal JSON Flows
	var flows []string
	if model.Flows != "" {
		if err := json.Unmarshal([]byte(model.Flows), &flows); err != nil {
			return nil, err
		}
	}

	return entity.ReconstructTask(
		model.UUID,
		model.Title,
		model.Description,
		model.StartDate,
		model.DueDate,
		constants.TaskPriority(model.Priority),
		constants.TaskStatus(model.Status),
		constants.TaskCategory(model.Category),
		model.AssigneeID,
		flows,
		model.Archived,
		model.SortOrder,
		model.CreatedAt,
		model.UpdatedAt,
		nil, // deletedAt não é usado na entity (soft delete no model)
	), nil
}

// ToModel converte uma Task entity para Task model
func (m *TaskMapper) ToModel(entity *entity.Task) (*model.Task, error) {
	// Marshal JSON Flows
	flowsJSON, err := json.Marshal(entity.Flows())
	if err != nil {
		return nil, err
	}

	return &model.Task{
		UUID:        entity.ID(),
		Title:       entity.Title(),
		Description: entity.Description(),
		StartDate:   entity.StartDate(),
		DueDate:     entity.DueDate(),
		Priority:    string(entity.Priority()),
		Status:      string(entity.Status()),
		Category:    string(entity.Category()),
		AssigneeID:  entity.AssigneeID(),
		Flows:       string(flowsJSON),
		Archived:    entity.Archived(),
		SortOrder:   entity.SortOrder(),
		CreatedAt:   entity.CreatedAt(),
		UpdatedAt:   entity.UpdatedAt(),
	}, nil
}

// ToEntityList converte uma lista de Task models para Task entities
func (m *TaskMapper) ToEntityList(models []*model.Task) ([]*entity.Task, error) {
	entities := make([]*entity.Task, 0, len(models))
	for _, model := range models {
		entity, err := m.ToEntity(model)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}
