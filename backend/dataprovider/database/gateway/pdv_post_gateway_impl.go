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

// pdvPostGatewayImpl implementa a interface PdvPostGateway usando GORM
type pdvPostGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.PdvPostMapper
}

// NewPdvPostGateway cria uma nova instância do PdvPostGateway
func NewPdvPostGateway(db *gorm.DB) gateway.PdvPostGateway {
	return &pdvPostGatewayImpl{
		db:     db,
		mapper: mapper.NewPdvPostMapper(),
	}
}

// Save salva um novo post de PDV no banco de dados
func (g *pdvPostGatewayImpl) Save(ctx context.Context, post *entity.PdvPost) error {
	postModel := g.mapper.EntityToModel(post)
	return g.db.WithContext(ctx).Create(postModel).Error
}

// FindByID busca um post de PDV por ID
func (g *pdvPostGatewayImpl) FindByID(ctx context.Context, id string) (*entity.PdvPost, error) {
	var postModel model.PdvPostModel
	err := g.db.WithContext(ctx).Where("uuid = ?", id).First(&postModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrPdvPostNotFound
		}
		return nil, err
	}
	return g.mapper.ModelToEntity(&postModel)
}

// Update atualiza um post de PDV existente
func (g *pdvPostGatewayImpl) Update(ctx context.Context, post *entity.PdvPost) error {
	postModel := g.mapper.EntityToModel(post)
	result := g.db.WithContext(ctx).Where("uuid = ?", postModel.UUID).Save(postModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrPdvPostNotFound
	}
	return nil
}

// Delete realiza soft delete de um post de PDV
func (g *pdvPostGatewayImpl) Delete(ctx context.Context, id string) error {
	result := g.db.WithContext(ctx).Delete(&model.PdvPostModel{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrPdvPostNotFound
	}
	return nil
}

// FindByCriteria busca posts de PDV baseado em critérios
func (g *pdvPostGatewayImpl) FindByCriteria(
	ctx context.Context,
	criteria *domain.PdvPostCriteria,
	pagination *valueobject.Pagination,
	sortOrder *valueobject.SortOrder,
) ([]*entity.PdvPost, error) {
	var postModels []model.PdvPostModel

	query := g.db.WithContext(ctx).Model(&postModels).Where("deleted_at IS NULL")

	// Aplicar criteria usando getters (sem dependência de GORM no domínio)
	if criteria.RepresentativeUUID() != nil {
		query = query.Where("representative_uuid = ?", *criteria.RepresentativeUUID())
	}
	if criteria.Month() != nil {
		query = query.Where("month = ?", *criteria.Month())
	}
	if criteria.Platform() != nil {
		query = query.Where("platform = ?", *criteria.Platform())
	}
	if criteria.Status() != nil {
		query = query.Where("status = ?", criteria.Status().String())
	}
	if criteria.StartDate() != nil {
		query = query.Where("post_date >= ?", *criteria.StartDate())
	}
	if criteria.EndDate() != nil {
		query = query.Where("post_date <= ?", *criteria.EndDate())
	}

	// Aplicar ordenação
	orderBy := "post_date DESC"
	if sortOrder != nil && sortOrder.GetField() == "post_date" {
		if sortOrder.GetDirection() == "ASC" {
			orderBy = "post_date ASC"
		}
	}
	query = query.Order(orderBy)

	// Aplicar paginação
	if pagination != nil {
		offset := pagination.Offset()
		limit := pagination.PageSize
		query = query.Offset(offset).Limit(limit)
	}

	if err := query.Find(&postModels).Error; err != nil {
		return nil, err
	}

	// Convert slice of models to slice of pointers
	postModelPointers := make([]*model.PdvPostModel, len(postModels))
	for i := range postModels {
		postModelPointers[i] = &postModels[i]
	}

	return g.mapper.ModelsToEntities(postModelPointers)
}

// CountByCriteria conta posts de PDV baseado em critérios
func (g *pdvPostGatewayImpl) CountByCriteria(ctx context.Context, criteria *domain.PdvPostCriteria) (int64, error) {
	var count int64

	query := g.db.WithContext(ctx).Model(&model.PdvPostModel{}).Where("deleted_at IS NULL")

	// Aplicar criteria usando getters
	if criteria.RepresentativeUUID() != nil {
		query = query.Where("representative_uuid = ?", *criteria.RepresentativeUUID())
	}
	if criteria.Month() != nil {
		query = query.Where("month = ?", *criteria.Month())
	}
	if criteria.Platform() != nil {
		query = query.Where("platform = ?", *criteria.Platform())
	}
	if criteria.Status() != nil {
		query = query.Where("status = ?", criteria.Status().String())
	}
	if criteria.StartDate() != nil {
		query = query.Where("post_date >= ?", *criteria.StartDate())
	}
	if criteria.EndDate() != nil {
		query = query.Where("post_date <= ?", *criteria.EndDate())
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// ExistsByID verifica se um post de PDV existe por ID
func (g *pdvPostGatewayImpl) ExistsByID(ctx context.Context, id string) (bool, error) {
	var count int64
	err := g.db.WithContext(ctx).
		Model(&model.PdvPostModel{}).
		Where("uuid = ? AND deleted_at IS NULL", id).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
