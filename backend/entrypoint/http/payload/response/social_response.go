package response

type SocialBenchmarkingResponse struct {
	ID             string  `json:"id"`
	BrandID        string  `json:"brand_id"`
	AvgLikes       float64 `json:"avg_likes"`
	AvgComments    float64 `json:"avg_comments"`
	Followers      *int    `json:"followers"`
	EngagementRate float64 `json:"engagement_rate"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}

// SocialBenchmarkingListData representa os dados da lista de benchmarkings
type SocialBenchmarkingListData struct {
	Benchmarkings []SocialBenchmarkingResponse `json:"benchmarkings"`
	Meta          MetaResponse                 `json:"meta"`
}
