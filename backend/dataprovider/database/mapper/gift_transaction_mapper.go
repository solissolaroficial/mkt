package mapper

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type GiftTransactionMapper struct{}

func NewGiftTransactionMapper() *GiftTransactionMapper {
	return &GiftTransactionMapper{}
}

// ModelToEntity converte Model para Entity
func (m *GiftTransactionMapper) ModelToEntity(model *model.GiftTransactionModel) (*entity.GiftTransaction, error) {
	// Converter deletedAt
	var deletedAt *time.Time
	if model.DeletedAt.Valid {
		deletedAt = &model.DeletedAt.Time
	}

	return entity.ReconstructGiftTransaction(
		model.UUID,
		model.ItemName,
		model.Quantity,
		model.TransactionType,
		model.Date,
		model.Time,
		model.Price,
		model.Representative,
		model.Unit,
		model.CreatedAt,
		model.UpdatedAt,
		deletedAt,
	), nil
}

// ModelsToEntities converte slice de Model para slice de Entity
func (m *GiftTransactionMapper) ModelsToEntities(models []*model.GiftTransactionModel) ([]*entity.GiftTransaction, error) {
	transactions := make([]*entity.GiftTransaction, len(models))
	for i, model := range models {
		transaction, err := m.ModelToEntity(model)
		if err != nil {
			return nil, err
		}
		transactions[i] = transaction
	}
	return transactions, nil
}

// EntityToModel converte Entity para Model
func (m *GiftTransactionMapper) EntityToModel(transaction *entity.GiftTransaction) *model.GiftTransactionModel {
	// Converter deletedAt
	var deletedAt gorm.DeletedAt
	if transaction.DeletedAt() != nil {
		deletedAt.Time = *transaction.DeletedAt()
		deletedAt.Valid = true
	}

	// Converter price
	var price *float64
	if transaction.Price() != nil {
		p := transaction.Price().Value()
		price = &p
	}

	return &model.GiftTransactionModel{
		UUID:            transaction.ID(),
		ItemName:        transaction.ItemName(),
		Quantity:        transaction.Quantity().Value(),
		TransactionType: string(*transaction.TransactionType()),
		Date:            transaction.Date(),
		Time:            transaction.Time(),
		Price:           price,
		Representative:  transaction.Representative(),
		Unit:            transaction.Unit(),
		CreatedAt:       transaction.CreatedAt(),
		UpdatedAt:       transaction.UpdatedAt(),
		DeletedAt:       deletedAt,
	}
}
