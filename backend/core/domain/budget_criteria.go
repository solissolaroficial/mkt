package domain

import (
	"github.com/google/uuid"
)

// BudgetCriteria define critérios de busca para BudgetItem
type BudgetCriteria struct {
	objectUUID   *uuid.UUID
	objectName   *string
	groupUUID    *uuid.UUID
	groupName    *string
	cod          *string
	desc         *string
	year         *int
	page         *int
	limit        *int
	sortBy       *string
	sortOrder    *string
}

// NewBudgetCriteria cria um novo BudgetCriteria com valores padrão
func NewBudgetCriteria() *BudgetCriteria {
	defaultPage := 1
	defaultLimit := 50
	defaultSortBy := "objectUUID"
	defaultSortOrder := "asc"

	return &BudgetCriteria{
		page:      &defaultPage,
		limit:     &defaultLimit,
		sortBy:    &defaultSortBy,
		sortOrder: &defaultSortOrder,
	}
}

// WithObjectUUID define filtro por UUID do objeto
func (c *BudgetCriteria) WithObjectUUID(objectUUID *uuid.UUID) *BudgetCriteria {
	c.objectUUID = objectUUID
	return c
}

// WithObjectName define filtro por nome do objeto (LIKE)
func (c *BudgetCriteria) WithObjectName(objectName *string) *BudgetCriteria {
	c.objectName = objectName
	return c
}

// WithGroupUUID define filtro por UUID do grupo
func (c *BudgetCriteria) WithGroupUUID(groupUUID *uuid.UUID) *BudgetCriteria {
	c.groupUUID = groupUUID
	return c
}

// WithGroupName define filtro por nome do grupo (LIKE)
func (c *BudgetCriteria) WithGroupName(groupName *string) *BudgetCriteria {
	c.groupName = groupName
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
// Opções: objectUUID, objectName, groupUUID, groupName, cod, desc, createdAt
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
func (c *BudgetCriteria) GetObjectUUID() *uuid.UUID {
	return c.objectUUID
}

func (c *BudgetCriteria) GetObjectName() *string {
	return c.objectName
}

func (c *BudgetCriteria) GetGroupUUID() *uuid.UUID {
	return c.groupUUID
}

func (c *BudgetCriteria) GetGroupName() *string {
	return c.groupName
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
