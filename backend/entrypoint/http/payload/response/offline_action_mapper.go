package response

import (
	"time"

	"github.com/seu-usuario/solis-backend/core/domain/entity"
)

type OfflineActionPayloadMapper struct{}

func NewOfflineActionPayloadMapper() *OfflineActionPayloadMapper {
	return &OfflineActionPayloadMapper{}
}

// ToOfflineActionResponse converte Entity para Response DTO
func (m *OfflineActionPayloadMapper) ToOfflineActionResponse(action *entity.OfflineAction) OfflineActionResponse {
	pdv := ""
	if v := action.PDV(); v != nil {
		pdv = *v
	}

	response := OfflineActionResponse{
		UUID:            action.ID().String(),
		RequestedAmount: action.RequestedAmount().Value(),
		ActionDate:      action.ActionDate().Value().Format(time.RFC3339),
		Category:        action.Category().String(),
		PDV:             pdv,
		RepName:         action.RepName(),
		Observation:     action.Observation(),
		Status:          action.Status().String(),
		Month:           action.Month(),
		CreatedAt:       action.CreatedAt().Format(time.RFC3339),
		UpdatedAt:       action.UpdatedAt().Format(time.RFC3339),
	}

	// Optional fields
	if action.ApprovedAmount() != nil {
		response.ApprovedAmount = action.ApprovedAmount()
	}

	if action.OrderNumber() != nil {
		response.OrderNumber = action.OrderNumber()
	}

	if action.DepartureDate() != nil {
		response.DepartureDate = action.DepartureDate()
	}

	if action.DeliveryForecast() != nil {
		response.DeliveryForecast = action.DeliveryForecast()
	}

	if action.DeliveryDate() != nil {
		response.DeliveryDate = action.DeliveryDate()
	}

	if action.City() != nil {
		response.City = action.City()
	}

	if action.UF() != nil {
		response.UF = action.UF()
	}

	if action.Scored() != nil {
		scored := action.Scored().String()
		response.Scored = &scored
	}

	if action.DeletedAt() != nil {
		deletedAt := action.DeletedAt().Format(time.RFC3339)
		response.DeletedAt = &deletedAt
	}

	return response
}

// ToOfflineActionResponseList converte slice de Entity para slice de Response DTO
func (m *OfflineActionPayloadMapper) ToOfflineActionResponseList(actions []*entity.OfflineAction) []OfflineActionResponse {
	responses := make([]OfflineActionResponse, len(actions))
	for i, action := range actions {
		responses[i] = m.ToOfflineActionResponse(action)
	}
	return responses
}

// ToOfflineActionsListResponse cria a resposta paginada de actions
func (m *OfflineActionPayloadMapper) ToOfflineActionsListResponse(
	actions []*entity.OfflineAction,
	total int64,
	page int,
	limit int,
) map[string]interface{} {
	return map[string]interface{}{
		"data": m.ToOfflineActionResponseList(actions),
		"pagination": PaginationResponse{
			Page:       page,
			PageSize:   limit,
			Total:      total,
			TotalPages: int((total + int64(limit) - 1) / int64(limit)),
		},
	}
}
