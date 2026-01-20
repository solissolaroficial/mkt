package request

type CreateSocialBenchmarkingRequest struct {
	BrandID     string  `json:"brand_id" validate:"required,max=200"`
	AvgLikes    float64 `json:"avg_likes" validate:"required,min=0"`
	AvgComments float64 `json:"avg_comments" validate:"required,min=0"`
	Followers   *int    `json:"followers" validate:"omitempty,min=0"`
}

type UpdateSocialBenchmarkingRequest struct {
	BrandID     *string  `json:"brand_id" validate:"omitempty,max=200"`
	AvgLikes    *float64 `json:"avg_likes" validate:"omitempty,min=0"`
	AvgComments *float64 `json:"avg_comments" validate:"omitempty,min=0"`
	Followers   *int     `json:"followers" validate:"omitempty,min=0"`
}

type ListSocialBenchmarkingsQuery struct {
	BrandID   *string `query:"brand_id" validate:"omitempty,max=200"`
	Active    *bool   `query:"active" validate:"omitempty"`
	StartDate *string `query:"start_date" validate:"omitempty,datetime=2006-01-02"`
	EndDate   *string `query:"end_date" validate:"omitempty,datetime=2006-01-02"`
	Page      int     `query:"page" validate:"omitempty,min=1"`
	Limit     int     `query:"limit" validate:"omitempty,min=1,max=100"`
	SortBy    *string `query:"sort_by" validate:"omitempty,oneof=engagement_rate avg_likes avg_comments created_at"`
	SortOrder *string `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}

type ListSocialPostsRequest struct {
	BrandID       *string  `query:"brand_id" validate:"omitempty,max=200"`
	Platform      *string  `query:"platform" validate:"omitempty,oneof=instagram facebook linkedin tiktok twitter"`
	PostType      *string  `query:"post_type" validate:"omitempty,oneof=image video carousel reel story"`
	StartDate     *string  `query:"start_date" validate:"omitempty,datetime=2006-01-02"`
	EndDate       *string  `query:"end_date" validate:"omitempty,datetime=2006-01-02"`
	MinFollowers  *int     `query:"min_followers" validate:"omitempty,min=0"`
	MaxFollowers  *int     `query:"max_followers" validate:"omitempty,min=0"`
	MinEngagement *float64 `query:"min_engagement" validate:"omitempty,min=0"`
	MaxEngagement *float64 `query:"max_engagement" validate:"omitempty,min=0"`
	OnlyActive    *bool    `query:"only_active"`
	Page          *int     `query:"page" validate:"omitempty,min=1"`
	PageSize      *int     `query:"page_size" validate:"omitempty,min=1,max=100"`
	SortBy        *string  `query:"sort_by" validate:"omitempty,oneof=post_date likes comments created_at"`
	SortOrder     *string  `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}

type ListSocialDailyAggregationsRequest struct {
	BrandID   *string `query:"brand_id" validate:"omitempty,max=200"`
	Platform  *string `query:"platform" validate:"omitempty,oneof=instagram facebook linkedin tiktok twitter"`
	StartDate *string `query:"start_date" validate:"omitempty,datetime=2006-01-02"`
	EndDate   *string `query:"end_date" validate:"omitempty,datetime=2006-01-02"`
	Page      *int    `query:"page" validate:"omitempty,min=1"`
	PageSize  *int    `query:"page_size" validate:"omitempty,min=1,max=100"`
	SortBy    *string `query:"sort_by" validate:"omitempty,oneof=aggregation_date total_posts avg_likes avg_comments engagement_rate"`
	SortOrder *string `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}

// Social Post Request DTOs
type CreateSocialPostRequest struct {
	BrandID         string  `json:"brand_id" validate:"required,max=200"`
	Platform        string  `json:"platform" validate:"required,oneof=instagram facebook linkedin tiktok twitter"`
	PostDate        string  `json:"post_date" validate:"required"`
	PostTime        *string `json:"post_time"`
	Likes           int     `json:"likes" validate:"required,min=0"`
	Comments        int     `json:"comments" validate:"required,min=0"`
	Shares          *int    `json:"shares" validate:"omitempty,min=0"`
	PostType        string  `json:"post_type" validate:"required,oneof=image video carousel reel story"`
	Caption         *string `json:"caption"`
	FollowersAtPost *int    `json:"followers_at_post" validate:"omitempty,min=0"`
}

type UpdateSocialPostRequest struct {
	ID              string  `json:"id" validate:"required"`
	BrandID         *string `json:"brand_id,omitempty"`
	Platform        *string `json:"platform,omitempty"`
	PostDate        *string `json:"post_date,omitempty"`
	PostTime        *string `json:"post_time,omitempty"`
	Likes           *int    `json:"likes,omitempty"`
	Comments        *int    `json:"comments,omitempty"`
	Shares          *int    `json:"shares,omitempty"`
	PostType        *string `json:"post_type,omitempty"`
	Caption         *string `json:"caption,omitempty"`
	FollowersAtPost *int    `json:"followers_at_post,omitempty"`
}
