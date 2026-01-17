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

// recurrentPdvGatewayImpl implementa a interface RecurrentPdvGateway usando GORM
type recurrentPdvGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.RecurrentPdvMapper
}

// NewRecurrentPdvGateway cria uma nova instância do RecurrentPdvGateway
func NewRecurrentPdvGateway(db *gorm.DB) gateway.RecurrentPdvGateway {
	return &recurrentPdvGatewayImpl{
		db:     db,
		mapper: mapper.NewRecurrentPdvMapper(),
	}
}

// Save salva um novo PDV recorrente no banco de dados
func (g *recurrentPdvGatewayImpl) Save(ctx context.Context, pdv *entity.RecurrentPdv) error {
	pdvModel := g.mapper.EntityToModel(pdv)
	return g.db.WithContext(ctx).Create(pdvModel).Error
}

// FindByID busca um PDV recorrente por ID
func (g *recurrentPdvGatewayImpl) FindByID(ctx context.Context, id string) (*entity.RecurrentPdv, error) {
	var pdvModel model.RecurrentPdvModel
	err := g.db.WithContext(ctx).Where("uuid = ?", id).First(&pdvModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrRecurrentPdvNotFound
		}
		return nil, err
	}
	return g.mapper.ModelToEntity(&pdvModel)
}

// Update atualiza um PDV recorrente existente
func (g *recurrentPdvGatewayImpl) Update(ctx context.Context, pdv *entity.RecurrentPdv) error {
	pdvModel := g.mapper.EntityToModel(pdv)
	result := g.db.WithContext(ctx).Where("uuid = ?", pdvModel.UUID).Save(pdvModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrRecurrentPdvNotFound
	}
	return nil
}

// Delete realiza soft delete de um PDV recorrente
func (g *recurrentPdvGatewayImpl) Delete(ctx context.Context, id string) error {
	result := g.db.WithContext(ctx).Delete(&model.RecurrentPdvModel{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrRecurrentPdvNotFound
	}
	return nil
}

// FindByCriteria busca PDVs recorrentes baseado em critérios
func (g *recurrentPdvGatewayImpl) FindByCriteria(
	ctx context.Context,
	criteria *domain.RecurrentPdvCriteria,
	pagination *valueobject.Pagination,
	sortOrder *valueobject.SortOrder,
) ([]*entity.RecurrentPdv, error) {
	var pdvModels []model.RecurrentPdvModel

	query := g.db.WithContext(ctx).Model(&pdvModels).Where("deleted_at IS NULL")

	// Aplicar criteria usando getters
	if criteria.RepresentativeUUID() != nil {
		query = query.Where("representative_uuid = ?", *criteria.RepresentativeUUID())
	}
	if criteria.City() != nil {
		query = query.Where("city = ?", *criteria.City())
	}

	// Aplicar ordenação
	orderBy := "name ASC"
	if sortOrder != nil && sortOrder.GetField() == "name" {
		if sortOrder.GetDirection() == "DESC" {
			orderBy = "name DESC"
		}
	}
	query = query.Order(orderBy)

	// Aplicar paginação
	if pagination != nil {
		offset := pagination.Offset()
		limit := pagination.PageSize
		query = query.Offset(offset).Limit(limit)
	}

	if err := query.Find(&pdvModels).Error; err != nil {
		return nil, err
	}

	// Convert slice of models to slice of pointers
	pdvModelPointers := make([]*model.RecurrentPdvModel, len(pdvModels))
	for i := range pdvModels {
		pdvModelPointers[i] = &pdvModels[i]
	}

	return g.mapper.ModelsToEntities(pdvModelPointers)
}

// CountByCriteria conta PDVs recorrentes baseado em critérios
func (g *recurrentPdvGatewayImpl) CountByCriteria(ctx context.Context, criteria *domain.RecurrentPdvCriteria) (int64, error) {
	var count int64

	query := g.db.WithContext(ctx).Model(&model.RecurrentPdvModel{}).Where("deleted_at IS NULL")

	// Aplicar criteria usando getters
	if criteria.RepresentativeUUID() != nil {
		query = query.Where("representative_uuid = ?", *criteria.RepresentativeUUID())
	}
	if criteria.City() != nil {
		query = query.Where("city = ?", *criteria.City())
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// ExistsByID verifica se um PDV recorrente existe por ID
func (g *recurrentPdvGatewayImpl) ExistsByID(ctx context.Context, id string) (bool, error) {
	var count int64
	err := g.db.WithContext(ctx).
		Model(&model.RecurrentPdvModel{}).
		Where("uuid = ? AND deleted_at IS NULL", id).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
