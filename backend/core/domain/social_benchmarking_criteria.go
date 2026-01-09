package domain

import (
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// SocialBenchmarkingCriteria representa filtros para busca de benchmarkings
// NOTA: Este criteria NÃO depende de GORM, seguindo Clean Architecture
type SocialBenchmarkingCriteria struct {
	brandName *valueobject.BrandName
	active    *bool
	startDate *string
	endDate   *string
}

func NewSocialBenchmarkingCriteria() *SocialBenchmarkingCriteria {
	return &SocialBenchmarkingCriteria{}
}

func (c *SocialBenchmarkingCriteria) WithBrandName(brandName *valueobject.BrandName) *SocialBenchmarkingCriteria {
	c.brandName = brandName
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
func (c *SocialBenchmarkingCriteria) BrandName() *valueobject.BrandName { return c.brandName }
func (c *SocialBenchmarkingCriteria) Active() *bool                     { return c.active }
func (c *SocialBenchmarkingCriteria) StartDate() *string                { return c.startDate }
func (c *SocialBenchmarkingCriteria) EndDate() *string                  { return c.endDate }
