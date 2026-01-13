package domain

import (
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// GiftItemCriteria representa filtros para busca de itens de brinde
// NOTA: Este criteria NÃO depende de GORM, seguindo Clean Architecture
type GiftItemCriteria struct {
	name     *valueobject.GiftName
	minStock *int
	maxStock *int
	minPrice *float64
	maxPrice *float64
	page     *int
	limit    *int
}

func NewGiftItemCriteria() *GiftItemCriteria {
	return &GiftItemCriteria{}
}

func (c *GiftItemCriteria) WithName(name *valueobject.GiftName) *GiftItemCriteria {
	c.name = name
	return c
}

func (c *GiftItemCriteria) WithMinStock(minStock *int) *GiftItemCriteria {
	c.minStock = minStock
	return c
}

func (c *GiftItemCriteria) WithMaxStock(maxStock *int) *GiftItemCriteria {
	c.maxStock = maxStock
	return c
}

func (c *GiftItemCriteria) WithMinPrice(minPrice *float64) *GiftItemCriteria {
	c.minPrice = minPrice
	return c
}

func (c *GiftItemCriteria) WithMaxPrice(maxPrice *float64) *GiftItemCriteria {
	c.maxPrice = maxPrice
	return c
}

func (c *GiftItemCriteria) WithPage(page *int) *GiftItemCriteria {
	c.page = page
	return c
}

func (c *GiftItemCriteria) WithLimit(limit *int) *GiftItemCriteria {
	c.limit = limit
	return c
}

// Getters para o gateway aplicar os filtros
func (c *GiftItemCriteria) Name() *valueobject.GiftName { return c.name }
func (c *GiftItemCriteria) MinStock() *int              { return c.minStock }
func (c *GiftItemCriteria) MaxStock() *int              { return c.maxStock }
func (c *GiftItemCriteria) MinPrice() *float64          { return c.minPrice }
func (c *GiftItemCriteria) MaxPrice() *float64          { return c.maxPrice }
func (c *GiftItemCriteria) Page() *int                  { return c.page }
func (c *GiftItemCriteria) Limit() *int                 { return c.limit }

// GetOffset calcula o offset para paginação
func (c *GiftItemCriteria) GetOffset() int {
	if c.page == nil || c.limit == nil {
		return 0
	}
	return (*c.page - 1) * *c.limit
}
