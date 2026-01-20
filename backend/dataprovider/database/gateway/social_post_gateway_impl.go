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

type SocialPostGatewayImpl struct {
	db *gorm.DB
}

func NewSocialPostGatewayImpl(db *gorm.DB) gateway.SocialPostGateway {
	return &SocialPostGatewayImpl{
		db: db,
	}
}

func (g *SocialPostGatewayImpl) Create(post *entity.SocialPost) error {
	model := mapper.EntityToSocialPostModel(post)
	return g.db.Create(model).Error
}

func (g *SocialPostGatewayImpl) Update(post *entity.SocialPost) error {
	model := mapper.EntityToSocialPostModel(post)
	return g.db.Save(model).Error
}

func (g *SocialPostGatewayImpl) Delete(id uuid.UUID) error {
	return g.db.Delete(&model.SocialPostModel{}, id).Error
}

func (g *SocialPostGatewayImpl) GetByID(id uuid.UUID) (*entity.SocialPost, error) {
	var model model.SocialPostModel
	err := g.db.Where("id = ?", id).Preload("Brand").First(&model).Error
	if err != nil {
		return nil, err
	}
	return mapper.ModelToSocialPostEntity(&model), nil
}

func (g *SocialPostGatewayImpl) List(
	criteria *domain.SocialPostCriteria,
	pagination *valueobject.Pagination,
) ([]*entity.SocialPost, int64, error) {
	var models []*model.SocialPostModel
	var total int64

	query := g.db.Model(&model.SocialPostModel{}).Preload("Brand")

	// Aplicar filtros
	if criteria.BrandName != nil {
		query = query.Where("brand_name = ?", *criteria.BrandName)
	}

	if criteria.Platform != nil {
		query = query.Where("platform = ?", criteria.Platform.String())
	}

	if criteria.PostType != nil {
		query = query.Where("post_type = ?", criteria.PostType.String())
	}

	if criteria.StartDate != nil {
		query = query.Where("post_date >= ?", *criteria.StartDate)
	}

	if criteria.EndDate != nil {
		query = query.Where("post_date <= ?", *criteria.EndDate)
	}

	if criteria.MinFollowers != nil {
		query = query.Where("followers_at_post >= ?", *criteria.MinFollowers)
	}

	if criteria.MaxFollowers != nil {
		query = query.Where("followers_at_post <= ?", *criteria.MaxFollowers)
	}

	if criteria.MinEngagement != nil {
		query = query.Where("((likes + comments) * 100.0 / NULLIF(followers_at_post, 0)) >= ?", *criteria.MinEngagement)
	}

	if criteria.MaxEngagement != nil {
		query = query.Where("((likes + comments) * 100.0 / NULLIF(followers_at_post, 0)) <= ?", *criteria.MaxEngagement)
	}

	if criteria.OnlyActive != nil && *criteria.OnlyActive {
		query = query.Where("deleted_at IS NULL")
	}

	// Contar total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Aplicar paginação
	offset := (pagination.Page - 1) * pagination.PageSize
	query = query.Offset(offset).Limit(pagination.PageSize)

	// Ordenar por data desc
	query = query.Order("post_date DESC, created_at DESC")

	// Executar query
	if err := query.Find(&models).Error; err != nil {
		return nil, 0, err
	}

	// Converter para entities
	entities := make([]*entity.SocialPost, len(models))
	for i, m := range models {
		entities[i] = mapper.ModelToSocialPostEntity(m)
	}

	return entities, total, nil
}

func (g *SocialPostGatewayImpl) ListByBrandAndDate(brandName string, date time.Time) ([]*entity.SocialPost, error) {
	var models []*model.SocialPostModel

	query := g.db.Where("brand_name = ?", brandName).
		Where("post_date = ?", date).
		Order("post_time DESC, created_at DESC")

	if err := query.Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*entity.SocialPost, len(models))
	for i, m := range models {
		entities[i] = mapper.ModelToSocialPostEntity(m)
	}

	return entities, nil
}

func (g *SocialPostGatewayImpl) ListByBrandIDAndDate(brandID uuid.UUID, date time.Time) ([]*entity.SocialPost, error) {
	var models []*model.SocialPostModel

	query := g.db.Where("brand_id = ?", brandID).
		Where("post_date = ?", date).
		Order("post_time DESC, created_at DESC").
		Preload("Brand")

	if err := query.Find(&models).Error; err != nil {
		return nil, err
	}

	entities := make([]*entity.SocialPost, len(models))
	for i, m := range models {
		entities[i] = mapper.ModelToSocialPostEntity(m)
	}

	return entities, nil
}
