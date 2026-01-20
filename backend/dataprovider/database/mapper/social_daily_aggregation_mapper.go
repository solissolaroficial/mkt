package mapper

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

// ModelToEntity converte model para entity
func ModelToSocialDailyAggregationEntity(m *model.SocialDailyAggregationModel) *entity.SocialDailyAggregation {
	return entity.ReconstructSocialDailyAggregation(
		m.ID,
		m.BrandID,
		m.AggregationDate,
		m.TotalPosts,
		m.TotalLikes,
		m.TotalComments,
		m.TotalShares,
		m.AvgLikes,
		m.AvgComments,
		m.AvgShares,
		m.FollowersAtDate,
		m.EngagementRate,
		m.CreatedAt,
		m.UpdatedAt,
		m.DeletedAt,
	)
}

// EntityToModel converte entity para model
func EntityToSocialDailyAggregationModel(e *entity.SocialDailyAggregation) *model.SocialDailyAggregationModel {
	return &model.SocialDailyAggregationModel{
		ID:              e.ID(),
		BrandID:         e.BrandID(),
		AggregationDate: e.AggregationDate(),
		TotalPosts:      e.TotalPosts(),
		TotalLikes:      e.TotalLikes(),
		TotalComments:   e.TotalComments(),
		TotalShares:     e.TotalShares(),
		AvgLikes:        e.AvgLikes(),
		AvgComments:     e.AvgComments(),
		AvgShares:       e.AvgShares(),
		FollowersAtDate: e.FollowersAtDate(),
		EngagementRate:  e.EngagementRate().Value(),
		CreatedAt:       e.CreatedAt(),
		UpdatedAt:       e.UpdatedAt(),
		DeletedAt:       e.DeletedAt(),
	}
}
