package mapper

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type RepresentativeMonthlyGoalMapper struct{}

func NewRepresentativeMonthlyGoalMapper() *RepresentativeMonthlyGoalMapper {
	return &RepresentativeMonthlyGoalMapper{}
}

// ModelToEntity converts Model to Entity
func (m *RepresentativeMonthlyGoalMapper) ModelToEntity(model *model.RepresentativeMonthlyGoalModel) (*entity.RepresentativeMonthlyGoal, error) {
	return entity.ReconstructRepresentativeMonthlyGoal(
		model.ID,
		model.RepresentativeUUID,
		model.Month,
		model.Year,
		model.Target,
		model.Realized,
		model.CreatedAt,
		model.UpdatedAt,
		model.DeletedAt,
	)
}

// ModelsToEntities converts slice of Model to slice of Entity
func (m *RepresentativeMonthlyGoalMapper) ModelsToEntities(models []*model.RepresentativeMonthlyGoalModel) ([]*entity.RepresentativeMonthlyGoal, error) {
	items := make([]*entity.RepresentativeMonthlyGoal, len(models))
	for i, model := range models {
		item, err := m.ModelToEntity(model)
		if err != nil {
			return nil, err
		}
		items[i] = item
	}
	return items, nil
}

// EntityToModel converts Entity to Model
func (m *RepresentativeMonthlyGoalMapper) EntityToModel(goal *entity.RepresentativeMonthlyGoal) *model.RepresentativeMonthlyGoalModel {
	return &model.RepresentativeMonthlyGoalModel{
		ID:                 goal.ID(),
		RepresentativeUUID: goal.RepresentativeUUID(),
		Month:              goal.Month(),
		Year:               goal.Year(),
		Target:             goal.Target(),
		Realized:           goal.Realized(),
		CreatedAt:          goal.CreatedAt(),
		UpdatedAt:          goal.UpdatedAt(),
		DeletedAt:          goal.DeletedAt(),
	}
}
