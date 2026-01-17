package domain

import (
	"github.com/google/uuid"
)

// ShowroomItemCriteria representa filtros para busca de itens de showroom
type ShowroomItemCriteria struct {
	representativeUUID *uuid.UUID
	delivered          *bool
	city               *string
}

func NewShowroomItemCriteria() *ShowroomItemCriteria {
	return &ShowroomItemCriteria{}
}

func (c *ShowroomItemCriteria) WithRepresentativeUUID(representativeUUID *uuid.UUID) *ShowroomItemCriteria {
	c.representativeUUID = representativeUUID
	return c
}

func (c *ShowroomItemCriteria) WithDelivered(delivered *bool) *ShowroomItemCriteria {
	c.delivered = delivered
	return c
}

func (c *ShowroomItemCriteria) WithCity(city *string) *ShowroomItemCriteria {
	c.city = city
	return c
}

// Getters para o gateway aplicar os filtros
func (c *ShowroomItemCriteria) RepresentativeUUID() *uuid.UUID { return c.representativeUUID }
func (c *ShowroomItemCriteria) Delivered() *bool               { return c.delivered }
func (c *ShowroomItemCriteria) City() *string                  { return c.city }

// Validate valida os critérios
func (c *ShowroomItemCriteria) Validate() error {
	return nil
}
