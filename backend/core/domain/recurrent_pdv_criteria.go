package domain

// RecurrentPdvCriteria representa filtros para busca de PDVs recorrentes
// NOTA: Este criteria NÃO depende de GORM, seguindo Clean Architecture
type RecurrentPdvCriteria struct {
	repName *string
	city    *string
}

// NewRecurrentPdvCriteria cria um novo RecurrentPdvCriteria vazio
func NewRecurrentPdvCriteria() *RecurrentPdvCriteria {
	return &RecurrentPdvCriteria{}
}

// WithRepName adiciona filtro por nome do representante
func (c *RecurrentPdvCriteria) WithRepName(repName *string) *RecurrentPdvCriteria {
	c.repName = repName
	return c
}

// WithCity adiciona filtro por cidade
func (c *RecurrentPdvCriteria) WithCity(city *string) *RecurrentPdvCriteria {
	c.city = city
	return c
}

// Getters para o gateway aplicar os filtros

// RepName retorna o filtro de nome do representante
func (c *RecurrentPdvCriteria) RepName() *string {
	return c.repName
}

// City retorna o filtro de cidade
func (c *RecurrentPdvCriteria) City() *string {
	return c.city
}

// Validate valida os critérios de busca
func (c *RecurrentPdvCriteria) Validate() error {
	// Sem validações específicas necessárias por enquanto
	return nil
}
