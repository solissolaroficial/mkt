package mapper

import (
	"time"

	"github.com/google/uuid"
	"github.com/seu-usuario/solis-backend/application/usecase/representatives"
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// RepresentativePayloadMapper converts between DTOs and Use Case Inputs/Outputs
type RepresentativePayloadMapper struct{}

func NewRepresentativePayloadMapper() *RepresentativePayloadMapper {
	return &RepresentativePayloadMapper{}
}

// CreateRequestToInput converts CreateRepresentativeRequest to CreateRepresentativeInput
func (m *RepresentativePayloadMapper) CreateRequestToInput(req *request.CreateRepresentativeRequest) representatives.CreateRepresentativeInput {
	return representatives.CreateRepresentativeInput{
		Code:      req.Code,
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Company:   req.Company,
		Region:    req.Region,
		City:      req.City,
		Attendant: req.Attendant,
	}
}

// UpdateRequestToInput converts UpdateRepresentativeRequest to UpdateRepresentativeInput
func (m *RepresentativePayloadMapper) UpdateRequestToInput(idStr string, req *request.UpdateRepresentativeRequest) (representatives.UpdateRepresentativeInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return representatives.UpdateRepresentativeInput{}, err
	}
	return representatives.UpdateRepresentativeInput{
		ID:        id,
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Company:   req.Company,
		Region:    req.Region,
		City:      req.City,
		Attendant: req.Attendant,
		Active:    req.Active,
	}, nil
}

// GetRequestToInput converts ID string to GetRepresentativeInput
func (m *RepresentativePayloadMapper) GetRequestToInput(idStr string) (representatives.GetRepresentativeInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return representatives.GetRepresentativeInput{}, err
	}
	return representatives.GetRepresentativeInput{
		ID: id,
	}, nil
}

// DeleteRequestToInput converts ID string to DeleteRepresentativeInput
func (m *RepresentativePayloadMapper) DeleteRequestToInput(idStr string) (representatives.DeleteRepresentativeInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return representatives.DeleteRepresentativeInput{}, err
	}
	return representatives.DeleteRepresentativeInput{
		ID: id,
	}, nil
}

// GetStatsRequestToInput converts ID string to GetRepresentativeStatsInput
func (m *RepresentativePayloadMapper) GetStatsRequestToInput(idStr string) (representatives.GetRepresentativeStatsInput, error) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return representatives.GetRepresentativeStatsInput{}, err
	}
	return representatives.GetRepresentativeStatsInput{
		ID: id,
	}, nil
}

// ListRequestToInput converts ListRepresentativesRequest to ListRepresentativesInput
func (m *RepresentativePayloadMapper) ListRequestToInput(req *request.ListRepresentativesRequest) representatives.ListRepresentativesInput {
	return representatives.ListRepresentativesInput{
		Page:      req.Page,
		PageSize:  req.PageSize,
		Name:      req.Name,
		Company:   req.Company,
		Email:     req.Email,
		Region:    req.Region,
		City:      req.City,
		Active:    req.Active,
		Code:      req.Code,
		SortBy:    req.SortBy,
		SortOrder: req.SortOrder,
	}
}

// CreateOutputToResponse converts CreateRepresentativeOutput to RepresentativeResponse
func (m *RepresentativePayloadMapper) CreateOutputToResponse(output *representatives.CreateRepresentativeOutput) *response.RepresentativeResponse {
	return &response.RepresentativeResponse{
		UUID:      output.UUID,
		Code:      output.Code,
		Name:      output.Name,
		Email:     output.Email,
		Phone:     output.Phone,
		Company:   output.Company,
		Region:    output.Region,
		City:      output.City,
		Attendant: output.Attendant,
		Active:    output.Active,
		CreatedAt: parseTime(output.CreatedAt),
		UpdatedAt: parseTime(output.UpdatedAt),
	}
}

// GetOutputToResponse converts GetRepresentativeOutput to RepresentativeResponse
func (m *RepresentativePayloadMapper) GetOutputToResponse(output *representatives.GetRepresentativeOutput) *response.RepresentativeResponse {
	return &response.RepresentativeResponse{
		UUID:      output.UUID,
		Code:      output.Code,
		Name:      output.Name,
		Email:     output.Email,
		Phone:     output.Phone,
		Company:   output.Company,
		Region:    output.Region,
		City:      output.City,
		Attendant: output.Attendant,
		Active:    output.Active,
		CreatedAt: parseTime(output.CreatedAt),
		UpdatedAt: parseTime(output.UpdatedAt),
		DeletedAt: parseTimePtr(output.DeletedAt),
	}
}

// UpdateOutputToResponse converts UpdateRepresentativeOutput to RepresentativeResponse
func (m *RepresentativePayloadMapper) UpdateOutputToResponse(output *representatives.UpdateRepresentativeOutput) *response.RepresentativeResponse {
	return &response.RepresentativeResponse{
		UUID:      output.UUID,
		Code:      output.Code,
		Name:      output.Name,
		Email:     output.Email,
		Phone:     output.Phone,
		Company:   output.Company,
		Region:    output.Region,
		City:      output.City,
		Attendant: output.Attendant,
		Active:    output.Active,
		CreatedAt: parseTime(output.CreatedAt),
		UpdatedAt: parseTime(output.UpdatedAt),
	}
}

// StatsOutputToResponse converts GetRepresentativeStatsOutput to RepresentativeStatsResponse
func (m *RepresentativePayloadMapper) StatsOutputToResponse(output *representatives.GetRepresentativeStatsOutput) *response.RepresentativeStatsResponse {
	return &response.RepresentativeStatsResponse{
		UUID:               output.UUID,
		OnlineCount:        output.OnlineCount,
		OnlineActionCount:  output.OnlineActionCount,
		OfflineCount:       output.OfflineCount,
		OfflineActionCount: output.OfflineActionCount,
		OfflineValue:       output.OfflineValue,
		OfflineActionValue: output.OfflineActionValue,
		ShowroomItemCount:  output.ShowroomItemCount,
		RepMarketingCount:  output.RepMarketingCount,
		TotalActions:       output.TotalActions,
	}
}

// ListOutputToResponse converts ListRepresentativesOutput to ListRepresentativesResponse
func (m *RepresentativePayloadMapper) ListOutputToResponse(output *representatives.ListRepresentativesOutput) *response.ListRepresentativesResponse {
	responses := make([]response.RepresentativeResponse, len(output.Data))
	for i, entity := range output.Data {
		responses[i] = entityToResponse(entity)
	}
	return &response.ListRepresentativesResponse{
		Data:       responses,
		Total:      output.Total,
		Page:       output.Page,
		PageSize:   output.PageSize,
		TotalPages: output.TotalPages,
	}
}

// entityToResponse converts a Representative entity to RepresentativeResponse
func entityToResponse(entity *entity.Representative) response.RepresentativeResponse {
	return response.RepresentativeResponse{
		UUID:      entity.UUID(),
		Code:      entity.Code().Value(),
		Name:      entity.Name(),
		Email:     entity.Email(),
		Phone:     entity.Phone(),
		Company:   entity.Company(),
		Region:    entity.Region(),
		City:      entity.City(),
		Attendant: entity.Attendant(),
		Active:    entity.Active(),
		CreatedAt: entity.CreatedAt(),
		UpdatedAt: entity.UpdatedAt(),
		DeletedAt: entity.DeletedAt(),
	}
}

// parseTime is a helper function to parse time string
func parseTime(timeStr string) time.Time {
	t, _ := time.Parse(time.RFC3339, timeStr)
	return t
}

// parseTimePtr is a helper function to parse time pointer string
func parseTimePtr(timeStr *string) *time.Time {
	if timeStr == nil {
		return nil
	}
	t := parseTime(*timeStr)
	return &t
}

// ProfileOutputToResponse converts GetRepresentativeProfileOutput to RepresentativeProfileResponse
func (m *RepresentativePayloadMapper) ProfileOutputToResponse(output *representatives.GetRepresentativeProfileOutput) *response.RepresentativeProfileResponse {
	return &response.RepresentativeProfileResponse{
		UUID:              output.UUID,
		Code:              output.Code,
		Name:              output.Name,
		Email:             output.Email,
		Phone:             output.Phone,
		Company:           output.Company,
		Region:            output.Region,
		City:              output.City,
		Attendant:         output.Attendant,
		Active:            output.Active,
		CreatedAt:         output.CreatedAt,
		UpdatedAt:         output.UpdatedAt,
		TrainingCount:     output.TrainingCount,
		OnlineCount:       output.OnlineCount,
		OfflineCount:      output.OfflineCount,
		OfflineValue:      output.OfflineValue,
		ShowroomItemCount: output.ShowroomItemCount,
		RepMarketingCount: output.RepMarketingCount,
		TotalActions:      output.TotalActions,
	}
}

// AllProfilesOutputToResponse converts GetAllRepresentativeProfilesOutput to GetAllRepresentativeProfilesResponse
func (m *RepresentativePayloadMapper) AllProfilesOutputToResponse(output []*representatives.GetAllRepresentativeProfilesOutput) *response.GetAllRepresentativeProfilesResponse {
	data := make([]response.RepresentativeProfileResponse, len(output))
	for i, profile := range output {
		data[i] = response.RepresentativeProfileResponse{
			UUID:              profile.UUID,
			Code:              profile.Code,
			Name:              profile.Name,
			Email:             profile.Email,
			Phone:             profile.Phone,
			Company:           profile.Company,
			Region:            profile.Region,
			City:              profile.City,
			Attendant:         profile.Attendant,
			Active:            profile.Active,
			CreatedAt:         profile.CreatedAt,
			UpdatedAt:         profile.UpdatedAt,
			TrainingCount:     profile.TrainingCount,
			OnlineCount:       profile.OnlineCount,
			OfflineCount:      profile.OfflineCount,
			OfflineValue:      profile.OfflineValue,
			ShowroomItemCount: profile.ShowroomItemCount,
			RepMarketingCount: profile.RepMarketingCount,
			TotalActions:      profile.TotalActions,
		}
	}
	return &response.GetAllRepresentativeProfilesResponse{
		Data: data,
	}
}
