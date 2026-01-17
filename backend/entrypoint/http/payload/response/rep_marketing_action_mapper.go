package response

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

type RepMarketingActionPayloadMapper struct{}

func NewRepMarketingActionPayloadMapper() *RepMarketingActionPayloadMapper {
	return &RepMarketingActionPayloadMapper{}
}

// ToRepMarketingActionResponse converte Entity para Response DTO
func (m *RepMarketingActionPayloadMapper) ToRepMarketingActionResponse(action *entity.RepMarketingAction) RepMarketingActionResponse {
	response := RepMarketingActionResponse{
		UUID:               action.ID().String(),
		RepresentativeUUID: action.RepresentativeUUID().String(),
		Date:               action.Date().Format(time.RFC3339),
		Description:        action.Description(),
		Month:              action.Month(),
		CreatedAt:          action.CreatedAt().Format(time.RFC3339),
		UpdatedAt:          action.UpdatedAt().Format(time.RFC3339),
	}

	// Optional field
	if action.DeletedAt() != nil {
		deletedAt := action.DeletedAt().Format(time.RFC3339)
		response.DeletedAt = &deletedAt
	}

	return response
}

// ToRepMarketingActionResponseList converte slice de Entity para slice de Response DTO
func (m *RepMarketingActionPayloadMapper) ToRepMarketingActionResponseList(actions []*entity.RepMarketingAction) []RepMarketingActionResponse {
	responses := make([]RepMarketingActionResponse, len(actions))
	for i, action := range actions {
		responses[i] = m.ToRepMarketingActionResponse(action)
	}
	return responses
}

// ToRepMarketingActionsListResponse cria a resposta paginada de actions
func (m *RepMarketingActionPayloadMapper) ToRepMarketingActionsListResponse(
	actions []*entity.RepMarketingAction,
	total int64,
	page int,
	limit int,
) map[string]interface{} {
	return map[string]interface{}{
		"data": m.ToRepMarketingActionResponseList(actions),
		"pagination": PaginationResponse{
			Page:       page,
			PageSize:   limit,
			Total:      total,
			TotalPages: int((total + int64(limit) - 1) / int64(limit)),
		},
	}
}
