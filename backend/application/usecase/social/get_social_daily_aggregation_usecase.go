package social

import (
	"context"
	"time"

	"github.com/google/uuid"
	socialerrors "github.com/seu-usuario/solis-backend/core/domain/errors"
	"github.com/seu-usuario/solis-backend/core/domain/gateway"
)

// GetSocialDailyAggregationUseCase handles getting a social daily aggregation by ID
type GetSocialDailyAggregationUseCase struct {
	aggregationGateway gateway.SocialDailyAggregationGateway
}

// NewGetSocialDailyAggregationUseCase creates a new get social daily aggregation use case
func NewGetSocialDailyAggregationUseCase(aggregationGateway gateway.SocialDailyAggregationGateway) *GetSocialDailyAggregationUseCase {
	return &GetSocialDailyAggregationUseCase{
		aggregationGateway: aggregationGateway,
	}
}

// GetSocialDailyAggregationOutput represents the output of getting a social daily aggregation
type GetSocialDailyAggregationOutput struct {
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
	EngagementRate  *float64 `json:"engagement_rate,omitempty"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
}

// Execute gets a social daily aggregation by ID
func (uc *GetSocialDailyAggregationUseCase) Execute(ctx context.Context, id string) (*GetSocialDailyAggregationOutput, error) {
	if id == "" {
		return nil, socialerrors.ErrSocialDailyAggregationIDRequired
	}

	// Parse UUID
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, socialerrors.ErrSocialDailyAggregationNotFound
	}

	// Get aggregation
	aggregation, err := uc.aggregationGateway.GetByID(uuidID)
	if err != nil {
		if err == socialerrors.ErrSocialDailyAggregationNotFound {
			return nil, err
		}
		return nil, err
	}

	// Get engagement rate value
	var engagementRate *float64
	if aggregation.EngagementRate() != nil {
		value := aggregation.EngagementRate().Value()
		engagementRate = &value
	}

	return &GetSocialDailyAggregationOutput{
		ID:              aggregation.ID().String(),
		BrandName:       aggregation.BrandName().Value(),
		AggregationDate: aggregation.AggregationDate().Format("2006-01-02"),
		TotalPosts:      aggregation.TotalPosts(),
		TotalLikes:      aggregation.TotalLikes(),
		TotalComments:   aggregation.TotalComments(),
		TotalShares:     aggregation.TotalShares(),
		AvgLikes:        aggregation.AvgLikes(),
		AvgComments:     aggregation.AvgComments(),
		AvgShares:       aggregation.AvgShares(),
		FollowersAtDate: aggregation.FollowersAtDate(),
		EngagementRate:  engagementRate,
		CreatedAt:       aggregation.CreatedAt().Format(time.RFC3339),
		UpdatedAt:       aggregation.UpdatedAt().Format(time.RFC3339),
	}, nil
}
