package social

import (
	"context"
	"time"

	"github.com/google/uuid"
	socialerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetSocialPostUseCase handles getting a social post by ID
type GetSocialPostUseCase struct {
	postGateway gateway.SocialPostGateway
}

// NewGetSocialPostUseCase creates a new get social post use case
func NewGetSocialPostUseCase(postGateway gateway.SocialPostGateway) *GetSocialPostUseCase {
	return &GetSocialPostUseCase{
		postGateway: postGateway,
	}
}

// GetSocialPostOutput represents the output of getting a social post
type GetSocialPostOutput struct {
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

// Execute gets a social post by ID
func (uc *GetSocialPostUseCase) Execute(ctx context.Context, id string) (*GetSocialPostOutput, error) {
	if id == "" {
		return nil, socialerrors.ErrSocialPostIDRequired
	}

	// Parse UUID
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, socialerrors.ErrSocialPostNotFound
	}

	// Get the post
	post, err := uc.postGateway.GetByID(uuidID)
	if err != nil {
		if err == socialerrors.ErrSocialPostNotFound {
			return nil, err
		}
		return nil, err
	}

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

	return &GetSocialPostOutput{
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
	}, nil
}
