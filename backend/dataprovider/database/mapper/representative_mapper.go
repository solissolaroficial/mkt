package mapper

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type RepresentativeMapper struct{}

func NewRepresentativeMapper() *RepresentativeMapper {
	return &RepresentativeMapper{}
}

// ModelToEntity converts Model to Entity
func (m *RepresentativeMapper) ModelToEntity(model *model.RepresentativeModel) (*entity.Representative, error) {
	return entity.ReconstructRepresentative(
		model.UUID,
		valueobject.ReconstructRepresentativeCode(model.Code),
		model.Name,
		model.Email,
		model.Phone,
		model.Company,
		model.Region,
		model.City,
		model.Attendant,
		model.Active,
		model.CreatedAt,
		model.UpdatedAt,
		model.DeletedAt,
	), nil
}

// ModelsToEntities converts slice of Model to slice of Entity
func (m *RepresentativeMapper) ModelsToEntities(models []*model.RepresentativeModel) ([]*entity.Representative, error) {
	items := make([]*entity.Representative, len(models))
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
func (m *RepresentativeMapper) EntityToModel(representative *entity.Representative) *model.RepresentativeModel {
	return &model.RepresentativeModel{
		UUID:      representative.UUID(),
		Code:      representative.Code().Value(),
		Name:      representative.Name(),
		Email:     representative.Email(),
		Phone:     representative.Phone(),
		Company:   representative.Company(),
		Region:    representative.Region(),
		City:      representative.City(),
		Attendant: representative.Attendant(),
		Active:    representative.Active(),
		CreatedAt: representative.CreatedAt(),
		UpdatedAt: representative.UpdatedAt(),
		DeletedAt: representative.DeletedAt(),
	}
}
