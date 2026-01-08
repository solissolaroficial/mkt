package mapper

import (
	"time"

	"gorm.io/gorm"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/core/domain/valueobject"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
)

type PdvPostMapper struct{}

func NewPdvPostMapper() *PdvPostMapper {
	return &PdvPostMapper{}
}

// ModelToEntity converte Model para Entity
func (m *PdvPostMapper) ModelToEntity(model *model.PdvPostModel) (*entity.PdvPost, error) {
	// Converter data usando ReconstructPdvPostDate (assume dados do banco são válidos)
	postDate := valueobject.ReconstructPdvPostDate(model.PostDate.Format("2006-01-02"))

	// Converter status
	status, err := valueobject.NewPdvStatus(model.Status)
	if err != nil {
		return nil, err
	}

	// Converter deletedAt
	var deletedAt *time.Time
	if model.DeletedAt.Valid {
		deletedAt = &model.DeletedAt.Time
	}

	return entity.ReconstructPdvPost(
		model.UUID,
		model.RepName,
		model.PdvName,
		postDate,
		model.Month,
		model.Platform,
		model.Link,
		model.ProofUrl,
		status,
		model.CreatedAt,
		model.UpdatedAt,
		deletedAt,
	)
}

// ModelsToEntities converte slice de Model para slice de Entity
func (m *PdvPostMapper) ModelsToEntities(models []*model.PdvPostModel) ([]*entity.PdvPost, error) {
	posts := make([]*entity.PdvPost, len(models))
	for i, model := range models {
		post, err := m.ModelToEntity(model)
		if err != nil {
			return nil, err
		}
		posts[i] = post
	}
	return posts, nil
}

// EntityToModel converte Entity para Model
func (m *PdvPostMapper) EntityToModel(post *entity.PdvPost) *model.PdvPostModel {
	// Converter deletedAt
	var deletedAt gorm.DeletedAt
	if post.DeletedAt() != nil {
		deletedAt.Time = *post.DeletedAt()
		deletedAt.Valid = true
	}

	// Obter status string
	var statusStr string
	if post.Status() != nil {
		statusStr = post.Status().Value()
	} else {
		statusStr = valueobject.PdvStatusPending
	}

	return &model.PdvPostModel{
		UUID:      post.ID(),
		RepName:   post.RepName(),
		PdvName:   post.PdvName(),
		PostDate:  post.PostDate().Value(),
		Month:     post.Month(),
		Platform:  post.Platform(),
		Link:      post.Link(),
		ProofUrl:  post.ProofUrl(),
		Status:    statusStr,
		CreatedAt: post.CreatedAt(),
		UpdatedAt: post.UpdatedAt(),
		DeletedAt: deletedAt,
	}
}
