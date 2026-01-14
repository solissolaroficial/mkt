package gateway

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainErrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/dataprovider/database/mapper"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type accountPayableGatewayImpl struct {
	db     *gorm.DB
	mapper *mapper.AccountPayableMapper
}

func NewAccountPayableGateway(db *gorm.DB) gateway.AccountPayableGateway {
	return &accountPayableGatewayImpl{
		db:     db,
		mapper: mapper.NewAccountPayableMapper(),
	}
}

func (g *accountPayableGatewayImpl) Create(account *entity.AccountPayable) error {
	accountModel := g.mapper.EntityToModel(account)
	return g.db.Create(accountModel).Error
}

func (g *accountPayableGatewayImpl) FindByID(id uuid.UUID) (*entity.AccountPayable, error) {
	var accountModel model.AccountPayableModel
	err := g.db.Where("uuid = ?", id).First(&accountModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domainErrors.ErrAccountPayableNotFound
		}
		return nil, err
	}
	return g.mapper.ModelToEntity(&accountModel)
}

func (g *accountPayableGatewayImpl) List(criteria interface{}) ([]*entity.AccountPayable, error) {
	var accountModels []model.AccountPayableModel

	query := g.db.Model(&accountModels)

	// Aplicar criteria se fornecido
	if crit, ok := criteria.(*domain.AccountPayableCriteria); ok {
		if crit.Supplier() != nil {
			query = query.Where("supplier = ?", crit.Supplier().String())
		}
		if crit.MinAmount() != nil {
			query = query.Where("amount >= ?", *crit.MinAmount())
		}
		if crit.MaxAmount() != nil {
			query = query.Where("amount <= ?", *crit.MaxAmount())
		}
		if crit.Status() != nil {
			query = query.Where("status = ?", crit.Status().String())
		}
		if crit.Recurrence() != nil {
			query = query.Where("recurrence = ?", crit.Recurrence().String())
		}
		if crit.StartDate() != nil {
			query = query.Where("due_date >= ?", *crit.StartDate())
		}
		if crit.EndDate() != nil {
			query = query.Where("due_date <= ?", *crit.EndDate())
		}
		// Paginação
		if crit.Limit() != nil && *crit.Limit() > 0 {
			query = query.Offset(crit.GetOffset()).Limit(*crit.Limit())
		}

		// Ordenação
		if crit.SortBy() != nil && crit.SortOrder() != nil {
			orderClause := fmt.Sprintf("%s %s", *crit.SortBy(), *crit.SortOrder())
			query = query.Order(orderClause)
		} else {
			// Ordenação padrão
			query = query.Order("due_date ASC")
		}
	}

	// Default: apenas ativos
	query = query.Where("deleted_at IS NULL")

	if err := query.Find(&accountModels).Error; err != nil {
		return nil, err
	}

	// Convert slice of models to slice of pointers
	accountModelPointers := make([]*model.AccountPayableModel, len(accountModels))
	for i := range accountModels {
		accountModelPointers[i] = &accountModels[i]
	}

	return g.mapper.ModelsToEntities(accountModelPointers)
}

func (g *accountPayableGatewayImpl) Update(account *entity.AccountPayable) error {
	accountModel := g.mapper.EntityToModel(account)
	result := g.db.Where("uuid = ?", accountModel.UUID).Save(accountModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrAccountPayableNotFound
	}
	return nil
}

func (g *accountPayableGatewayImpl) Delete(id uuid.UUID) error {
	result := g.db.Delete(&model.AccountPayableModel{}, "uuid = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domainErrors.ErrAccountPayableNotFound
	}
	return nil
}

func (g *accountPayableGatewayImpl) Count(criteria interface{}) (int64, error) {
	var count int64

	query := g.db.Model(&model.AccountPayableModel{})

	// Aplicar criteria se fornecido
	if crit, ok := criteria.(*domain.AccountPayableCriteria); ok {
		if crit.Supplier() != nil {
			query = query.Where("supplier = ?", crit.Supplier().String())
		}
		if crit.MinAmount() != nil {
			query = query.Where("amount >= ?", *crit.MinAmount())
		}
		if crit.MaxAmount() != nil {
			query = query.Where("amount <= ?", *crit.MaxAmount())
		}
		if crit.Status() != nil {
			query = query.Where("status = ?", crit.Status().String())
		}
		if crit.Recurrence() != nil {
			query = query.Where("recurrence = ?", crit.Recurrence().String())
		}
		if crit.StartDate() != nil {
			query = query.Where("due_date >= ?", *crit.StartDate())
		}
		if crit.EndDate() != nil {
			query = query.Where("due_date <= ?", *crit.EndDate())
		}
	}

	// Default: apenas ativos
	query = query.Where("deleted_at IS NULL")

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (g *accountPayableGatewayImpl) ExistsBySupplierAndDueDate(supplier string, dueDate time.Time) (bool, error) {
	var count int64
	err := g.db.Model(&model.AccountPayableModel{}).
		Where("supplier = ? AND due_date = ? AND deleted_at IS NULL", supplier, dueDate).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
