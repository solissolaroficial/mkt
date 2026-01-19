package mapper

import (
	"errors"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

// ToBudgetGroupDomain converte Model para Entity
func ToBudgetGroupDomain(model *model.BudgetGroupModel) (*entity.BudgetGroup, error) {
	if model == nil {
		return nil, errors.New("model cannot be nil")
	}

	return entity.ReconstructBudgetGroup(
		model.UUID,
		model.Code,
		model.Name,
		model.CreatedAt,
		model.UpdatedAt,
		getDeletedAt(model.DeletedAt),
	), nil
}

// ToBudgetGroupModel converte Entity para Model
func ToBudgetGroupModel(entity *entity.BudgetGroup) (*model.BudgetGroupModel, error) {
	if entity == nil {
		return nil, errors.New("entity cannot be nil")
	}

	return &model.BudgetGroupModel{
		UUID:      entity.ID(),
		Code:      entity.Code(),
		Name:      entity.Name(),
		CreatedAt: entity.CreatedAt(),
		UpdatedAt: entity.UpdatedAt(),
		DeletedAt: getGormDeletedAt(entity.DeletedAt()),
	}, nil
}

// ToBudgetGroupDomainList converte lista de Model para lista de Entity
func ToBudgetGroupDomainList(models []*model.BudgetGroupModel) ([]*entity.BudgetGroup, error) {
	if models == nil {
		return nil, nil
	}

	entities := make([]*entity.BudgetGroup, 0, len(models))
	for _, m := range models {
		entity, err := ToBudgetGroupDomain(m)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

// ToBudgetGroupModelList converte lista de Entity para lista de Model
func ToBudgetGroupModelList(entities []*entity.BudgetGroup) ([]*model.BudgetGroupModel, error) {
	if entities == nil {
		return nil, nil
	}

	models := make([]*model.BudgetGroupModel, 0, len(entities))
	for _, e := range entities {
		model, err := ToBudgetGroupModel(e)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}
