package gateway

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	kpiErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type monthlyDataGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.MonthlyDataMapper
}

// NewMonthlyDataGateway creates a new MonthlyDataGateway implementation
func NewMonthlyDataGateway(db *gorm.DB) gateway.MonthlyDataGateway {
	return &monthlyDataGatewayImpl{
		db:     db,
		mapper: &mapper.MonthlyDataMapper{},
	}
}

// Save creates new monthly data or updates an existing one
func (g *monthlyDataGatewayImpl) Save(ctx context.Context, data *entity.MonthlyData) error {
	dataModel, err := g.mapper.ToModel(data)
	if err != nil {
		return err
	}

	// Check if monthly data already exists to decide between Create and Update
	var existingData model.MonthlyData
	err = g.db.WithContext(ctx).Where("uuid = ?", dataModel.UUID).First(&existingData).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Create new monthly data
		if err := g.db.WithContext(ctx).Create(dataModel).Error; err != nil {
			return err
		}
	} else {
		// Update existing monthly data
		if err := g.db.WithContext(ctx).Save(dataModel).Error; err != nil {
			return err
		}
	}

	return nil
}

// FindByKpiAndMonth retrieves monthly data for a specific KPI, year and month
func (g *monthlyDataGatewayImpl) FindByKpiAndMonth(ctx context.Context, kpiID uuid.UUID, year int, month string) (*entity.MonthlyData, error) {
	var dataModel model.MonthlyData

	err := g.db.WithContext(ctx).
		Where("kpi_category_id = ? AND year = ? AND month = ?", kpiID, year, month).
		First(&dataModel).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, kpiErrors.ErrMonthDataNotFound
		}
		return nil, err
	}

	return g.mapper.ToDomain(&dataModel)
}

// FindAllByKpi retrieves all monthly data for a specific KPI
func (g *monthlyDataGatewayImpl) FindAllByKpi(ctx context.Context, kpiID uuid.UUID) ([]*entity.MonthlyData, error) {
	var dataModels []model.MonthlyData

	err := g.db.WithContext(ctx).
		Where("kpi_category_id = ?", kpiID).
		Order("month").
		Find(&dataModels).Error

	if err != nil {
		return nil, err
	}

	// Convert slice of models to slice of entities
	dataEntities := make([]*entity.MonthlyData, len(dataModels))
	for i, dataModel := range dataModels {
		dataEntity, err := g.mapper.ToDomain(&dataModel)
		if err != nil {
			return nil, err
		}
		dataEntities[i] = dataEntity
	}

	return dataEntities, nil
}

// Update modifies existing monthly data
func (g *monthlyDataGatewayImpl) Update(ctx context.Context, data *entity.MonthlyData) error {
	dataModel, err := g.mapper.ToModel(data)
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

// UpdateMeta updates only the meta value for a specific KPI, year and month
func (g *monthlyDataGatewayImpl) UpdateMeta(ctx context.Context, kpiID uuid.UUID, year int, month string, meta float64) error {
	result := g.db.WithContext(ctx).
		Model(&model.MonthlyData{}).
		Where("kpi_category_id = ? AND year = ? AND month = ?", kpiID, year, month).
		Update("meta", meta)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return kpiErrors.ErrMonthDataNotFound
	}

	return nil
}

// UpdateRealized updates only the realized value for a specific KPI, year and month
func (g *monthlyDataGatewayImpl) UpdateRealized(ctx context.Context, kpiID uuid.UUID, year int, month string, realized float64) error {
	result := g.db.WithContext(ctx).
		Model(&model.MonthlyData{}).
		Where("kpi_category_id = ? AND year = ? AND month = ?", kpiID, year, month).
		Update("realized", realized)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return kpiErrors.ErrMonthDataNotFound
	}

	return nil
}

// UpdateBreakdown updates only the breakdown data for a specific KPI, year and month
func (g *monthlyDataGatewayImpl) UpdateBreakdown(ctx context.Context, kpiID uuid.UUID, year int, month string, breakdown interface{}) error {
	result := g.db.WithContext(ctx).
		Model(&model.MonthlyData{}).
		Where("kpi_category_id = ? AND year = ? AND month = ?", kpiID, year, month).
		Update("breakdown", breakdown)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return kpiErrors.ErrMonthDataNotFound
	}

	return nil
}
