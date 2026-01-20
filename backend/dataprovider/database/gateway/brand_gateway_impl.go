package gateway

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type brandGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.BrandMapper
}

func NewBrandGateway(db *gorm.DB) gateway.BrandGateway {
	return &brandGatewayImpl{
		db:     db,
		mapper: mapper.NewBrandMapper(),
	}
}

func (g *brandGatewayImpl) Save(ctx context.Context, brand *entity.Brand) error {
	brandModel := g.mapper.EntityToModel(brand)
	return g.db.WithContext(ctx).Create(brandModel).Error
}

func (g *brandGatewayImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Brand, error) {
	var brandModel model.BrandModel
	err := g.db.WithContext(ctx).
		Where("uuid = ? AND deleted_at IS NULL", id).
		First(&brandModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrBrandNotFound
		}
		return nil, err
	}
	return g.mapper.ModelToEntity(&brandModel)
}

func (g *brandGatewayImpl) FindByName(ctx context.Context, name string) (*entity.Brand, error) {
	var brandModel model.BrandModel
	err := g.db.WithContext(ctx).
		Where("name = ? AND deleted_at IS NULL", name).
		First(&brandModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrBrandNotFound
		}
		return nil, err
	}
	return g.mapper.ModelToEntity(&brandModel)
}

func (g *brandGatewayImpl) List(ctx context.Context) ([]*entity.Brand, error) {
	var brandModels []model.BrandModel
	err := g.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Find(&brandModels).Error
	if err != nil {
		return nil, err
	}

	brands := make([]*entity.Brand, len(brandModels))
	for i, m := range brandModels {
		brand, err := g.mapper.ModelToEntity(&m)
		if err != nil {
			return nil, err
		}
		brands[i] = brand
	}
	return brands, nil
}

func (g *brandGatewayImpl) Update(ctx context.Context, brand *entity.Brand) error {
	brandModel := g.mapper.EntityToModel(brand)
	result := g.db.WithContext(ctx).
		Where("uuid = ? AND deleted_at IS NULL", brandModel.UUID).
		Save(brandModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrBrandNotFound
	}
	return nil
}

func (g *brandGatewayImpl) Delete(ctx context.Context, id uuid.UUID) error {
	// Soft delete: atualiza deleted_at em vez de remover o registro
	result := g.db.WithContext(ctx).
		Model(&model.BrandModel{}).
		Where("uuid = ? AND deleted_at IS NULL", id).
		Update("deleted_at", gorm.Expr("NOW()"))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrBrandNotFound
	}
	return nil
}

func (g *brandGatewayImpl) ExistsByName(ctx context.Context, name string) (bool, error) {
	var count int64
	err := g.db.WithContext(ctx).
		Model(&model.BrandModel{}).
		Where("name = ? AND deleted_at IS NULL", name).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
