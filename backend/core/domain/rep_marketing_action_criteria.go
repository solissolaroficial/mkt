package domain

import (
	"github.com/google/uuid"
)

// RepMarketingActionCriteria representa filtros para busca de ações de marketing de representantes
type RepMarketingActionCriteria struct {
	representativeUUID *uuid.UUID
	month              *string
}

func NewRepMarketingActionCriteria() *RepMarketingActionCriteria {
	return &RepMarketingActionCriteria{}
}

func (c *RepMarketingActionCriteria) WithRepresentativeUUID(representativeUUID *uuid.UUID) *RepMarketingActionCriteria {
	c.representativeUUID = representativeUUID
	return c
}

func (c *RepMarketingActionCriteria) WithMonth(month *string) *RepMarketingActionCriteria {
	c.month = month
	return c
}

// Getters para o gateway aplicar os filtros
func (c *RepMarketingActionCriteria) RepresentativeUUID() *uuid.UUID { return c.representativeUUID }
func (c *RepMarketingActionCriteria) Month() *string                 { return c.month }

// Validate valida os critérios
func (c *RepMarketingActionCriteria) Validate() error {
	return nil
}
