package mapper

import (
	"time"

	"gorm.io/gorm"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type RecurrentPdvMapper struct{}

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
		model.RepresentativeUUID,
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
	recurrentPdvs := make([]*entity.RecurrentPdv, len(models))
	for i, model := range models {
		recurrentPdv, err := m.ModelToEntity(model)
		if err != nil {
			return nil, err
		}
		recurrentPdvs[i] = recurrentPdv
	}
	return recurrentPdvs, nil
}

// EntityToModel converte Entity para Model
func (m *RecurrentPdvMapper) EntityToModel(recurrentPdv *entity.RecurrentPdv) *model.RecurrentPdvModel {
	// Converter deletedAt
	var deletedAt gorm.DeletedAt
	if recurrentPdv.DeletedAt() != nil {
		deletedAt.Time = *recurrentPdv.DeletedAt()
		deletedAt.Valid = true
	}

	return &model.RecurrentPdvModel{
		UUID:               recurrentPdv.ID(),
		RepresentativeUUID: recurrentPdv.RepresentativeUUID(),
		Name:               recurrentPdv.Name(),
		City:               recurrentPdv.City(),
		Followers:          recurrentPdv.Followers(),
		InstagramProfile:   recurrentPdv.InstagramProfile(),
		CreatedAt:          recurrentPdv.CreatedAt(),
		UpdatedAt:          recurrentPdv.UpdatedAt(),
		DeletedAt:          deletedAt,
	}
}
