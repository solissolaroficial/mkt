package domain

// RepMarketingActionCriteria representa filtros para busca de ações de marketing de representantes
type RepMarketingActionCriteria struct {
	repName *string
	month   *string
}

func NewRepMarketingActionCriteria() *RepMarketingActionCriteria {
	return &RepMarketingActionCriteria{}
}

func (c *RepMarketingActionCriteria) WithRepName(repName *string) *RepMarketingActionCriteria {
	c.repName = repName
	return c
}

func (c *RepMarketingActionCriteria) WithMonth(month *string) *RepMarketingActionCriteria {
	c.month = month
	return c
}

// Getters para o gateway aplicar os filtros
func (c *RepMarketingActionCriteria) RepName() *string { return c.repName }
func (c *RepMarketingActionCriteria) Month() *string   { return c.month }

// Validate valida os critérios
func (c *RepMarketingActionCriteria) Validate() error {
	return nil
}
