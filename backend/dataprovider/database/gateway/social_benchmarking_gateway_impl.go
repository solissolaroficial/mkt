package gateway

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type socialBenchmarkingGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.SocialBenchmarkingMapper
}

func NewSocialBenchmarkingGateway(db *gorm.DB) gateway.SocialBenchmarkingGateway {
	return &socialBenchmarkingGatewayImpl{
		db:     db,
		mapper: mapper.NewSocialBenchmarkingMapper(),
	}
}

func (g *socialBenchmarkingGatewayImpl) Save(ctx context.Context, benchmarking *entity.SocialBenchmarking) error {
	benchmarkingModel := g.mapper.EntityToModel(benchmarking)
	return g.db.WithContext(ctx).Create(benchmarkingModel).Error
}

func (g *socialBenchmarkingGatewayImpl) FindByID(ctx context.Context, id string) (*entity.SocialBenchmarking, error) {
	var benchmarkingModel model.SocialBenchmarkingModel
	err := g.db.WithContext(ctx).Where("uuid = ?", id).First(&benchmarkingModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrSocialBenchmarkingNotFound
		}
		return nil, err
	}
	return g.mapper.ModelToEntity(&benchmarkingModel)
}

func (g *socialBenchmarkingGatewayImpl) Update(ctx context.Context, benchmarking *entity.SocialBenchmarking) error {
	benchmarkingModel := g.mapper.EntityToModel(benchmarking)
	result := g.db.WithContext(ctx).Where("uuid = ?", benchmarkingModel.UUID).Save(benchmarkingModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrSocialBenchmarkingNotFound
	}
	return nil
}

func (g *socialBenchmarkingGatewayImpl) Delete(ctx context.Context, id string) error {
	result := g.db.WithContext(ctx).Delete(&model.SocialBenchmarkingModel{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrSocialBenchmarkingNotFound
	}
	return nil
}

func (g *socialBenchmarkingGatewayImpl) FindByCriteria(
	ctx context.Context,
	criteria *domain.SocialBenchmarkingCriteria,
	pagination *valueobject.Pagination,
	sortOrder *valueobject.SortOrder,
) ([]*entity.SocialBenchmarking, error) {
	var benchmarkingModels []model.SocialBenchmarkingModel

	query := g.db.WithContext(ctx).Model(&benchmarkingModels)

	// Aplicar criteria usando getters (sem dependência de GORM no domínio)
	if criteria.BrandID() != nil {
		query = query.Where("brand_id = ?", criteria.BrandID())
	}
	if criteria.Active() != nil {
		if *criteria.Active() {
			query = query.Where("deleted_at IS NULL")
		} else {
			query = query.Where("deleted_at IS NOT NULL")
		}
	} else {
		// Default: apenas ativos
		query = query.Where("deleted_at IS NULL")
	}
	if criteria.StartDate() != nil {
		query = query.Where("created_at >= ?", *criteria.StartDate())
	}
	if criteria.EndDate() != nil {
		query = query.Where("created_at <= ?", *criteria.EndDate())
	}

	// Aplicar ordenação
	if sortOrder != nil {
		query = query.Order(sortOrder.ToSQLString())
	} else {
		// Default: ordenar por engagement_rate DESC
		query = query.Order("engagement_rate DESC")
	}

	// Aplicar paginação
	if pagination != nil {
		offset := pagination.Offset()
		limit := pagination.PageSize
		query = query.Offset(offset).Limit(limit)
	}

	if err := query.Find(&benchmarkingModels).Error; err != nil {
		return nil, err
	}

	// Convert slice of models to slice of pointers
	benchmarkingModelPointers := make([]*model.SocialBenchmarkingModel, len(benchmarkingModels))
	for i := range benchmarkingModels {
		benchmarkingModelPointers[i] = &benchmarkingModels[i]
	}

	return g.mapper.ModelsToEntities(benchmarkingModelPointers)
}

func (g *socialBenchmarkingGatewayImpl) CountByCriteria(ctx context.Context, criteria *domain.SocialBenchmarkingCriteria) (int64, error) {
	var count int64

	query := g.db.WithContext(ctx).Model(&model.SocialBenchmarkingModel{})

	// Aplicar criteria usando getters
	if criteria.BrandID() != nil {
		query = query.Where("brand_id = ?", criteria.BrandID())
	}
	if criteria.Active() != nil {
		if *criteria.Active() {
			query = query.Where("deleted_at IS NULL")
		} else {
			query = query.Where("deleted_at IS NOT NULL")
		}
	} else {
		// Default: apenas ativos
		query = query.Where("deleted_at IS NULL")
	}
	if criteria.StartDate() != nil {
		query = query.Where("created_at >= ?", *criteria.StartDate())
	}
	if criteria.EndDate() != nil {
		query = query.Where("created_at <= ?", *criteria.EndDate())
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (g *socialBenchmarkingGatewayImpl) ExistsByID(ctx context.Context, id string) (bool, error) {
	var count int64
	err := g.db.WithContext(ctx).
		Model(&model.SocialBenchmarkingModel{}).
		Where("uuid = ? AND deleted_at IS NULL", id).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (g *socialBenchmarkingGatewayImpl) GetByBrandID(brandID uuid.UUID) (*entity.SocialBenchmarking, error) {
	var benchmarkingModel model.SocialBenchmarkingModel
	err := g.db.Where("brand_id = ? AND deleted_at IS NULL", brandID).First(&benchmarkingModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrSocialBenchmarkingNotFound
		}
		return nil, err
	}
	return g.mapper.ModelToEntity(&benchmarkingModel)
}
