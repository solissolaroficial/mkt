package social

import (
	"context"
	stderrors "errors"
	"time"

	"github.com/google/uuid"
	socialerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
)

// UpdateSocialPostUseCase handles updating a social post
type UpdateSocialPostUseCase struct {
	postGateway                    gateway.SocialPostGateway
	recalculateAggregationsUseCase *RecalculateDailyAggregationsUseCase
}

// NewUpdateSocialPostUseCase creates a new update social post use case
func NewUpdateSocialPostUseCase(
	postGateway gateway.SocialPostGateway,
	recalculateAggregationsUseCase *RecalculateDailyAggregationsUseCase,
) *UpdateSocialPostUseCase {
	return &UpdateSocialPostUseCase{
		postGateway:                    postGateway,
		recalculateAggregationsUseCase: recalculateAggregationsUseCase,
	}
}

// UpdateSocialPostInput represents the input for updating a social post
type UpdateSocialPostInput struct {
	ID              string  `json:"id"`
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

// UpdateSocialPostOutput represents the output of updating a social post
type UpdateSocialPostOutput struct {
	ID              string  `json:"id"`
	BrandName       string  `json:"brand_name"`
	Platform        string  `json:"platform"`
	PostDate        string  `json:"post_date"`
	PostTime        *string `json:"post_time,omitempty"`
	Likes           int     `json:"likes"`
	Comments        int     `json:"comments"`
	Shares          *int    `json:"shares,omitempty"`
	PostType        string  `json:"post_type"`
	Caption         *string `json:"caption,omitempty"`
	FollowersAtPost *int    `json:"followers_at_post,omitempty"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

// Execute updates a social post
func (uc *UpdateSocialPostUseCase) Execute(ctx context.Context, input UpdateSocialPostInput) (*UpdateSocialPostOutput, error) {
	if input.ID == "" {
		return nil, socialerrors.ErrSocialPostIDRequired
	}

	// Parse UUID
	id, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, socialerrors.ErrSocialPostNotFound
	}

	// Get the existing post
	post, err := uc.postGateway.GetByID(id)
	if err != nil {
		if stderrors.Is(err, socialerrors.ErrSocialPostNotFound) {
			return nil, err
		}
		return nil, err
	}

	// Update fields if provided
	if input.Platform != nil {
		platform, err := valueobject.NewSocialPlatform(*input.Platform)
		if err != nil {
			return nil, err
		}
		if err := post.UpdatePlatform(platform); err != nil {
			return nil, err
		}
	}

	if input.PostDate != nil {
		postDate, err := time.Parse("2006-01-02", *input.PostDate)
		if err != nil {
			return nil, stderrors.New("invalid post_date format, expected YYYY-MM-DD")
		}
		if err := post.UpdatePostDate(postDate); err != nil {
			return nil, err
		}
	}

	if input.PostTime != nil {
		postTime, err := time.Parse("15:04:05", *input.PostTime)
		if err != nil {
			return nil, stderrors.New("invalid post_time format, expected HH:MM:SS")
		}
		post.SetPostTime(&postTime)
	}

	if input.Likes != nil {
		if err := post.UpdateMetrics(*input.Likes, post.Comments(), post.Shares()); err != nil {
			return nil, err
		}
	}

	if input.Comments != nil {
		if err := post.UpdateMetrics(post.Likes(), *input.Comments, post.Shares()); err != nil {
			return nil, err
		}
	}

	if input.Shares != nil {
		if err := post.UpdateMetrics(post.Likes(), post.Comments(), input.Shares); err != nil {
			return nil, err
		}
	}

	if input.PostType != nil {
		postType, err := valueobject.NewSocialPostType(*input.PostType)
		if err != nil {
			return nil, err
		}
		post.SetPostType(postType)
	}

	if input.Caption != nil {
		post.UpdateCaption(input.Caption)
	}

	if input.FollowersAtPost != nil {
		post.SetFollowersAtPost(*input.FollowersAtPost)
	}

	// Validate the updated post
	if err := post.Validate(); err != nil {
		return nil, err
	}

	// Update in database
	if err := uc.postGateway.Update(post); err != nil {
		return nil, err
	}

	// Recalculate daily aggregations
	_, err = uc.recalculateAggregationsUseCase.Execute(post.BrandName().Value(), post.PostDate())
	if err != nil {
		// Log error but don't fail the update
		// In production, this should be asynchronous
		// We still return the updated post even if recalculation fails
	}

	// Format post time for output
	var postTimeStr *string
	if post.PostTime() != nil {
		formatted := post.PostTime().Format("15:04:05")
		postTimeStr = &formatted
	}

	return &UpdateSocialPostOutput{
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
		CreatedAt:       post.CreatedAt().Format(time.RFC3339),
		UpdatedAt:       post.UpdatedAt().Format(time.RFC3339),
	}, nil
}
