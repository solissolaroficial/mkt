package gateway

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type giftItemGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.GiftItemMapper
}

func NewGiftItemGateway(db *gorm.DB) gateway.GiftItemGateway {
	return &giftItemGatewayImpl{
		db:     db,
		mapper: mapper.NewGiftItemMapper(),
	}
}

func (g *giftItemGatewayImpl) Create(item *entity.GiftItem) error {
	itemModel := g.mapper.EntityToModel(item)
	return g.db.Create(itemModel).Error
}

func (g *giftItemGatewayImpl) FindByID(id uuid.UUID) (*entity.GiftItem, error) {
	var itemModel model.GiftItemModel
	err := g.db.Where("uuid = ?", id).First(&itemModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrGiftItemNotFound
		}
		return nil, err
	}
	return g.mapper.ModelToEntity(&itemModel)
}

func (g *giftItemGatewayImpl) FindByName(name string) (*entity.GiftItem, error) {
	var itemModel model.GiftItemModel
	err := g.db.Where("name = ? AND deleted_at IS NULL", name).First(&itemModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrGiftItemNotFound
		}
		return nil, err
	}
	return g.mapper.ModelToEntity(&itemModel)
}

func (g *giftItemGatewayImpl) ExistsByName(name string) (bool, error) {
	var count int64
	err := g.db.Model(&model.GiftItemModel{}).
		Where("name = ? AND deleted_at IS NULL", name).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (g *giftItemGatewayImpl) List(criteria interface{}) ([]*entity.GiftItem, error) {
	var itemModels []model.GiftItemModel

	query := g.db.Model(&itemModels)

	// Aplicar criteria se fornecido
	if crit, ok := criteria.(*domain.GiftItemCriteria); ok {
		if crit.Name() != nil {
			query = query.Where("name = ?", crit.Name().Value())
		}
		if crit.MinStock() != nil {
			query = query.Where("stock >= ?", *crit.MinStock())
		}
		if crit.MaxStock() != nil {
			query = query.Where("stock <= ?", *crit.MaxStock())
		}
		if crit.MinPrice() != nil {
			query = query.Where("price >= ?", *crit.MinPrice())
		}
		if crit.MaxPrice() != nil {
			query = query.Where("price <= ?", *crit.MaxPrice())
		}
		// Paginação
		if crit.Limit() != nil && *crit.Limit() > 0 {
			query = query.Offset(crit.GetOffset()).Limit(*crit.Limit())
		}
	}

	// Default: apenas ativos
	query = query.Where("deleted_at IS NULL")

	if err := query.Find(&itemModels).Error; err != nil {
		return nil, err
	}

	// Convert slice of models to slice of pointers
	itemModelPointers := make([]*model.GiftItemModel, len(itemModels))
	for i := range itemModels {
		itemModelPointers[i] = &itemModels[i]
	}

	return g.mapper.ModelsToEntities(itemModelPointers)
}

func (g *giftItemGatewayImpl) Update(item *entity.GiftItem) error {
	itemModel := g.mapper.EntityToModel(item)
	result := g.db.Where("uuid = ?", itemModel.UUID).Save(itemModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainerrors.ErrGiftItemNotFound
	}
	return nil
}

func (g *giftItemGatewayImpl) Delete(id uuid.UUID) error {
	result := g.db.Delete(&model.GiftItemModel{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainerrors.ErrGiftItemNotFound
	}
	return nil
}

func (g *giftItemGatewayImpl) AdjustStock(id uuid.UUID, quantity int) error {
	result := g.db.Model(&model.GiftItemModel{}).
		Where("uuid = ?", id).
		Update("stock", gorm.Expr("stock + ?", quantity))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainerrors.ErrGiftItemNotFound
	}
	return nil
}

func (g *giftItemGatewayImpl) Count(criteria interface{}) (int64, error) {
	var count int64

	query := g.db.Model(&model.GiftItemModel{})

	// Aplicar criteria se fornecido
	if crit, ok := criteria.(*domain.GiftItemCriteria); ok {
		if crit.Name() != nil {
			query = query.Where("name = ?", crit.Name().Value())
		}
		if crit.MinStock() != nil {
			query = query.Where("stock >= ?", *crit.MinStock())
		}
		if crit.MaxStock() != nil {
			query = query.Where("stock <= ?", *crit.MaxStock())
		}
		if crit.MinPrice() != nil {
			query = query.Where("price >= ?", *crit.MinPrice())
		}
		if crit.MaxPrice() != nil {
			query = query.Where("price <= ?", *crit.MaxPrice())
		}
	}

	// Default: apenas ativos
	query = query.Where("deleted_at IS NULL")

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
