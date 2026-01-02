package mapper

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

// CommentMapper converte entre Comment entity e Comment model
type CommentMapper struct{}

// NewCommentMapper cria um novo CommentMapper
func NewCommentMapper() *CommentMapper {
	return &CommentMapper{}
}

// ToEntity converte um Comment model para Comment entity
func (m *CommentMapper) ToEntity(model *model.Comment) *entity.Comment {
	return entity.ReconstructComment(
		model.UUID,
		model.TaskID,
		model.UserID,
		model.Text,
		model.Timestamp,
	)
}

// ToModel converte uma Comment entity para Comment model
func (m *CommentMapper) ToModel(entity *entity.Comment) *model.Comment {
	return &model.Comment{
		UUID:      entity.ID(),
		TaskID:    entity.TaskID(),
		UserID:    entity.UserID(),
		Text:      entity.Text(),
		Timestamp: entity.Timestamp(),
	}
}

// ToEntityList converte uma lista de Comment models para Comment entities
func (m *CommentMapper) ToEntityList(models []*model.Comment) []*entity.Comment {
	entities := make([]*entity.Comment, 0, len(models))
	for _, model := range models {
		entities = append(entities, m.ToEntity(model))
	}
	return entities
}
