package gifts

import (
	"github.com/google/uuid"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetGiftTransactionInput define os dados de entrada para buscar uma GiftTransaction
type GetGiftTransactionInput struct {
	ID uuid.UUID
}

// GetGiftTransactionOutput define os dados de saída após buscar uma GiftTransaction
type GetGiftTransactionOutput struct {
	ID              uuid.UUID
	ItemName        string
	Quantity        int
	TransactionType string
	Date            string
	Time            string
	Price           *float64
	Representative  *string
	Unit            string
	CreatedAt       string
	UpdatedAt       string
}

// GetGiftTransactionUseCase define o use case para buscar uma GiftTransaction
type GetGiftTransactionUseCase struct {
	giftTransactionGateway gateway.GiftTransactionGateway
}

// NewGetGiftTransactionUseCase cria uma nova instância de GetGiftTransactionUseCase
func NewGetGiftTransactionUseCase(giftTransactionGateway gateway.GiftTransactionGateway) *GetGiftTransactionUseCase {
	return &GetGiftTransactionUseCase{
		giftTransactionGateway: giftTransactionGateway,
	}
}

// Execute executa o use case para buscar uma GiftTransaction
func (uc *GetGiftTransactionUseCase) Execute(input GetGiftTransactionInput) (*GetGiftTransactionOutput, error) {
	transaction, err := uc.giftTransactionGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	if transaction == nil {
		return nil, domainerrors.ErrGiftTransactionNotFound
	}

	// Preparar output
	var price *float64
	if transaction.Price() != nil {
		p := transaction.Price().Value()
		price = &p
	}

	return &GetGiftTransactionOutput{
		ID:              transaction.ID(),
		ItemName:        transaction.ItemName(),
		Quantity:        transaction.Quantity().Value(),
		TransactionType: string(*transaction.TransactionType()),
		Date:            transaction.Date().Format("02/01/2006"),
		Time:            transaction.Time(),
		Price:           price,
		Representative:  transaction.Representative(),
		Unit:            transaction.Unit(),
		CreatedAt:       transaction.CreatedAt().Format("2006-01-02 15:04:05"),
		UpdatedAt:       transaction.UpdatedAt().Format("2006-01-02 15:04:05"),
	}, nil
}
