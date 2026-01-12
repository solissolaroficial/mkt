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

type offlineActionGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.OfflineActionMapper
}

func NewOfflineActionGateway(db *gorm.DB) gateway.OfflineActionGateway {
	return &offlineActionGatewayImpl{
		db:     db,
		mapper: mapper.NewOfflineActionMapper(),
	}
}

func (g *offlineActionGatewayImpl) Save(ctx context.Context, action *entity.OfflineAction) error {
	actionModel := g.mapper.EntityToModel(action)
	return g.db.WithContext(ctx).Create(actionModel).Error
}

func (g *offlineActionGatewayImpl) FindByID(ctx context.Context, id string) (*entity.OfflineAction, error) {
	var actionModel model.OfflineActionModel
	err := g.db.WithContext(ctx).Where("uuid = ?", id).First(&actionModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrOfflineActionNotFound
		}
		return nil, err
	}
	return g.mapper.ModelToEntity(&actionModel)
}

func (g *offlineActionGatewayImpl) Update(ctx context.Context, action *entity.OfflineAction) error {
	actionModel := g.mapper.EntityToModel(action)
	result := g.db.WithContext(ctx).Where("uuid = ?", actionModel.UUID).Save(actionModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrOfflineActionNotFound
	}
	return nil
}

func (g *offlineActionGatewayImpl) Delete(ctx context.Context, id string) error {
	result := g.db.WithContext(ctx).Delete(&model.OfflineActionModel{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrOfflineActionNotFound
	}
	return nil
}

func (g *offlineActionGatewayImpl) FindByCriteria(
	ctx context.Context,
	criteria *domain.OfflineActionCriteria,
	pagination *valueobject.Pagination,
	sortOrder *valueobject.SortOrder,
) ([]*entity.OfflineAction, error) {
	var actionModels []model.OfflineActionModel

	query := g.db.WithContext(ctx).Model(&actionModels).Where("deleted_at IS NULL")

	// Aplicar criteria usando getters (sem dependência de GORM no domínio)
	if criteria.Category() != nil {
		query = query.Where("category = ?", criteria.Category().String())
	}
	if criteria.RepName() != nil {
		query = query.Where("rep_name = ?", *criteria.RepName())
	}
	if criteria.Month() != nil {
		query = query.Where("month = ?", *criteria.Month())
	}
	if criteria.StartDate() != nil {
		query = query.Where("action_date >= ?", *criteria.StartDate())
	}
	if criteria.EndDate() != nil {
		query = query.Where("action_date <= ?", *criteria.EndDate())
	}
	if criteria.Status() != nil {
		query = query.Where("status = ?", criteria.Status().String())
	}

	// Aplicar ordenação
	orderBy := "action_date DESC"
	if sortOrder != nil && sortOrder.GetField() == "action_date" {
		if sortOrder.GetDirection() == valueobject.SortDirectionAsc {
			orderBy = "action_date ASC"
		}
	}
	query = query.Order(orderBy)

	// Aplicar paginação
	if pagination != nil {
		offset := pagination.Offset()
		limit := pagination.PageSize
		query = query.Offset(offset).Limit(limit)
	}

	if err := query.Find(&actionModels).Error; err != nil {
		return nil, err
	}

	// Convert slice of models to slice of pointers
	actionModelPointers := make([]*model.OfflineActionModel, len(actionModels))
	for i := range actionModels {
		actionModelPointers[i] = &actionModels[i]
	}

	return g.mapper.ModelsToEntities(actionModelPointers)
}

func (g *offlineActionGatewayImpl) CountByCriteria(ctx context.Context, criteria *domain.OfflineActionCriteria) (int64, error) {
	var count int64

	query := g.db.WithContext(ctx).Model(&model.OfflineActionModel{}).Where("deleted_at IS NULL")

	// Aplicar criteria usando getters
	if criteria.Category() != nil {
		query = query.Where("category = ?", criteria.Category().String())
	}
	if criteria.RepName() != nil {
		query = query.Where("rep_name = ?", *criteria.RepName())
	}
	if criteria.Month() != nil {
		query = query.Where("month = ?", *criteria.Month())
	}
	if criteria.StartDate() != nil {
		query = query.Where("action_date >= ?", *criteria.StartDate())
	}
	if criteria.EndDate() != nil {
		query = query.Where("action_date <= ?", *criteria.EndDate())
	}
	if criteria.Status() != nil {
		query = query.Where("status = ?", criteria.Status().String())
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (g *offlineActionGatewayImpl) ExistsByID(ctx context.Context, id string) (bool, error) {
	var count int64
	err := g.db.WithContext(ctx).
		Model(&model.OfflineActionModel{}).
		Where("uuid = ? AND deleted_at IS NULL", id).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
