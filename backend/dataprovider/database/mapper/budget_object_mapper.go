package mapper

import (
	"errors"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

// ToBudgetObjectDomain converte Model para Entity
func ToBudgetObjectDomain(model *model.BudgetObjectModel) (*entity.BudgetObject, error) {
	if model == nil {
		return nil, errors.New("model cannot be nil")
	}

	return entity.ReconstructBudgetObject(
		model.UUID,
		model.Code,
		model.Name,
		model.CreatedAt,
		model.UpdatedAt,
		getDeletedAt(model.DeletedAt),
	), nil
}

// ToBudgetObjectModel converte Entity para Model
func ToBudgetObjectModel(entity *entity.BudgetObject) (*model.BudgetObjectModel, error) {
	if entity == nil {
		return nil, errors.New("entity cannot be nil")
	}

	return &model.BudgetObjectModel{
		UUID:      entity.ID(),
		Code:      entity.Code(),
		Name:      entity.Name(),
		CreatedAt: entity.CreatedAt(),
		UpdatedAt: entity.UpdatedAt(),
		DeletedAt: getGormDeletedAt(entity.DeletedAt()),
	}, nil
}

// ToBudgetObjectDomainList converte lista de Model para lista de Entity
func ToBudgetObjectDomainList(models []*model.BudgetObjectModel) ([]*entity.BudgetObject, error) {
	if models == nil {
		return nil, nil
	}

	entities := make([]*entity.BudgetObject, 0, len(models))
	for _, m := range models {
		entity, err := ToBudgetObjectDomain(m)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

// ToBudgetObjectModelList converte lista de Entity para lista de Model
func ToBudgetObjectModelList(entities []*entity.BudgetObject) ([]*model.BudgetObjectModel, error) {
	if entities == nil {
		return nil, nil
	}

	models := make([]*model.BudgetObjectModel, 0, len(entities))
	for _, e := range entities {
		model, err := ToBudgetObjectModel(e)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}
