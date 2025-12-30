package mapper

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type KpiMapper struct{}

// ToDomain converte Model (GORM) para Entity (Domínio)
func (m *KpiMapper) ToDomain(kpiModel *model.KpiCategory) (*entity.KpiCategory, error) {
	// Convert MonthlyData models to entities
	monthlyDataEntities := make([]*entity.MonthlyData, len(kpiModel.MonthlyData))
	monthlyDataMapper := &MonthlyDataMapper{}

	for i, monthlyDataModel := range kpiModel.MonthlyData {
		monthlyDataEntity, err := monthlyDataMapper.ToDomain(&monthlyDataModel)
		if err != nil {
			return nil, err
		}
		monthlyDataEntities[i] = monthlyDataEntity
	}

	return entity.ReconstructKpiCategoryWithMonthlyData(
		kpiModel.UUID,
		kpiModel.Title,
		kpiModel.Color,
		kpiModel.Unit,
		monthlyDataEntities,
		kpiModel.CreatedAt,
		kpiModel.UpdatedAt,
	)
}

// ToModel converte Entity (Domínio) para Model (GORM)
func (m *KpiMapper) ToModel(kpi *entity.KpiCategory) (*model.KpiCategory, error) {
	// Convert MonthlyData entities to models
	monthlyDataModels := make([]model.MonthlyData, len(kpi.MonthlyDatas()))
	monthlyDataMapper := &MonthlyDataMapper{}

	for i, monthlyDataEntity := range kpi.MonthlyDatas() {
		monthlyDataModel, err := monthlyDataMapper.ToModel(monthlyDataEntity)
		if err != nil {
			return nil, err
		}
		monthlyDataModels[i] = *monthlyDataModel
	}

	return &model.KpiCategory{
		UUID:        kpi.ID(),
		Title:       kpi.Title(),
		Slug:        kpi.Slug(),
		Color:       kpi.Color(),
		Unit:        kpi.Unit(),
		MonthlyData: monthlyDataModels,
		CreatedAt:   kpi.CreatedAt(),
		UpdatedAt:   kpi.UpdatedAt(),
	}, nil
}
