package mapper

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

// ModelToEntity converte model para entity
func ModelToSocialPostEntity(m *model.SocialPostModel) *entity.SocialPost {
	return entity.ReconstructSocialPost(
		m.ID,
		m.BrandName,
		m.Platform,
		m.PostDate,
		m.PostTime,
		m.Likes,
		m.Comments,
		m.Shares,
		m.PostType,
		m.Caption,
		m.FollowersAtPost,
		m.CreatedAt,
		m.UpdatedAt,
		m.DeletedAt,
	)
}

// EntityToModel converte entity para model
func EntityToSocialPostModel(e *entity.SocialPost) *model.SocialPostModel {
	return &model.SocialPostModel{
		ID:              e.ID(),
		BrandName:       e.BrandName().String(),
		Platform:        e.Platform().String(),
		PostDate:        e.PostDate(),
		PostTime:        e.PostTime(),
		Likes:           e.Likes(),
		Comments:        e.Comments(),
		Shares:          e.Shares(),
		PostType:        e.PostType().String(),
		Caption:         e.Caption(),
		FollowersAtPost: e.FollowersAtPost(),
		CreatedAt:       e.CreatedAt(),
		UpdatedAt:       e.UpdatedAt(),
		DeletedAt:       e.DeletedAt(),
	}
}
