package domain

// BaseCriteria define campos comuns para todos os Criteria
type BaseCriteria struct {
	month     *string
	year      *int
	page      *int
	limit     *int
	sortBy    *string
	sortOrder *string
}

// Builder methods comuns
func (c *BaseCriteria) WithMonth(month string) *BaseCriteria {
	c.month = &month
	return c
}

func (c *BaseCriteria) WithYear(year int) *BaseCriteria {
	c.year = &year
	return c
}

func (c *BaseCriteria) WithPage(page int) *BaseCriteria {
	c.page = &page
	return c
}

func (c *BaseCriteria) WithLimit(limit int) *BaseCriteria {
	c.limit = &limit
	return c
}

func (c *BaseCriteria) WithSortBy(sortBy string) *BaseCriteria {
	c.sortBy = &sortBy
	return c
}

func (c *BaseCriteria) WithSortOrder(sortOrder string) *BaseCriteria {
	c.sortOrder = &sortOrder
	return c
}

// Getters
func (c *BaseCriteria) GetMonth() *string {
	return c.month
}

func (c *BaseCriteria) GetYear() *int {
	return c.year
}

func (c *BaseCriteria) GetPage() *int {
	return c.page
}

func (c *BaseCriteria) GetLimit() *int {
	return c.limit
}

func (c *BaseCriteria) GetSortBy() *string {
	return c.sortBy
}

func (c *BaseCriteria) GetSortOrder() *string {
	return c.sortOrder
}

// IsFullYear retorna true se o mês selecionado for "---" (Ano Completo)
func (c *BaseCriteria) IsFullYear() bool {
	return c.month != nil && *c.month == "---"
}
