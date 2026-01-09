package request

type CreateSocialBenchmarkingRequest struct {
	BrandName   string  `json:"brand_name" validate:"required,max=200"`
	AvgLikes    float64 `json:"avg_likes" validate:"required,min=0"`
	AvgComments float64 `json:"avg_comments" validate:"required,min=0"`
	Followers   *int    `json:"followers" validate:"omitempty,min=0"`
}

type UpdateSocialBenchmarkingRequest struct {
	BrandName   *string  `json:"brand_name" validate:"omitempty,max=200"`
	AvgLikes    *float64 `json:"avg_likes" validate:"omitempty,min=0"`
	AvgComments *float64 `json:"avg_comments" validate:"omitempty,min=0"`
	Followers   *int     `json:"followers" validate:"omitempty,min=0"`
}

type ListSocialBenchmarkingsQuery struct {
	BrandName *string `query:"brand_name" validate:"omitempty,max=200"`
	Active    *bool   `query:"active" validate:"omitempty"`
	StartDate *string `query:"start_date" validate:"omitempty,datetime=2006-01-02"`
	EndDate   *string `query:"end_date" validate:"omitempty,datetime=2006-01-02"`
	Page      int     `query:"page" validate:"omitempty,min=1"`
	Limit     int     `query:"limit" validate:"omitempty,min=1,max=100"`
	SortBy    *string `query:"sort_by" validate:"omitempty,oneof=engagement_rate avg_likes avg_comments created_at"`
	SortOrder *string `query:"sort_order" validate:"omitempty,oneof=asc desc"`
}
