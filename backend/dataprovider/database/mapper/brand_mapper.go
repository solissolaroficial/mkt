package mapper

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type BrandMapper struct{}

func NewBrandMapper() *BrandMapper {
	return &BrandMapper{}
}

func (m *BrandMapper) ModelToEntity(model *model.BrandModel) (*entity.Brand, error) {
	var deletedAt *time.Time
	if model.DeletedAt != nil {
		deletedAt = model.DeletedAt
	}

	return entity.ReconstructBrand(
		model.UUID,
		model.Name,
		model.CreatedAt,
		model.UpdatedAt,
		deletedAt,
	), nil
}

func (m *BrandMapper) ModelsToEntities(models []model.BrandModel) ([]*entity.Brand, error) {
	brands := make([]*entity.Brand, len(models))
	for i, model := range models {
		brand, err := m.ModelToEntity(&model)
		if err != nil {
			return nil, err
		}
		brands[i] = brand
	}
	return brands, nil
}

func (m *BrandMapper) EntityToModel(brand *entity.Brand) *model.BrandModel {
	return &model.BrandModel{
		UUID:      brand.UUID(),
		Name:      brand.Name(),
		CreatedAt: brand.CreatedAt(),
		UpdatedAt: brand.UpdatedAt(),
		DeletedAt: brand.DeletedAt(),
	}
}
