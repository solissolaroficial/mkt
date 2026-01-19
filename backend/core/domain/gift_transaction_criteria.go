package domain

import (
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// GiftTransactionCriteria representa filtros para busca de transações de brinde
// CORREÇÃO: itemName usa string para consistência com GiftTransaction entity
type GiftTransactionCriteria struct {
	itemName           *string // CORRIGIDO: era *valueobject.GiftName
	transactionType    *valueobject.TransactionType
	representativeUUID *string
	startDate          *string
	endDate            *string
	page               *int
	limit              *int
}

func NewGiftTransactionCriteria() *GiftTransactionCriteria {
	return &GiftTransactionCriteria{}
}

func (c *GiftTransactionCriteria) WithItemName(itemName *string) *GiftTransactionCriteria {
	c.itemName = itemName
	return c
}

func (c *GiftTransactionCriteria) WithTransactionType(transactionType *valueobject.TransactionType) *GiftTransactionCriteria {
	c.transactionType = transactionType
	return c
}

func (c *GiftTransactionCriteria) WithRepresentativeUUID(representativeUUID *string) *GiftTransactionCriteria {
	c.representativeUUID = representativeUUID
	return c
}

func (c *GiftTransactionCriteria) WithStartDate(startDate *string) *GiftTransactionCriteria {
	c.startDate = startDate
	return c
}

func (c *GiftTransactionCriteria) WithEndDate(endDate *string) *GiftTransactionCriteria {
	c.endDate = endDate
	return c
}

func (c *GiftTransactionCriteria) WithPage(page *int) *GiftTransactionCriteria {
	c.page = page
	return c
}

func (c *GiftTransactionCriteria) WithLimit(limit *int) *GiftTransactionCriteria {
	c.limit = limit
	return c
}

// Getters para o gateway aplicar os filtros
func (c *GiftTransactionCriteria) ItemName() *string { return c.itemName }
func (c *GiftTransactionCriteria) TransactionType() *valueobject.TransactionType {
	return c.transactionType
}
func (c *GiftTransactionCriteria) RepresentativeUUID() *string {
	if c.representativeUUID == nil {
		return nil
	}
	return c.representativeUUID
}
func (c *GiftTransactionCriteria) StartDate() *string { return c.startDate }
func (c *GiftTransactionCriteria) EndDate() *string   { return c.endDate }
func (c *GiftTransactionCriteria) Page() *int         { return c.page }
func (c *GiftTransactionCriteria) Limit() *int        { return c.limit }

// GetOffset calcula o offset para paginação
func (c *GiftTransactionCriteria) GetOffset() int {
	if c.page == nil || c.limit == nil {
		return 0
	}
	return (*c.page - 1) * *c.limit
}
