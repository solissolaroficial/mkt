package mapper

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	budgeterrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

// ToBudgetItemDomain converte Model para Entity
// Realiza validações dos dados antes de criar a entidade
// ObjectName e GroupName são populados via JOIN com BudgetObject e BudgetGroup
func ToBudgetItemDomain(model *model.BudgetItemModel) (*entity.BudgetItem, error) {
	if model == nil {
		return nil, errors.New("model cannot be nil")
	}

	vals, err := parseFloatArrayFromJSON(model.Vals)
	if err != nil {
		return nil, fmt.Errorf("failed to parse vals: %w", err)
	}

	realizedVals, err := parseFloatArrayFromJSON(model.RealizedVals)
	if err != nil {
		return nil, fmt.Errorf("failed to parse realizedVals: %w", err)
	}

	// Popula ObjectName e GroupName dos relacionamentos
	var objectName string
	if model.Object != nil {
		objectName = model.Object.Name
	}

	var groupName string
	if model.Group != nil {
		groupName = model.Group.Name
	}

	return entity.ReconstructBudgetItem(
		model.UUID,
		model.ObjectUUID,
		objectName,
		model.GroupUUID,
		groupName,
		model.Cod,
		model.Desc,
		vals,
		realizedVals,
		model.Year,
		model.CreatedAt,
		model.UpdatedAt,
		getDeletedAt(model.DeletedAt),
	), nil
}

// ToBudgetItemModel converte Entity para Model
func ToBudgetItemModel(entity *entity.BudgetItem) (*model.BudgetItemModel, error) {
	if entity == nil {
		return nil, errors.New("entity cannot be nil")
	}

	vals, err := floatArrayToJSON(entity.Vals())
	if err != nil {
		return nil, fmt.Errorf("failed to serialize vals: %w", err)
	}

	realizedVals, err := floatArrayToJSON(entity.RealizedVals())
	if err != nil {
		return nil, fmt.Errorf("failed to serialize realizedVals: %w", err)
	}

	return &model.BudgetItemModel{
		UUID:         entity.ID(),
		ObjectUUID:   entity.ObjectUUID(),
		GroupUUID:    entity.GroupUUID(),
		Cod:          entity.Cod(),
		Desc:         entity.Desc(),
		Vals:         vals,
		RealizedVals: realizedVals,
		Year:         entity.Year(),
		CreatedAt:    entity.CreatedAt(),
		UpdatedAt:    entity.UpdatedAt(),
		DeletedAt:    getGormDeletedAt(entity.DeletedAt()),
	}, nil
}

// ToBudgetItemDomainList converte lista de Model para lista de Entity
// Se houver erro em algum item, retorna o erro e nil
func ToBudgetItemDomainList(models []*model.BudgetItemModel) ([]*entity.BudgetItem, error) {
	if models == nil {
		return nil, nil
	}

	entities := make([]*entity.BudgetItem, 0, len(models))
	for _, m := range models {
		entity, err := ToBudgetItemDomain(m)
		if err != nil {
			return nil, fmt.Errorf("failed to convert model %s: %w", m.UUID, err)
		}
		entities = append(entities, entity)
	}

	return entities, nil
}

// ToBudgetItemModelList converte lista de Entity para lista de Model
// Se houver erro em algum item, retorna o erro e nil
func ToBudgetItemModelList(entities []*entity.BudgetItem) ([]*model.BudgetItemModel, error) {
	if entities == nil {
		return nil, nil
	}

	models := make([]*model.BudgetItemModel, 0, len(entities))
	for _, e := range entities {
		model, err := ToBudgetItemModel(e)
		if err != nil {
			return nil, fmt.Errorf("failed to convert entity %s: %w", e.ID(), err)
		}
		models = append(models, model)
	}

	return models, nil
}

// parseFloatArrayFromJSON converte JSON para array de float
// Valida tamanho (12) e valores não-negativos
func parseFloatArrayFromJSON(jsonData []byte) ([]float64, error) {
	if len(jsonData) == 0 {
		return make([]float64, 12), nil
	}

	var vals []float64
	if err := json.Unmarshal(jsonData, &vals); err != nil {
		return nil, errors.New("failed to parse JSON array")
	}

	// Validar tamanho
	if len(vals) != 12 {
		return nil, budgeterrors.ErrInvalidMonthCount
	}

	// Validar valores não-negativos
	for i, v := range vals {
		if v < 0 {
			return nil, fmt.Errorf("%w at index %d", budgeterrors.ErrNegativeValue, i)
		}
	}

	return vals, nil
}

// floatArrayToJSON converte array de float para JSON
// Valida tamanho (12)
func floatArrayToJSON(vals []float64) ([]byte, error) {
	if len(vals) != 12 {
		return nil, budgeterrors.ErrInvalidMonthCount
	}

	return json.Marshal(vals)
}

// getDeletedAt converte gorm.DeletedAt para *time.Time
func getDeletedAt(deletedAt gorm.DeletedAt) *time.Time {
	if !deletedAt.Valid {
		return nil
	}
	return &deletedAt.Time
}

// getGormDeletedAt converte *time.Time para gorm.DeletedAt
func getGormDeletedAt(deletedAt *time.Time) gorm.DeletedAt {
	if deletedAt == nil {
		return gorm.DeletedAt{}
	}
	return gorm.DeletedAt{Time: *deletedAt, Valid: true}
}
