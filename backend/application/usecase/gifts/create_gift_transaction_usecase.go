package gifts

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// CreateGiftTransactionInput define os dados de entrada para criar uma GiftTransaction
type CreateGiftTransactionInput struct {
	ItemName           string
	Quantity           int
	TransactionType    string
	Date               string
	Time               string
	Price              *float64
	RepresentativeUUID *string
	Unit               string
}

// CreateGiftTransactionOutput define os dados de saída após criar uma GiftTransaction
type CreateGiftTransactionOutput struct {
	ID                 uuid.UUID
	ItemName           string
	Quantity           int
	TransactionType    string
	Date               string
	Time               string
	Price              *float64
	RepresentativeUUID *string
	Unit               string
}

// CreateGiftTransactionUseCase define o use case para criar uma GiftTransaction
type CreateGiftTransactionUseCase struct {
	giftTransactionGateway gateway.GiftTransactionGateway
	giftItemGateway        gateway.GiftItemGateway
}

// NewCreateGiftTransactionUseCase cria uma nova instância de CreateGiftTransactionUseCase
func NewCreateGiftTransactionUseCase(
	giftTransactionGateway gateway.GiftTransactionGateway,
	giftItemGateway gateway.GiftItemGateway,
) *CreateGiftTransactionUseCase {
	return &CreateGiftTransactionUseCase{
		giftTransactionGateway: giftTransactionGateway,
		giftItemGateway:        giftItemGateway,
	}
}

// Execute executa o use case para criar uma GiftTransaction
func (uc *CreateGiftTransactionUseCase) Execute(input CreateGiftTransactionInput) (*CreateGiftTransactionOutput, error) {
	// Buscar o item pelo nome
	giftItem, err := uc.giftItemGateway.FindByName(input.ItemName)
	if err != nil {
		return nil, err
	}
	if giftItem == nil {
		return nil, domainerrors.ErrGiftItemNotFound
	}

	// Criar a transação
	var representativeUUID *uuid.UUID
	if input.RepresentativeUUID != nil {
		parsedUUID, err := uuid.Parse(*input.RepresentativeUUID)
		if err != nil {
			return nil, err
		}
		representativeUUID = &parsedUUID
	}

	transaction, err := entity.NewGiftTransaction(
		input.ItemName,
		input.Quantity,
		input.TransactionType,
		input.Date,
		input.Time,
		input.Price,
		representativeUUID,
		input.Unit,
	)
	if err != nil {
		return nil, err
	}

	// Salvar a transação
	if err := uc.giftTransactionGateway.Create(transaction); err != nil {
		return nil, err
	}

	// Ajustar o estoque do item
	stockAdjustment := input.Quantity
	if input.TransactionType == "out" {
		stockAdjustment = -input.Quantity
	}
	if err := uc.giftItemGateway.AdjustStock(giftItem.ID(), stockAdjustment); err != nil {
		// Rollback: deletar a transação recém-criada se o ajuste de estoque falhar
		_ = uc.giftTransactionGateway.Delete(transaction.ID())
		return nil, err
	}

	// Preparar output
	var price *float64
	if transaction.Price() != nil {
		p := transaction.Price().Value()
		price = &p
	}

	// Converter UUID para string no output
	var representativeUUIDStr *string
	if transaction.RepresentativeUUID() != (uuid.UUID{}) {
		uuidStr := transaction.RepresentativeUUID().String()
		representativeUUIDStr = &uuidStr
	}

	return &CreateGiftTransactionOutput{
		ID:                 transaction.ID(),
		ItemName:           transaction.ItemName(),
		Quantity:           transaction.Quantity().Value(),
		TransactionType:    string(*transaction.TransactionType()),
		Date:               transaction.Date().Format("02/01/2006"),
		Time:               transaction.Time(),
		Price:              price,
		RepresentativeUUID: representativeUUIDStr,
		Unit:               transaction.Unit(),
	}, nil
}
