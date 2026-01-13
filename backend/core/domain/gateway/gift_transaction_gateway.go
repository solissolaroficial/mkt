package gateway

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

// GiftTransactionGateway define a interface para operações de banco de dados de GiftTransaction
type GiftTransactionGateway interface {
	// Create cria uma nova GiftTransaction
	Create(transaction *entity.GiftTransaction) error

	// FindByID busca uma GiftTransaction pelo ID
	FindByID(id uuid.UUID) (*entity.GiftTransaction, error)

	// List busca GiftTransactions baseados em critérios
	List(criteria interface{}) ([]*entity.GiftTransaction, error)

	// Update atualiza uma GiftTransaction existente
	Update(transaction *entity.GiftTransaction) error

	// Delete remove uma GiftTransaction (soft delete)
	Delete(id uuid.UUID) error

	// Count conta o número de GiftTransactions baseados em critérios
	Count(criteria interface{}) (int64, error)

	// CountByItemName conta o número de transações para um item específico
	CountByItemName(itemName string) (int64, error)
}
