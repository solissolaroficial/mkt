package gateway

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	kpiErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type kpiGatewayImpl struct {
	db            *gorm.DB
	mapper        *mapper.KpiMapper
	monthlyMapper *mapper.MonthlyDataMapper
}

// NewKpiGateway creates a new KpiGateway implementation
func NewKpiGateway(db *gorm.DB) gateway.KpiGateway {
	return &kpiGatewayImpl{
		db:            db,
		mapper:        &mapper.KpiMapper{},
		monthlyMapper: &mapper.MonthlyDataMapper{},
	}
}

// Save creates a new KPI category or updates an existing one
func (g *kpiGatewayImpl) Save(ctx context.Context, kpi *entity.KpiCategory) error {
	kpiModel, err := g.mapper.ToModel(kpi)
	if err != nil {
		return err
	}

	// Check if KPI already exists to decide between Create and Update
	var existingKpi model.KpiCategory
	err = g.db.WithContext(ctx).Where("uuid = ?", kpiModel.UUID).First(&existingKpi).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Create new KPI
		if err := g.db.WithContext(ctx).Create(kpiModel).Error; err != nil {
			return err
		}
	} else {
		// Update existing KPI
		if err := g.db.WithContext(ctx).Save(kpiModel).Error; err != nil {
			return err
		}
	}

	return nil
}

// FindByID retrieves a KPI category by its ID
func (g *kpiGatewayImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.KpiCategory, error) {
	var kpiModel model.KpiCategory

	err := g.db.WithContext(ctx).
		Preload("MonthlyData").
		Where("uuid = ?", id).
		First(&kpiModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kpiErrors.ErrKpiNotFound
		}
		return nil, err
	}

	return g.mapper.ToDomain(&kpiModel)
}

// FindByTitle retrieves a KPI category by its title
func (g *kpiGatewayImpl) FindByTitle(ctx context.Context, title string) (*entity.KpiCategory, error) {
	var kpiModel model.KpiCategory

	err := g.db.WithContext(ctx).
		Preload("MonthlyData").
		Where("title = ?", title).
		First(&kpiModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil instead of error for seeder
		}
		return nil, err
	}

	return g.mapper.ToDomain(&kpiModel)
}

// FindBySlug retrieves a KPI category by its slug
func (g *kpiGatewayImpl) FindBySlug(ctx context.Context, slug string) (*entity.KpiCategory, error) {
	var kpiModel model.KpiCategory

	err := g.db.WithContext(ctx).
		Preload("MonthlyData").
		Where("slug = ?", slug).
		First(&kpiModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kpiErrors.ErrKpiNotFound
		}
		return nil, err
	}

	return g.mapper.ToDomain(&kpiModel)
}

// FindBySlugs retrieves KPI categories by a list of slugs
func (g *kpiGatewayImpl) FindBySlugs(ctx context.Context, slugs []string) ([]*entity.KpiCategory, error) {
	var kpiModels []model.KpiCategory

	err := g.db.WithContext(ctx).
		Preload("MonthlyData").
		Where("slug IN ?", slugs).
		Find(&kpiModels).Error

	if err != nil {
		return nil, err
	}

	// Convert slice of models to slice of entities
	kpiEntities := make([]*entity.KpiCategory, len(kpiModels))
	for i, kpiModel := range kpiModels {
		kpiEntity, err := g.mapper.ToDomain(&kpiModel)
		if err != nil {
			return nil, err
		}
		kpiEntities[i] = kpiEntity
	}

	return kpiEntities, nil
}

// FindAll retrieves KPI categories with pagination
func (g *kpiGatewayImpl) FindAll(ctx context.Context, pagination valueobject.Pagination) ([]*entity.KpiCategory, int64, error) {
	var kpiModels []model.KpiCategory
	var total int64

	// Get total count
	if err := g.db.WithContext(ctx).Model(&model.KpiCategory{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	err := g.db.WithContext(ctx).
		Preload("MonthlyData").
		Offset(pagination.Offset()).
		Limit(pagination.Limit()).
		Find(&kpiModels).Error

	if err != nil {
		return nil, 0, err
	}

	// Convert slice of models to slice of entities
	kpiEntities := make([]*entity.KpiCategory, len(kpiModels))
	for i, kpiModel := range kpiModels {
		kpiEntity, err := g.mapper.ToDomain(&kpiModel)
		if err != nil {
			return nil, 0, err
		}
		kpiEntities[i] = kpiEntity
	}

	return kpiEntities, total, nil
}

// Update modifies an existing KPI category
func (g *kpiGatewayImpl) Update(ctx context.Context, kpi *entity.KpiCategory) error {
	kpiModel, err := g.mapper.ToModel(kpi)
	if err != nil {
		return err
	}

	result := g.db.WithContext(ctx).Where("uuid = ?", kpiModel.UUID).Save(kpiModel)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return kpiErrors.ErrKpiNotFound
	}

	return nil
}

// Delete removes a KPI category by its ID
func (g *kpiGatewayImpl) Delete(ctx context.Context, id uuid.UUID) error {
	result := g.db.WithContext(ctx).Delete(&model.KpiCategory{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return kpiErrors.ErrKpiNotFound
	}

	return nil
}

// FindMonthlyDataByID retrieves monthly data by its ID
func (g *kpiGatewayImpl) FindMonthlyDataByID(ctx context.Context, id uuid.UUID) (*entity.MonthlyData, error) {
	var dataModel model.MonthlyData

	err := g.db.WithContext(ctx).
		Where("uuid = ? AND deleted_at IS NULL", id).
		First(&dataModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kpiErrors.ErrMonthDataNotFound
		}
		return nil, err
	}

	return g.monthlyMapper.ToDomain(&dataModel)
}

// UpdateMonthlyData modifies an existing monthly data
func (g *kpiGatewayImpl) UpdateMonthlyData(ctx context.Context, monthlyData *entity.MonthlyData) error {
	dataModel, err := g.monthlyMapper.ToModel(monthlyData)
	if err != nil {
		return err
	}

	result := g.db.WithContext(ctx).Where("uuid = ?", dataModel.UUID).Save(dataModel)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return kpiErrors.ErrMonthDataNotFound
	}

	return nil
}

// DeleteMonthlyData soft deletes a monthly data record by its ID
func (g *kpiGatewayImpl) DeleteMonthlyData(ctx context.Context, id uuid.UUID) error {
	result := g.db.WithContext(ctx).Delete(&model.MonthlyData{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return kpiErrors.ErrMonthDataNotFound
	}

	return nil
}
