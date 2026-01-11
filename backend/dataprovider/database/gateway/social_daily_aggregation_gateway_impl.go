package gateway

import (
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type SocialDailyAggregationGatewayImpl struct {
	db *gorm.DB
}

func NewSocialDailyAggregationGatewayImpl(db *gorm.DB) gateway.SocialDailyAggregationGateway {
	return &SocialDailyAggregationGatewayImpl{
		db: db,
	}
}

func (g *SocialDailyAggregationGatewayImpl) Create(aggregation *entity.SocialDailyAggregation) error {
	model := mapper.EntityToSocialDailyAggregationModel(aggregation)
	return g.db.Create(model).Error
}

func (g *SocialDailyAggregationGatewayImpl) Update(aggregation *entity.SocialDailyAggregation) error {
	model := mapper.EntityToSocialDailyAggregationModel(aggregation)
	return g.db.Save(model).Error
}

func (g *SocialDailyAggregationGatewayImpl) Delete(id uuid.UUID) error {
	return g.db.Delete(&model.SocialDailyAggregationModel{}, id).Error
}

func (g *SocialDailyAggregationGatewayImpl) GetByID(id uuid.UUID) (*entity.SocialDailyAggregation, error) {
	var model model.SocialDailyAggregationModel
	err := g.db.Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return mapper.ModelToSocialDailyAggregationEntity(&model), nil
}

func (g *SocialDailyAggregationGatewayImpl) GetByBrandAndDate(
	brandName string,
	date time.Time,
) (*entity.SocialDailyAggregation, error) {
	var model model.SocialDailyAggregationModel
	err := g.db.Where("brand_name = ?", brandName).
		Where("aggregation_date = ?", date).
		First(&model).Error

	if err != nil {
		return nil, err
	}

	return mapper.ModelToSocialDailyAggregationEntity(&model), nil
}

func (g *SocialDailyAggregationGatewayImpl) List(
	criteria *domain.SocialDailyAggregationCriteria,
	pagination *valueobject.Pagination,
) ([]*entity.SocialDailyAggregation, int64, error) {
	var models []*model.SocialDailyAggregationModel
	var total int64

	query := g.db.Model(&model.SocialDailyAggregationModel{})

	// Aplicar filtros
	if criteria.BrandName != nil {
		query = query.Where("brand_name = ?", *criteria.BrandName)
	}

	if criteria.StartDate != nil {
		query = query.Where("aggregation_date >= ?", *criteria.StartDate)
	}

	if criteria.EndDate != nil {
		query = query.Where("aggregation_date <= ?", *criteria.EndDate)
	}

	// Contar total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Aplicar paginação
	offset := (pagination.Page - 1) * pagination.PageSize
	query = query.Offset(offset).Limit(pagination.PageSize)

	// Ordenar por data desc
	query = query.Order("aggregation_date DESC")

	// Executar query
	if err := query.Find(&models).Error; err != nil {
		return nil, 0, err
	}

	// Converter para entities
	entities := make([]*entity.SocialDailyAggregation, len(models))
	for i, m := range models {
		entities[i] = mapper.ModelToSocialDailyAggregationEntity(m)
	}

	return entities, total, nil
}

// GetByBrand busca todas as agregações de uma marca
func (g *SocialDailyAggregationGatewayImpl) GetByBrand(brandName string) ([]*entity.SocialDailyAggregation, error) {
	var models []*model.SocialDailyAggregationModel

	query := g.db.Where("brand_name = ?", brandName).
		Order("aggregation_date DESC")

	if err := query.Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*entity.SocialDailyAggregation, len(models))
	for i, m := range models {
		entities[i] = mapper.ModelToSocialDailyAggregationEntity(m)
	}

	return entities, nil
}
