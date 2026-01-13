package gateway

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

// GiftItemGateway define a interface para operações de banco de dados de GiftItem
type GiftItemGateway interface {
	// Create cria um novo GiftItem
	Create(item *entity.GiftItem) error

	// FindByID busca um GiftItem pelo ID
	FindByID(id uuid.UUID) (*entity.GiftItem, error)

	// FindByName busca um GiftItem pelo nome (adicionado para CreateGiftTransactionUseCase)
	FindByName(name string) (*entity.GiftItem, error)

	// ExistsByName verifica se existe um GiftItem com o nome fornecido
	ExistsByName(name string) (bool, error)

	// List busca GiftItems baseados em critérios
	List(criteria interface{}) ([]*entity.GiftItem, error)

	// Update atualiza um GiftItem existente
	Update(item *entity.GiftItem) error

	// Delete remove um GiftItem (soft delete)
	Delete(id uuid.UUID) error

	// AdjustStock ajusta o estoque de um GiftItem
	AdjustStock(id uuid.UUID, quantity int) error

	// Count conta o número de GiftItems baseados em critérios
	Count(criteria interface{}) (int64, error)
}
