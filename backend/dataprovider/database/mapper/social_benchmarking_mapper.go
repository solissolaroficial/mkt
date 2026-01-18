package mapper

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/dataprovider/database/model"
	"gorm.io/gorm"
)

type SocialBenchmarkingMapper struct{}

func NewSocialBenchmarkingMapper() *SocialBenchmarkingMapper {
	return &SocialBenchmarkingMapper{}
}

// ModelToEntity converte Model para Entity
func (m *SocialBenchmarkingMapper) ModelToEntity(model *model.SocialBenchmarkingModel) (*entity.SocialBenchmarking, error) {
	// Converter deletedAt
	var deletedAt *time.Time
	if model.DeletedAt.Valid {
		deletedAt = &model.DeletedAt.Time
	}

	// Get brand name from Brand relationship if available
	var brandName string
	if model.Brand != nil {
		brandName = model.Brand.Name
	} else {
		brandName = "Unknown" // Fallback if brand not loaded
	}

	return entity.ReconstructSocialBenchmarking(
		model.UUID,
		model.BrandID,
		brandName,
		model.AvgLikes,
		model.AvgComments,
		model.Followers,
		model.EngagementRate,
		model.CreatedAt,
		model.UpdatedAt,
		deletedAt,
	), nil
}

// ModelsToEntities converte slice de Model para slice de Entity
func (m *SocialBenchmarkingMapper) ModelsToEntities(models []*model.SocialBenchmarkingModel) ([]*entity.SocialBenchmarking, error) {
	benchmarkings := make([]*entity.SocialBenchmarking, len(models))
	for i, model := range models {
		benchmarking, err := m.ModelToEntity(model)
		if err != nil {
			return nil, err
		}
		benchmarkings[i] = benchmarking
	}
	return benchmarkings, nil
}

// EntityToModel converte Entity para Model
func (m *SocialBenchmarkingMapper) EntityToModel(benchmarking *entity.SocialBenchmarking) *model.SocialBenchmarkingModel {
	// Converter deletedAt
	var deletedAt gorm.DeletedAt
	if benchmarking.DeletedAt() != nil {
		deletedAt.Time = *benchmarking.DeletedAt()
		deletedAt.Valid = true
	}

	return &model.SocialBenchmarkingModel{
		UUID:           benchmarking.ID(),
		BrandID:        benchmarking.BrandID(),
		AvgLikes:       benchmarking.AvgLikes(),
		AvgComments:    benchmarking.AvgComments(),
		Followers:      benchmarking.Followers(),
		EngagementRate: benchmarking.EngagementRate().Value(),
		CreatedAt:      benchmarking.CreatedAt(),
		UpdatedAt:      benchmarking.UpdatedAt(),
		DeletedAt:      deletedAt,
	}
}
