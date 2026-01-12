package gateway

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"solis/backend/core/domain"
	"solis/backend/core/domain/entity"
	domainErrors "solis/backend/core/domain/errors"
	"solis/backend/core/domain/gateway"
	"solis/backend/core/domain/valueobject"
	"solis/backend/dataprovider/database/mapper"
	"solis/backend/dataprovider/database/model"
)

type showroomItemGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.ShowroomItemMapper
}

func NewShowroomItemGateway(db *gorm.DB) gateway.ShowroomItemGateway {
	return &showroomItemGatewayImpl{
		db:     db,
		mapper: mapper.NewShowroomItemMapper(),
	}
}

func (g *showroomItemGatewayImpl) Save(ctx context.Context, item *entity.ShowroomItem) error {
	itemModel := g.mapper.EntityToModel(item)
	return g.db.WithContext(ctx).Create(itemModel).Error
}

func (g *showroomItemGatewayImpl) FindByID(ctx context.Context, id string) (*entity.ShowroomItem, error) {
	var itemModel model.ShowroomItemModel
	err := g.db.WithContext(ctx).Where("uuid = ?", id).First(&itemModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrShowroomItemNotFound
		}
		return nil, err
	}
	return g.mapper.ModelToEntity(&itemModel)
}

func (g *showroomItemGatewayImpl) Update(ctx context.Context, item *entity.ShowroomItem) error {
	itemModel := g.mapper.EntityToModel(item)
	result := g.db.WithContext(ctx).Where("uuid = ?", itemModel.UUID).Save(itemModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrShowroomItemNotFound
	}
	return nil
}

func (g *showroomItemGatewayImpl) Delete(ctx context.Context, id string) error {
	result := g.db.WithContext(ctx).Delete(&model.ShowroomItemModel{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrShowroomItemNotFound
	}
	return nil
}

func (g *showroomItemGatewayImpl) FindByCriteria(
	ctx context.Context,
	criteria *domain.ShowroomItemCriteria,
	pagination *valueobject.Pagination,
	sortOrder *valueobject.SortOrder,
) ([]*entity.ShowroomItem, error) {
	var itemModels []model.ShowroomItemModel

	query := g.db.WithContext(ctx).Model(&itemModels).Where("deleted_at IS NULL")

	// Aplicar criteria usando getters
	if criteria.RepName() != nil {
		query = query.Where("rep_name = ?", *criteria.RepName())
	}
	if criteria.City() != nil {
		query = query.Where("city = ?", *criteria.City())
	}
	if criteria.Delivered() != nil {
		query = query.Where("delivered = ?", *criteria.Delivered())
	}

	// Aplicar ordenação
	orderBy := "created_at DESC"
	if sortOrder != nil && sortOrder.GetField() == "delivery_forecast" {
		if sortOrder.GetDirection() == valueobject.SortDirectionAsc {
			orderBy = "delivery_forecast ASC"
		}
	}
	query = query.Order(orderBy)

	// Aplicar paginação
	if pagination != nil {
		offset := pagination.Offset()
		limit := pagination.PageSize
		query = query.Offset(offset).Limit(limit)
	}

	if err := query.Find(&itemModels).Error; err != nil {
		return nil, err
	}

	// Convert slice of models to slice of pointers
	itemModelPointers := make([]*model.ShowroomItemModel, len(itemModels))
	for i := range itemModels {
		itemModelPointers[i] = &itemModels[i]
	}

	return g.mapper.ModelsToEntities(itemModelPointers)
}

func (g *showroomItemGatewayImpl) CountByCriteria(ctx context.Context, criteria *domain.ShowroomItemCriteria) (int64, error) {
	var count int64

	query := g.db.WithContext(ctx).Model(&model.ShowroomItemModel{}).Where("deleted_at IS NULL")

	// Aplicar criteria usando getters
	if criteria.RepName() != nil {
		query = query.Where("rep_name = ?", *criteria.RepName())
	}
	if criteria.City() != nil {
		query = query.Where("city = ?", *criteria.City())
	}
	if criteria.Delivered() != nil {
		query = query.Where("delivered = ?", *criteria.Delivered())
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (g *showroomItemGatewayImpl) ExistsByID(ctx context.Context, id string) (bool, error) {
	var count int64
	err := g.db.WithContext(ctx).
		Model(&model.ShowroomItemModel{}).
		Where("uuid = ? AND deleted_at IS NULL", id).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
