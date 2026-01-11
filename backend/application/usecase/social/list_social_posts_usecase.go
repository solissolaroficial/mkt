package social

import (
	"context"
	"time"

	"github.com/seu-usuario/solis-backend/core/domain"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// ListSocialPostsUseCase handles listing social posts
type ListSocialPostsUseCase struct {
	postGateway gateway.SocialPostGateway
}

// NewListSocialPostsUseCase creates a new list social posts use case
func NewListSocialPostsUseCase(postGateway gateway.SocialPostGateway) *ListSocialPostsUseCase {
	return &ListSocialPostsUseCase{
		postGateway: postGateway,
	}
}

// ListSocialPostsInput represents the input for listing social posts
type ListSocialPostsInput struct {
	BrandName     *string  `query:"brand_name"`
	Platform      *string  `query:"platform"`
	PostType      *string  `query:"post_type"`
	StartDate     *string  `query:"start_date"`
	EndDate       *string  `query:"end_date"`
	MinFollowers  *int     `query:"min_followers"`
	MaxFollowers  *int     `query:"max_followers"`
	MinEngagement *float64 `query:"min_engagement"`
	MaxEngagement *float64 `query:"max_engagement"`
	OnlyActive    *bool    `query:"only_active"`
	Page          *int     `query:"page"`
	PageSize      *int     `query:"page_size"`
	SortBy        *string  `query:"sort_by"`
	SortOrder     *string  `query:"sort_order"`
}

// ListSocialPostsOutput represents the output of listing social posts
type ListSocialPostsOutput struct {
	Posts      []SocialPostItem `json:"posts"`
	Total      int64            `json:"total"`
	Page       int              `json:"page"`
	PageSize   int              `json:"page_size"`
	TotalPages int              `json:"total_pages"`
}

// SocialPostItem represents a social post in list
type SocialPostItem struct {
	ID              string   `json:"id"`
	BrandName       string   `json:"brand_name"`
	Platform        string   `json:"platform"`
	PostDate        string   `json:"post_date"`
	PostTime        *string  `json:"post_time,omitempty"`
	Likes           int      `json:"likes"`
	Comments        int      `json:"comments"`
	Shares          *int     `json:"shares,omitempty"`
	PostType        string   `json:"post_type"`
	Caption         *string  `json:"caption,omitempty"`
	FollowersAtPost *int     `json:"followers_at_post,omitempty"`
	EngagementRate  *float64 `json:"engagement_rate,omitempty"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
}

// Execute lists social posts
func (uc *ListSocialPostsUseCase) Execute(ctx context.Context, input ListSocialPostsInput) (*ListSocialPostsOutput, error) {
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
	criteria := domain.NewSocialPostCriteria()
	if input.BrandName != nil {
		criteria = criteria.WithBrandName(*input.BrandName)
	}
	if input.Platform != nil {
		platform, err := valueobject.NewSocialPlatform(*input.Platform)
		if err == nil {
			criteria = criteria.WithPlatform(platform)
		}
	}
	if input.PostType != nil {
		postType, err := valueobject.NewSocialPostType(*input.PostType)
		if err == nil {
			criteria = criteria.WithPostType(postType)
		}
	}
	if startDate != nil && endDate != nil {
		criteria = criteria.WithDateRange(*startDate, *endDate)
	}
	if input.MinFollowers != nil && input.MaxFollowers != nil {
		criteria = criteria.WithFollowersRange(*input.MinFollowers, *input.MaxFollowers)
	}
	if input.MinEngagement != nil && input.MaxEngagement != nil {
		criteria = criteria.WithEngagementRange(*input.MinEngagement, *input.MaxEngagement)
	}
	if input.OnlyActive != nil {
		criteria = criteria.WithOnlyActive(*input.OnlyActive)
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

	// List posts
	pagination := valueobject.NewPagination(page, pageSize)
	posts, total, err := uc.postGateway.List(criteria, &pagination)
	if err != nil {
		return nil, err
	}

	// Convert to output format
	items := make([]SocialPostItem, len(posts))
	for i, post := range posts {
		items[i] = SocialPostItem{
			ID:              post.ID().String(),
			BrandName:       post.BrandName().Value(),
			Platform:        post.Platform().String(),
			PostDate:        post.PostDate().Format("2006-01-02"),
			Likes:           post.Likes(),
			Comments:        post.Comments(),
			Shares:          post.Shares(),
			PostType:        post.PostType().String(),
			Caption:         post.Caption(),
			FollowersAtPost: post.FollowersAtPost(),
			CreatedAt:       post.CreatedAt().Format(time.RFC3339),
			UpdatedAt:       post.UpdatedAt().Format(time.RFC3339),
		}

		// Format post time
		if post.PostTime() != nil {
			formatted := post.PostTime().Format("15:04:05")
			items[i].PostTime = &formatted
		}

		// Calculate engagement rate
		if post.FollowersAtPost() != nil && *post.FollowersAtPost() > 0 {
			engagement := float64(post.Likes()+post.Comments()) / float64(*post.FollowersAtPost()) * 100
			items[i].EngagementRate = &engagement
		}
	}

	// Calculate total pages
	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	return &ListSocialPostsOutput{
		Posts:      items,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}
