package mapper

import (
	"time"

	"gorm.io/gorm"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

// RecurrentPdvMapper converte entre Model e Entity para PDVs recorrentes
type RecurrentPdvMapper struct{}

// NewRecurrentPdvMapper cria uma nova instância do RecurrentPdvMapper
func NewRecurrentPdvMapper() *RecurrentPdvMapper {
	return &RecurrentPdvMapper{}
}

// ModelToEntity converte Model para Entity
func (m *RecurrentPdvMapper) ModelToEntity(model *model.RecurrentPdvModel) (*entity.RecurrentPdv, error) {
	// Converter deletedAt
	var deletedAt *time.Time
	if model.DeletedAt.Valid {
		deletedAt = &model.DeletedAt.Time
	}

	return entity.ReconstructRecurrentPdv(
		model.UUID,
		model.Name,
		model.RepName,
		model.City,
		model.Followers,
		model.InstagramProfile,
		model.CreatedAt,
		model.UpdatedAt,
		deletedAt,
	), nil
}

// ModelsToEntities converte slice de Model para slice de Entity
func (m *RecurrentPdvMapper) ModelsToEntities(models []*model.RecurrentPdvModel) ([]*entity.RecurrentPdv, error) {
	pdvs := make([]*entity.RecurrentPdv, len(models))
	for i, model := range models {
		pdv, err := m.ModelToEntity(model)
		if err != nil {
			return nil, err
		}
		pdvs[i] = pdv
	}
	return pdvs, nil
}

// EntityToModel converte Entity para Model
func (m *RecurrentPdvMapper) EntityToModel(pdv *entity.RecurrentPdv) *model.RecurrentPdvModel {
	// Converter deletedAt
	var deletedAt gorm.DeletedAt
	if pdv.DeletedAt() != nil {
		deletedAt.Time = *pdv.DeletedAt()
		deletedAt.Valid = true
	}

	return &model.RecurrentPdvModel{
		UUID:             pdv.ID(),
		Name:             pdv.Name(),
		RepName:          pdv.RepName(),
		City:             pdv.City(),
		Followers:        pdv.Followers(),
		InstagramProfile: pdv.InstagramProfile(),
		CreatedAt:        pdv.CreatedAt(),
		UpdatedAt:        pdv.UpdatedAt(),
		DeletedAt:        deletedAt,
	}
}
