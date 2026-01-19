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

type giftTransactionGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.GiftTransactionMapper
}

func NewGiftTransactionGateway(db *gorm.DB) gateway.GiftTransactionGateway {
	return &giftTransactionGatewayImpl{
		db:     db,
		mapper: mapper.NewGiftTransactionMapper(),
	}
}

func (g *giftTransactionGatewayImpl) Create(transaction *entity.GiftTransaction) error {
	transactionModel := g.mapper.EntityToModel(transaction)
	return g.db.Create(transactionModel).Error
}

func (g *giftTransactionGatewayImpl) FindByID(id uuid.UUID) (*entity.GiftTransaction, error) {
	var transactionModel model.GiftTransactionModel
	err := g.db.Where("uuid = ?", id).First(&transactionModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainerrors.ErrGiftTransactionNotFound
		}
		return nil, err
	}
	return g.mapper.ModelToEntity(&transactionModel)
}

func (g *giftTransactionGatewayImpl) List(criteria interface{}) ([]*entity.GiftTransaction, error) {
	var transactionModels []model.GiftTransactionModel

	query := g.db.Model(&transactionModels)

	// Aplicar criteria se fornecido
	if crit, ok := criteria.(*domain.GiftTransactionCriteria); ok {
		if crit.ItemName() != nil {
			// Verificar se o item existe antes de filtrar por nome
			var itemModel model.GiftItemModel
			err := g.db.Where("name = ? AND deleted_at IS NULL", *crit.ItemName()).First(&itemModel).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					// Item não existe, retornar lista vazia
					return []*entity.GiftTransaction{}, nil
				}
				return nil, err
			}
			query = query.Where("item_name = ?", *crit.ItemName())
		}
		if crit.TransactionType() != nil {
			query = query.Where("transaction_type = ?", string(*crit.TransactionType()))
		}
		if crit.RepresentativeUUID() != nil {
			query = query.Where("representative_uuid = ?", *crit.RepresentativeUUID())
		}
		if crit.StartDate() != nil {
			query = query.Where("date >= ?", *crit.StartDate())
		}
		if crit.EndDate() != nil {
			query = query.Where("date <= ?", *crit.EndDate())
		}
		// Paginação
		if crit.Limit() != nil && *crit.Limit() > 0 {
			query = query.Offset(crit.GetOffset()).Limit(*crit.Limit())
		}
	}

	// Default: apenas ativos
	query = query.Where("deleted_at IS NULL")

	// Ordenar por data e hora decrescente (mais recentes primeiro)
	query = query.Order("date DESC, time DESC")

	if err := query.Find(&transactionModels).Error; err != nil {
		return nil, err
	}

	// Convert slice of models to slice of pointers
	transactionModelPointers := make([]*model.GiftTransactionModel, len(transactionModels))
	for i := range transactionModels {
		transactionModelPointers[i] = &transactionModels[i]
	}

	return g.mapper.ModelsToEntities(transactionModelPointers)
}

func (g *giftTransactionGatewayImpl) Update(transaction *entity.GiftTransaction) error {
	transactionModel := g.mapper.EntityToModel(transaction)
	result := g.db.Where("uuid = ?", transactionModel.UUID).Save(transactionModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainerrors.ErrGiftTransactionNotFound
	}
	return nil
}

func (g *giftTransactionGatewayImpl) Delete(id uuid.UUID) error {
	result := g.db.Delete(&model.GiftTransactionModel{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainerrors.ErrGiftTransactionNotFound
	}
	return nil
}

func (g *giftTransactionGatewayImpl) Count(criteria interface{}) (int64, error) {
	var count int64

	query := g.db.Model(&model.GiftTransactionModel{})

	// Aplicar criteria se fornecido
	if crit, ok := criteria.(*domain.GiftTransactionCriteria); ok {
		if crit.ItemName() != nil {
			query = query.Where("item_name = ?", *crit.ItemName())
		}
		if crit.TransactionType() != nil {
			query = query.Where("transaction_type = ?", string(*crit.TransactionType()))
		}
		if crit.RepresentativeUUID() != nil {
			query = query.Where("representative_uuid = ?", *crit.RepresentativeUUID())
		}
		if crit.StartDate() != nil {
			query = query.Where("date >= ?", *crit.StartDate())
		}
		if crit.EndDate() != nil {
			query = query.Where("date <= ?", *crit.EndDate())
		}
	}

	// Default: apenas ativos
	query = query.Where("deleted_at IS NULL")

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (g *giftTransactionGatewayImpl) CountByItemName(itemName string) (int64, error) {
	var count int64
	err := g.db.Model(&model.GiftTransactionModel{}).
		Where("item_name = ? AND deleted_at IS NULL", itemName).
		Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
