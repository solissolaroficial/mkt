package response

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

type SocialPayloadMapper struct{}

func NewSocialPayloadMapper() *SocialPayloadMapper {
	return &SocialPayloadMapper{}
}

// ToSocialBenchmarkingResponse converte Entity para Response DTO
func (m *SocialPayloadMapper) ToSocialBenchmarkingResponse(benchmarking *entity.SocialBenchmarking) SocialBenchmarkingResponse {
	return SocialBenchmarkingResponse{
		ID:             benchmarking.ID().String(),
		BrandName:      benchmarking.BrandName().String(),
		AvgLikes:       benchmarking.AvgLikes(),
		AvgComments:    benchmarking.AvgComments(),
		Followers:      benchmarking.Followers(),
		EngagementRate: benchmarking.EngagementRate().Value(),
		CreatedAt:      benchmarking.CreatedAt().Format(time.RFC3339),
		UpdatedAt:      benchmarking.UpdatedAt().Format(time.RFC3339),
	}
}

// ToSocialBenchmarkingResponseList converte slice de Entity para slice de Response DTO
func (m *SocialPayloadMapper) ToSocialBenchmarkingResponseList(benchmarkings []*entity.SocialBenchmarking) []SocialBenchmarkingResponse {
	responses := make([]SocialBenchmarkingResponse, len(benchmarkings))
	for i, benchmarking := range benchmarkings {
		responses[i] = m.ToSocialBenchmarkingResponse(benchmarking)
	}
	return responses
}

// ToSocialBenchmarkingsListResponse cria a resposta paginada de benchmarkings
func (m *SocialPayloadMapper) ToSocialBenchmarkingsListResponse(
	benchmarkings []*entity.SocialBenchmarking,
	total int64,
	page int,
	limit int,
) SocialBenchmarkingListData {
	return SocialBenchmarkingListData{
		Benchmarkings: m.ToSocialBenchmarkingResponseList(benchmarkings),
		Meta: MetaResponse{
			Total:      total,
			Page:       page,
			Limit:      limit,
			TotalPages: int((total + int64(limit) - 1) / int64(limit)),
		},
	}
}
