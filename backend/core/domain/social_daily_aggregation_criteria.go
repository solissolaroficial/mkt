package domain

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// SocialDailyAggregationCriteria defines criteria for filtering daily aggregations
// This follows the DDD pattern of keeping criteria in the domain layer
type SocialDailyAggregationCriteria struct {
	// Filter fields
	BrandID   *string
	Platform  *valueobject.SocialPlatform
	StartDate *time.Time
	EndDate   *time.Time

	// Pagination
	Page     int
	PageSize int

	// Sorting
	SortBy    string
	SortOrder *valueobject.SortOrder
}

// NewSocialDailyAggregationCriteria creates a new SocialDailyAggregationCriteria with default values
func NewSocialDailyAggregationCriteria() *SocialDailyAggregationCriteria {
	sortOrder, _ := valueobject.NewSortOrder("date", valueobject.SortDirectionDesc)
	return &SocialDailyAggregationCriteria{
		Page:      1,
		PageSize:  20,
		SortBy:    "date",
		SortOrder: sortOrder,
	}
}

// WithBrandID sets the brand ID filter
func (c *SocialDailyAggregationCriteria) WithBrandID(brandID string) *SocialDailyAggregationCriteria {
	c.BrandID = &brandID
	return c
}

// WithPlatform sets the platform filter
func (c *SocialDailyAggregationCriteria) WithPlatform(platform valueobject.SocialPlatform) *SocialDailyAggregationCriteria {
	c.Platform = &platform
	return c
}

// WithDateRange sets the date range filter
func (c *SocialDailyAggregationCriteria) WithDateRange(startDate, endDate time.Time) *SocialDailyAggregationCriteria {
	c.StartDate = &startDate
	c.EndDate = &endDate
	return c
}

// WithPagination sets pagination
func (c *SocialDailyAggregationCriteria) WithPagination(page, pageSize int) *SocialDailyAggregationCriteria {
	c.Page = page
	c.PageSize = pageSize
	return c
}

// WithSorting sets sorting
func (c *SocialDailyAggregationCriteria) WithSorting(sortBy string, sortOrder *valueobject.SortOrder) *SocialDailyAggregationCriteria {
	c.SortBy = sortBy
	c.SortOrder = sortOrder
	return c
}

// HasFilters returns true if any filter is set
func (c *SocialDailyAggregationCriteria) HasFilters() bool {
	return c.BrandID != nil ||
		c.Platform != nil ||
		c.StartDate != nil ||
		c.EndDate != nil
}

// GetOffset returns the offset for pagination
func (c *SocialDailyAggregationCriteria) GetOffset() int {
	if c.Page < 1 {
		c.Page = 1
	}
	return (c.Page - 1) * c.PageSize
}

// GetLimit returns the limit for pagination
func (c *SocialDailyAggregationCriteria) GetLimit() int {
	if c.PageSize < 1 {
		c.PageSize = 20
	}
	return c.PageSize
}
