package domain

import (
	"github.com/google/uuid"
)

// SocialBenchmarkingCriteria representa filtros para busca de benchmarkings
// NOTA: Este criteria NÃO depende de GORM, seguindo Clean Architecture
type SocialBenchmarkingCriteria struct {
	brandID   *uuid.UUID
	active    *bool
	startDate *string
	endDate   *string
}

func NewSocialBenchmarkingCriteria() *SocialBenchmarkingCriteria {
	return &SocialBenchmarkingCriteria{}
}

func (c *SocialBenchmarkingCriteria) WithBrandID(brandID *uuid.UUID) *SocialBenchmarkingCriteria {
	c.brandID = brandID
	return c
}

func (c *SocialBenchmarkingCriteria) WithActive(active *bool) *SocialBenchmarkingCriteria {
	c.active = active
	return c
}

func (c *SocialBenchmarkingCriteria) WithStartDate(startDate *string) *SocialBenchmarkingCriteria {
	c.startDate = startDate
	return c
}

func (c *SocialBenchmarkingCriteria) WithEndDate(endDate *string) *SocialBenchmarkingCriteria {
	c.endDate = endDate
	return c
}

// Getters para o gateway aplicar os filtros
func (c *SocialBenchmarkingCriteria) BrandID() *uuid.UUID { return c.brandID }
func (c *SocialBenchmarkingCriteria) Active() *bool       { return c.active }
func (c *SocialBenchmarkingCriteria) StartDate() *string  { return c.startDate }
func (c *SocialBenchmarkingCriteria) EndDate() *string    { return c.endDate }
