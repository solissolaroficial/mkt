package gifts

import (
	"github.com/google/uuid"
	domainerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// UpdateGiftTransactionInput define os dados de entrada para atualizar uma GiftTransaction
type UpdateGiftTransactionInput struct {
	ID              uuid.UUID
	ItemName        *string
	Quantity        *int
	TransactionType *string
	Date            *string
	Time            *string
	Price           *float64
	Representative  *string
	Unit            *string
}

// UpdateGiftTransactionOutput define os dados de saída após atualizar uma GiftTransaction
type UpdateGiftTransactionOutput struct {
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

// UpdateGiftTransactionUseCase define o use case para atualizar uma GiftTransaction
type UpdateGiftTransactionUseCase struct {
	giftTransactionGateway gateway.GiftTransactionGateway
}

// NewUpdateGiftTransactionUseCase cria uma nova instância de UpdateGiftTransactionUseCase
func NewUpdateGiftTransactionUseCase(giftTransactionGateway gateway.GiftTransactionGateway) *UpdateGiftTransactionUseCase {
	return &UpdateGiftTransactionUseCase{
		giftTransactionGateway: giftTransactionGateway,
	}
}

// Execute executa o use case para atualizar uma GiftTransaction
func (uc *UpdateGiftTransactionUseCase) Execute(input UpdateGiftTransactionInput) (*UpdateGiftTransactionOutput, error) {
	// Buscar a transação existente
	transaction, err := uc.giftTransactionGateway.FindByID(input.ID)
	if err != nil {
		return nil, err
	}
	if transaction == nil {
		return nil, domainerrors.ErrGiftTransactionNotFound
	}

	// Atualizar campos se fornecidos
	if input.ItemName != nil {
		if err := transaction.UpdateItemName(*input.ItemName); err != nil {
			return nil, err
		}
	}
	if input.Quantity != nil {
		if err := transaction.UpdateQuantity(*input.Quantity); err != nil {
			return nil, err
		}
	}
	if input.TransactionType != nil {
		newType := *input.TransactionType
		currentType := string(*transaction.TransactionType())

		// Se mudou o tipo, validar incompatibilidades
		if newType != currentType {
			// Transação "out" não deve ter price
			if newType == "out" && (input.Price != nil || transaction.Price() != nil) {
				return nil, domainerrors.ErrInvalidFieldForTransactionType
			}
			// Transação "in" não deve ter representative
			if newType == "in" && (input.Representative != nil || (transaction.Representative() != nil && *transaction.Representative() != "")) {
				return nil, domainerrors.ErrInvalidFieldForTransactionType
			}
		}

		if err := transaction.UpdateTransactionType(newType); err != nil {
			return nil, err
		}
	}
	if input.Price != nil {
		if err := transaction.UpdatePrice(input.Price); err != nil {
			return nil, err
		}
	}
	if input.Representative != nil {
		if err := transaction.UpdateRepresentative(input.Representative); err != nil {
			return nil, err
		}
	}
	if input.Unit != nil {
		// Unit não tem método de update na entidade, atualizamos diretamente
		// Mas precisamos de um método na entidade para isso
		// Por enquanto, assumimos que o unit não é atualizável
	}

	// Validar a entidade após as atualizações
	if err := transaction.Validate(); err != nil {
		return nil, err
	}

	// Salvar as alterações
	if err := uc.giftTransactionGateway.Update(transaction); err != nil {
		return nil, err
	}

	// Preparar output
	var price *float64
	if transaction.Price() != nil {
		p := transaction.Price().Value()
		price = &p
	}

	return &UpdateGiftTransactionOutput{
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
