package gateway

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type repMarketingActionGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.RepMarketingActionMapper
}

func NewRepMarketingActionGateway(db *gorm.DB) gateway.RepMarketingActionGateway {
	return &repMarketingActionGatewayImpl{
		db:     db,
		mapper: mapper.NewRepMarketingActionMapper(),
	}
}

func (g *repMarketingActionGatewayImpl) Save(ctx context.Context, action *entity.RepMarketingAction) error {
	actionModel := g.mapper.EntityToModel(action)
	return g.db.WithContext(ctx).Create(actionModel).Error
}

func (g *repMarketingActionGatewayImpl) FindByID(ctx context.Context, id string) (*entity.RepMarketingAction, error) {
	var actionModel model.RepMarketingActionModel
	err := g.db.WithContext(ctx).Where("uuid = ?", id).First(&actionModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrRepMarketingActionNotFound
		}
		return nil, err
	}
	return g.mapper.ModelToEntity(&actionModel)
}

func (g *repMarketingActionGatewayImpl) Update(ctx context.Context, action *entity.RepMarketingAction) error {
	actionModel := g.mapper.EntityToModel(action)
	result := g.db.WithContext(ctx).Where("uuid = ?", actionModel.UUID).Save(actionModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrRepMarketingActionNotFound
	}
	return nil
}

func (g *repMarketingActionGatewayImpl) Delete(ctx context.Context, id string) error {
	result := g.db.WithContext(ctx).Delete(&model.RepMarketingActionModel{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrRepMarketingActionNotFound
	}
	return nil
}

func (g *repMarketingActionGatewayImpl) FindByCriteria(
	ctx context.Context,
	criteria *domain.RepMarketingActionCriteria,
	pagination *valueobject.Pagination,
	sortOrder *valueobject.SortOrder,
) ([]*entity.RepMarketingAction, error) {
	var actionModels []model.RepMarketingActionModel

	query := g.db.WithContext(ctx).Model(&actionModels).Where("deleted_at IS NULL")

	// Aplicar criteria usando getters
	if criteria.RepresentativeUUID() != nil {
		query = query.Where("representative_uuid = ?", *criteria.RepresentativeUUID())
	}
	if criteria.Month() != nil {
		query = query.Where("month = ?", *criteria.Month())
	}

	// Aplicar ordenação
	orderBy := "date DESC"
	if sortOrder != nil && sortOrder.GetField() == "date" {
		if sortOrder.GetDirection() == valueobject.SortDirectionAsc {
			orderBy = "date ASC"
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
	actionModelPointers := make([]*model.RepMarketingActionModel, len(actionModels))
	for i := range actionModels {
		actionModelPointers[i] = &actionModels[i]
	}

	return g.mapper.ModelsToEntities(actionModelPointers)
}

func (g *repMarketingActionGatewayImpl) CountByCriteria(ctx context.Context, criteria *domain.RepMarketingActionCriteria) (int64, error) {
	var count int64

	query := g.db.WithContext(ctx).Model(&model.RepMarketingActionModel{}).Where("deleted_at IS NULL")

	// Aplicar criteria usando getters
	if criteria.RepresentativeUUID() != nil {
		query = query.Where("representative_uuid = ?", *criteria.RepresentativeUUID())
	}
	if criteria.Month() != nil {
		query = query.Where("month = ?", *criteria.Month())
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (g *repMarketingActionGatewayImpl) ExistsByID(ctx context.Context, id string) (bool, error) {
	var count int64
	err := g.db.WithContext(ctx).
		Model(&model.RepMarketingActionModel{}).
		Where("uuid = ? AND deleted_at IS NULL", id).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
