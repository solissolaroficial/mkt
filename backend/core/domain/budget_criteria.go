package domain

// BudgetCriteria define critérios de busca para BudgetItem
type BudgetCriteria struct {
	codObj    *string
	obj       *string
	codGrp    *string
	grp       *string
	cod       *string
	desc      *string
	year      *int
	page      *int
	limit     *int
	sortBy    *string
	sortOrder *string
}

// NewBudgetCriteria cria um novo BudgetCriteria com valores padrão
func NewBudgetCriteria() *BudgetCriteria {
	defaultPage := 1
	defaultLimit := 50
	defaultSortBy := "codObj"
	defaultSortOrder := "asc"

	return &BudgetCriteria{
		page:      &defaultPage,
		limit:     &defaultLimit,
		sortBy:    &defaultSortBy,
		sortOrder: &defaultSortOrder,
	}
}

// WithCodObj define filtro por código do objeto
func (c *BudgetCriteria) WithCodObj(codObj string) *BudgetCriteria {
	c.codObj = &codObj
	return c
}

// WithObj define filtro por nome do objeto (LIKE)
func (c *BudgetCriteria) WithObj(obj string) *BudgetCriteria {
	c.obj = &obj
	return c
}

// WithCodGrp define filtro por código do grupo
func (c *BudgetCriteria) WithCodGrp(codGrp string) *BudgetCriteria {
	c.codGrp = &codGrp
	return c
}

// WithGrp define filtro por nome do grupo (LIKE)
func (c *BudgetCriteria) WithGrp(grp string) *BudgetCriteria {
	c.grp = &grp
	return c
}

// WithCod define filtro por código do item
func (c *BudgetCriteria) WithCod(cod string) *BudgetCriteria {
	c.cod = &cod
	return c
}

// WithDesc define filtro por descrição (LIKE)
func (c *BudgetCriteria) WithDesc(desc string) *BudgetCriteria {
	c.desc = &desc
	return c
}

// WithYear define filtro por ano
func (c *BudgetCriteria) WithYear(year int) *BudgetCriteria {
	c.year = &year
	return c
}

// WithPage define a página atual
func (c *BudgetCriteria) WithPage(page int) *BudgetCriteria {
	c.page = &page
	return c
}

// WithLimit define o limite de itens por página
func (c *BudgetCriteria) WithLimit(limit int) *BudgetCriteria {
	c.limit = &limit
	return c
}

// WithSortBy define o campo de ordenação
// Opções: codObj, obj, codGrp, grp, cod, desc, createdAt
func (c *BudgetCriteria) WithSortBy(sortBy string) *BudgetCriteria {
	c.sortBy = &sortBy
	return c
}

// WithSortOrder define a ordem de ordenação
// Opções: asc, desc
func (c *BudgetCriteria) WithSortOrder(sortOrder string) *BudgetCriteria {
	c.sortOrder = &sortOrder
	return c
}

// Getters
func (c *BudgetCriteria) GetCodObj() *string {
	return c.codObj
}

func (c *BudgetCriteria) GetObj() *string {
	return c.obj
}

func (c *BudgetCriteria) GetCodGrp() *string {
	return c.codGrp
}

func (c *BudgetCriteria) GetGrp() *string {
	return c.grp
}

func (c *BudgetCriteria) GetCod() *string {
	return c.cod
}

func (c *BudgetCriteria) GetDesc() *string {
	return c.desc
}

func (c *BudgetCriteria) GetYear() *int {
	return c.year
}

func (c *BudgetCriteria) GetPage() *int {
	return c.page
}

func (c *BudgetCriteria) GetLimit() *int {
	return c.limit
}

func (c *BudgetCriteria) GetSortBy() *string {
	return c.sortBy
}

func (c *BudgetCriteria) GetSortOrder() *string {
	return c.sortOrder
}
