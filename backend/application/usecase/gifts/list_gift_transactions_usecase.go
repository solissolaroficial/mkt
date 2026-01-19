package gifts

import (
	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// ListGiftTransactionsInput define os dados de entrada para listar GiftTransactions
type ListGiftTransactionsInput struct {
	ItemName           *string
	TransactionType    *string
	RepresentativeUUID *string
	StartDate          *string
	EndDate            *string
	Page               *int
	Limit              *int
}

// ListGiftTransactionsOutput define os dados de saída após listar GiftTransactions
type ListGiftTransactionsOutput struct {
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

// ListGiftTransactionsUseCase define o use case para listar GiftTransactions
type ListGiftTransactionsUseCase struct {
	giftTransactionGateway gateway.GiftTransactionGateway
}

// NewListGiftTransactionsUseCase cria uma nova instância de ListGiftTransactionsUseCase
func NewListGiftTransactionsUseCase(giftTransactionGateway gateway.GiftTransactionGateway) *ListGiftTransactionsUseCase {
	return &ListGiftTransactionsUseCase{
		giftTransactionGateway: giftTransactionGateway,
	}
}

// Execute executa o use case para listar GiftTransactions
func (uc *ListGiftTransactionsUseCase) Execute(input ListGiftTransactionsInput) ([]*ListGiftTransactionsOutput, error) {
	// Criar critérios de busca
	criteria := domain.NewGiftTransactionCriteria()

	if input.ItemName != nil {
		criteria.WithItemName(input.ItemName)
	}
	if input.TransactionType != nil {
		txType, err := valueobject.NewTransactionType(*input.TransactionType)
		if err == nil {
			criteria.WithTransactionType(&txType)
		}
	}
	if input.RepresentativeUUID != nil {
		criteria.WithRepresentativeUUID(input.RepresentativeUUID)
	}
	if input.StartDate != nil {
		criteria.WithStartDate(input.StartDate)
	}
	if input.EndDate != nil {
		criteria.WithEndDate(input.EndDate)
	}
	if input.Page != nil {
		criteria.WithPage(input.Page)
	}
	if input.Limit != nil {
		criteria.WithLimit(input.Limit)
	}

	// Buscar transações
	transactions, err := uc.giftTransactionGateway.List(criteria)
	if err != nil {
		return nil, err
	}

	// Converter para output
	output := make([]*ListGiftTransactionsOutput, len(transactions))
	for i, tx := range transactions {
		var price *float64
		if tx.Price() != nil {
			p := tx.Price().Value()
			price = &p
		}

		// Converter UUID para string no output
		var representativeUUIDStr *string
		if tx.RepresentativeUUID() != (uuid.UUID{}) {
			uuidStr := tx.RepresentativeUUID().String()
			representativeUUIDStr = &uuidStr
		}

		output[i] = &ListGiftTransactionsOutput{
			ID:                 tx.ID(),
			ItemName:           tx.ItemName(),
			Quantity:           tx.Quantity().Value(),
			TransactionType:    string(*tx.TransactionType()),
			Date:               tx.Date().Format("02/01/2006"),
			Time:               tx.Time(),
			Price:              price,
			RepresentativeUUID: representativeUUIDStr,
			Unit:               tx.Unit(),
		}
	}

	return output, nil
}
