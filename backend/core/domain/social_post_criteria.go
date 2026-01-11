package domain

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// SocialPostCriteria defines criteria for filtering social posts
// This follows the DDD pattern of keeping criteria in the domain layer
type SocialPostCriteria struct {
	// Filter fields
	BrandName     *string
	Platform      *valueobject.SocialPlatform
	PostType      *valueobject.SocialPostType
	StartDate     *time.Time
	EndDate       *time.Time
	MinFollowers  *int
	MaxFollowers  *int
	MinEngagement *float64
	MaxEngagement *float64
	OnlyActive    *bool

	// Pagination
	Page     int
	PageSize int

	// Sorting
	SortBy    string
	SortOrder *valueobject.SortOrder
}

// NewSocialPostCriteria creates a new SocialPostCriteria with default values
func NewSocialPostCriteria() *SocialPostCriteria {
	sortOrder, _ := valueobject.NewSortOrder("post_date", valueobject.SortDirectionDesc)
	return &SocialPostCriteria{
		Page:      1,
		PageSize:  20,
		SortBy:    "post_date",
		SortOrder: sortOrder,
	}
}

// WithBrandName sets the brand name filter
func (c *SocialPostCriteria) WithBrandName(brandName string) *SocialPostCriteria {
	c.BrandName = &brandName
	return c
}

// WithPlatform sets the platform filter
func (c *SocialPostCriteria) WithPlatform(platform valueobject.SocialPlatform) *SocialPostCriteria {
	c.Platform = &platform
	return c
}

// WithPostType sets the post type filter
func (c *SocialPostCriteria) WithPostType(postType valueobject.SocialPostType) *SocialPostCriteria {
	c.PostType = &postType
	return c
}

// WithDateRange sets the date range filter
func (c *SocialPostCriteria) WithDateRange(startDate, endDate time.Time) *SocialPostCriteria {
	c.StartDate = &startDate
	c.EndDate = &endDate
	return c
}

// WithFollowersRange sets the followers range filter
func (c *SocialPostCriteria) WithFollowersRange(minFollowers, maxFollowers int) *SocialPostCriteria {
	c.MinFollowers = &minFollowers
	c.MaxFollowers = &maxFollowers
	return c
}

// WithEngagementRange sets the engagement rate range filter
func (c *SocialPostCriteria) WithEngagementRange(minEngagement, maxEngagement float64) *SocialPostCriteria {
	c.MinEngagement = &minEngagement
	c.MaxEngagement = &maxEngagement
	return c
}

// WithOnlyActive sets the filter to only show active (non-deleted) posts
func (c *SocialPostCriteria) WithOnlyActive(onlyActive bool) *SocialPostCriteria {
	c.OnlyActive = &onlyActive
	return c
}

// WithPagination sets pagination
func (c *SocialPostCriteria) WithPagination(page, pageSize int) *SocialPostCriteria {
	c.Page = page
	c.PageSize = pageSize
	return c
}

// WithSorting sets sorting
func (c *SocialPostCriteria) WithSorting(sortBy string, sortOrder *valueobject.SortOrder) *SocialPostCriteria {
	c.SortBy = sortBy
	c.SortOrder = sortOrder
	return c
}

// HasFilters returns true if any filter is set
func (c *SocialPostCriteria) HasFilters() bool {
	return c.BrandName != nil ||
		c.Platform != nil ||
		c.PostType != nil ||
		c.StartDate != nil ||
		c.EndDate != nil ||
		c.MinFollowers != nil ||
		c.MaxFollowers != nil ||
		c.MinEngagement != nil ||
		c.MaxEngagement != nil ||
		c.OnlyActive != nil
}

// GetOffset returns the offset for pagination
func (c *SocialPostCriteria) GetOffset() int {
	if c.Page < 1 {
		c.Page = 1
	}
	return (c.Page - 1) * c.PageSize
}

// GetLimit returns the limit for pagination
func (c *SocialPostCriteria) GetLimit() int {
	if c.PageSize < 1 {
		c.PageSize = 20
	}
	return c.PageSize
}
