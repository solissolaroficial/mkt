package response

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

type SocialPostResponse struct {
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

// SocialPostToResponse converts a SocialPost entity to a SocialPostResponse
func SocialPostToResponse(post *entity.SocialPost) SocialPostResponse {
	// Format post time for output
	var postTimeStr *string
	if post.PostTime() != nil {
		formatted := post.PostTime().Format("15:04:05")
		postTimeStr = &formatted
	}

	// Calculate engagement rate
	var engagementRate *float64
	if post.FollowersAtPost() != nil && *post.FollowersAtPost() > 0 {
		engagement := float64(post.Likes()+post.Comments()) / float64(*post.FollowersAtPost()) * 100
		engagementRate = &engagement
	}

	return SocialPostResponse{
		ID:              post.ID().String(),
		BrandName:       post.BrandName().Value(),
		Platform:        post.Platform().String(),
		PostDate:        post.PostDate().Format("2006-01-02"),
		PostTime:        postTimeStr,
		Likes:           post.Likes(),
		Comments:        post.Comments(),
		Shares:          post.Shares(),
		PostType:        post.PostType().String(),
		Caption:         post.Caption(),
		FollowersAtPost: post.FollowersAtPost(),
		EngagementRate:  engagementRate,
		CreatedAt:       post.CreatedAt().Format(time.RFC3339),
		UpdatedAt:       post.UpdatedAt().Format(time.RFC3339),
	}
}

type SocialPostsListResponse struct {
	Posts      []SocialPostResponse `json:"posts"`
	Total      int64                `json:"total"`
	Page       int                  `json:"page"`
	PageSize   int                  `json:"page_size"`
	TotalPages int                  `json:"total_pages"`
}

type SocialDailyAggregationResponse struct {
	ID              string   `json:"id"`
	BrandName       string   `json:"brand_name"`
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

// SocialDailyAggregationToResponse converts a SocialDailyAggregation entity to a SocialDailyAggregationResponse
func SocialDailyAggregationToResponse(agg *entity.SocialDailyAggregation) SocialDailyAggregationResponse {
	return SocialDailyAggregationResponse{
		ID:              agg.ID().String(),
		BrandName:       agg.BrandName().Value(),
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

type SocialDailyAggregationsListResponse struct {
	Aggregations []SocialDailyAggregationResponse `json:"aggregations"`
	Total        int64                            `json:"total"`
	Page         int                              `json:"page"`
	PageSize     int                              `json:"page_size"`
	TotalPages   int                              `json:"total_pages"`
}
