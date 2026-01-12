package domain

// ShowroomItemCriteria representa filtros para busca de itens de showroom
type ShowroomItemCriteria struct {
	repName   *string
	delivered *bool
	city      *string
}

func NewShowroomItemCriteria() *ShowroomItemCriteria {
	return &ShowroomItemCriteria{}
}

func (c *ShowroomItemCriteria) WithRepName(repName *string) *ShowroomItemCriteria {
	c.repName = repName
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
func (c *ShowroomItemCriteria) RepName() *string { return c.repName }
func (c *ShowroomItemCriteria) Delivered() *bool { return c.delivered }
func (c *ShowroomItemCriteria) City() *string    { return c.city }

// Validate valida os critérios
func (c *ShowroomItemCriteria) Validate() error {
	return nil
}
