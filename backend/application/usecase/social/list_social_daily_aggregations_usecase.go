package social

import (
	"context"
	"time"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// ListSocialDailyAggregationsUseCase handles listing social daily aggregations
type ListSocialDailyAggregationsUseCase struct {
	aggregationGateway gateway.SocialDailyAggregationGateway
}

// NewListSocialDailyAggregationsUseCase creates a new list social daily aggregations use case
func NewListSocialDailyAggregationsUseCase(aggregationGateway gateway.SocialDailyAggregationGateway) *ListSocialDailyAggregationsUseCase {
	return &ListSocialDailyAggregationsUseCase{
		aggregationGateway: aggregationGateway,
	}
}

// ListSocialDailyAggregationsInput represents input for listing social daily aggregations
type ListSocialDailyAggregationsInput struct {
	BrandID   *string `query:"brand_id"`
	Platform  *string `query:"platform"`
	StartDate *string `query:"start_date"`
	EndDate   *string `query:"end_date"`
	Page      *int    `query:"page"`
	PageSize  *int    `query:"page_size"`
	SortBy    *string `query:"sort_by"`
	SortOrder *string `query:"sort_order"`
}

// ListSocialDailyAggregationsOutput represents output of listing social daily aggregations
type ListSocialDailyAggregationsOutput struct {
	Aggregations []SocialDailyAggregationItem `json:"aggregations"`
	Total        int64                        `json:"total"`
	Page         int                          `json:"page"`
	PageSize     int                          `json:"page_size"`
	TotalPages   int                          `json:"total_pages"`
}

// SocialDailyAggregationItem represents a daily aggregation in list
type SocialDailyAggregationItem struct {
	ID              string   `json:"id"`
	BrandID         string   `json:"brand_id"`
	AggregationDate string   `json:"aggregation_date"`
	TotalPosts      int      `json:"total_posts"`
	TotalLikes      int      `json:"total_likes"`
	TotalComments   int      `json:"total_comments"`
	TotalShares     *int     `json:"total_shares,omitempty"`
	AvgLikes        float64  `json:"avg_likes"`
	AvgComments     float64  `json:"avg_comments"`
	AvgShares       *float64 `json:"avg_shares,omitempty"`
	FollowersAtDate *int     `json:"followers_at_date,omitempty"`
	EngagementRate  float64  `json:"engagement_rate"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
}

// Execute lists social daily aggregations
func (uc *ListSocialDailyAggregationsUseCase) Execute(ctx context.Context, input ListSocialDailyAggregationsInput) (*ListSocialDailyAggregationsOutput, error) {
	// Parse dates
	var startDate, endDate *time.Time
	if input.StartDate != nil {
		t, err := time.Parse("2006-01-02", *input.StartDate)
		if err != nil {
			return nil, err
		}
		startDate = &t
	}
	if input.EndDate != nil {
		t, err := time.Parse("2006-01-02", *input.EndDate)
		if err != nil {
			return nil, err
		}
		endDate = &t
	}

	// Parse sort order
	var sortOrder *valueobject.SortOrder
	if input.SortOrder != nil && input.SortBy != nil {
		so, err := valueobject.NewSortOrder(*input.SortBy, valueobject.SortDirection(*input.SortOrder))
		if err != nil {
			return nil, err
		}
		sortOrder = so
	}

	// Create criteria
	criteria := domain.NewSocialDailyAggregationCriteria()
	if input.BrandID != nil {
		criteria = criteria.WithBrandID(*input.BrandID)
	}
	if input.Platform != nil {
		platform, err := valueobject.NewSocialPlatform(*input.Platform)
		if err == nil {
			criteria = criteria.WithPlatform(platform)
		}
	}
	if startDate != nil && endDate != nil {
		criteria = criteria.WithDateRange(*startDate, *endDate)
	}

	// Set pagination
	page := 1
	if input.Page != nil && *input.Page > 0 {
		page = *input.Page
	}
	pageSize := 20
	if input.PageSize != nil && *input.PageSize > 0 {
		pageSize = *input.PageSize
	}
	criteria = criteria.WithPagination(page, pageSize)

	// Set sorting
	if input.SortBy != nil {
		criteria = criteria.WithSorting(*input.SortBy, sortOrder)
	}

	// List aggregations
	pagination := valueobject.NewPagination(page, pageSize)
	aggregations, total, err := uc.aggregationGateway.List(criteria, &pagination)
	if err != nil {
		return nil, err
	}

	// Convert to output format
	items := make([]SocialDailyAggregationItem, len(aggregations))
	for i, agg := range aggregations {
		items[i] = SocialDailyAggregationItem{
			ID:              agg.ID().String(),
			BrandID:         agg.BrandID().String(),
			AggregationDate: agg.AggregationDate().Format("2006-01-02"),
			TotalPosts:      agg.TotalPosts(),
			TotalLikes:      agg.TotalLikes(),
			TotalComments:   agg.TotalComments(),
			TotalShares:     agg.TotalShares(),
			AvgLikes:        agg.AvgLikes(),
			AvgComments:     agg.AvgComments(),
			AvgShares:       agg.AvgShares(),
			FollowersAtDate: agg.FollowersAtDate(),
			EngagementRate:  agg.EngagementRate().Value(),
			CreatedAt:       agg.CreatedAt().Format(time.RFC3339),
			UpdatedAt:       agg.UpdatedAt().Format(time.RFC3339),
		}
	}

	// Calculate total pages
	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	return &ListSocialDailyAggregationsOutput{
		Aggregations: items,
		Total:        total,
		Page:         page,
		PageSize:     pageSize,
		TotalPages:   totalPages,
	}, nil
}
