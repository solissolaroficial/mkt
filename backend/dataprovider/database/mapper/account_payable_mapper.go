package mapper

import (
	"time"

	"gorm.io/gorm"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type AccountPayableMapper struct{}

func NewAccountPayableMapper() *AccountPayableMapper {
	return &AccountPayableMapper{}
}

// ModelToEntity converte Model para Entity
func (m *AccountPayableMapper) ModelToEntity(model *model.AccountPayableModel) (*entity.AccountPayable, error) {
	// Converter deletedAt
	var deletedAt *time.Time
	if model.DeletedAt.Valid {
		deletedAt = &model.DeletedAt.Time
	}

	return entity.ReconstructAccountPayable(
		model.UUID,
		model.Supplier,
		model.Description,
		model.Amount,
		model.DueDate,
		model.NFArrived,
		model.BoletoArrived,
		model.Status,
		model.Recurrence,
		model.CreatedAt,
		model.UpdatedAt,
		deletedAt,
	), nil
}

// ModelsToEntities converte slice de Model para slice de Entity
func (m *AccountPayableMapper) ModelsToEntities(models []*model.AccountPayableModel) ([]*entity.AccountPayable, error) {
	items := make([]*entity.AccountPayable, len(models))
	for i, model := range models {
		item, err := m.ModelToEntity(model)
		if err != nil {
			return nil, err
		}
		items[i] = item
	}
	return items, nil
}

// EntityToModel converte Entity para Model
func (m *AccountPayableMapper) EntityToModel(account *entity.AccountPayable) *model.AccountPayableModel {
	// Converter deletedAt
	var deletedAt gorm.DeletedAt
	if account.DeletedAt() != nil {
		deletedAt.Time = *account.DeletedAt()
		deletedAt.Valid = true
	}

	return &model.AccountPayableModel{
		UUID:          account.ID(),
		Supplier:      account.Supplier().String(),
		Description:   account.Description(),
		Amount:        account.Amount().Value(),
		DueDate:       account.DueDate().Value(),
		NFArrived:     account.NFArrived(),
		BoletoArrived: account.BoletoArrived(),
		Status:        account.Status().String(),
		Recurrence:    account.Recurrence().String(),
		CreatedAt:     account.CreatedAt(),
		UpdatedAt:     account.UpdatedAt(),
		DeletedAt:     deletedAt,
	}
}
