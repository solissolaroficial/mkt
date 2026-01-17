package domain

import (
	"github.com/google/uuid"
)

// RecurrentPdvCriteria representa filtros para busca de PDVs recorrentes
// NOTA: Este criteria NÃO depende de GORM, seguindo Clean Architecture
type RecurrentPdvCriteria struct {
	representativeUUID *uuid.UUID
	city               *string
}

// NewRecurrentPdvCriteria cria um novo RecurrentPdvCriteria vazio
func NewRecurrentPdvCriteria() *RecurrentPdvCriteria {
	return &RecurrentPdvCriteria{}
}

// WithRepresentativeUUID adiciona filtro por UUID do representante
func (c *RecurrentPdvCriteria) WithRepresentativeUUID(representativeUUID *uuid.UUID) *RecurrentPdvCriteria {
	c.representativeUUID = representativeUUID
	return c
}

// WithCity adiciona filtro por cidade
func (c *RecurrentPdvCriteria) WithCity(city *string) *RecurrentPdvCriteria {
	c.city = city
	return c
}

// Getters para o gateway aplicar os filtros

// RepresentativeUUID retorna o filtro de UUID do representante
func (c *RecurrentPdvCriteria) RepresentativeUUID() *uuid.UUID {
	return c.representativeUUID
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
