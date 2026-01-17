package response

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

type ShowroomItemPayloadMapper struct{}

func NewShowroomItemPayloadMapper() *ShowroomItemPayloadMapper {
	return &ShowroomItemPayloadMapper{}
}

// ToShowroomItemResponse converte Entity para Response DTO
func (m *ShowroomItemPayloadMapper) ToShowroomItemResponse(item *entity.ShowroomItem) ShowroomItemResponse {
	response := ShowroomItemResponse{
		UUID:               item.ID().String(),
		PDV:                item.PDV(),
		City:               item.City(),
		Contact:            item.Contact(),
		RepresentativeUUID: item.RepresentativeUUID().String(),
		DeliveryForecast:   item.DeliveryForecast(),
		Delivered:          item.Delivered(),
		CreatedAt:          item.CreatedAt().Format(time.RFC3339),
		UpdatedAt:          item.UpdatedAt().Format(time.RFC3339),
	}

	// Optional fields
	if item.WorkshopDate() != nil {
		response.WorkshopDate = item.WorkshopDate()
	}

	if item.DeletedAt() != nil {
		deletedAt := item.DeletedAt().Format(time.RFC3339)
		response.DeletedAt = &deletedAt
	}

	return response
}

// ToShowroomItemResponseList converte slice de Entity para slice de Response DTO
func (m *ShowroomItemPayloadMapper) ToShowroomItemResponseList(items []*entity.ShowroomItem) []ShowroomItemResponse {
	responses := make([]ShowroomItemResponse, len(items))
	for i, item := range items {
		responses[i] = m.ToShowroomItemResponse(item)
	}
	return responses
}

// ToShowroomItemsListResponse cria a resposta paginada de items
func (m *ShowroomItemPayloadMapper) ToShowroomItemsListResponse(
	items []*entity.ShowroomItem,
	total int64,
	page int,
	limit int,
) map[string]interface{} {
	return map[string]interface{}{
		"data": m.ToShowroomItemResponseList(items),
		"pagination": PaginationResponse{
			Page:       page,
			PageSize:   limit,
			Total:      total,
			TotalPages: int((total + int64(limit) - 1) / int64(limit)),
		},
	}
}
