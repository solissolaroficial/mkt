package mapper

import (
	"github.com/google/uuid"

	"github.com/seu-usuario/solis-backend/application/usecase/representativemonthlygoal"
	representativemonthlygoalrequest "github.com/seu-usuario/solis-backend/entrypoint/http/payload/request"
	representativemonthlygoalresponse "github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

type RepresentativeMonthlyGoalPayloadMapper struct{}

func NewRepresentativeMonthlyGoalPayloadMapper() *RepresentativeMonthlyGoalPayloadMapper {
	return &RepresentativeMonthlyGoalPayloadMapper{}
}

// CreateRequestToInput converts request to use case input
func (m *RepresentativeMonthlyGoalPayloadMapper) CreateRequestToInput(req *representativemonthlygoalrequest.CreateRepresentativeMonthlyGoalRequest) *representativemonthlygoal.CreateRepresentativeMonthlyGoalInput {
	representativeID, _ := uuid.Parse(req.RepresentativeID)
	return &representativemonthlygoal.CreateRepresentativeMonthlyGoalInput{
		RepresentativeUUID: representativeID,
		Month:              req.Month,
		Year:               req.Year,
		Target:             req.Target,
	}
}

// CreateOutputToResponse converts use case output to response
func (m *RepresentativeMonthlyGoalPayloadMapper) CreateOutputToResponse(output *representativemonthlygoal.CreateRepresentativeMonthlyGoalOutput) *representativemonthlygoalresponse.CreateRepresentativeMonthlyGoalResponse {
	return &representativemonthlygoalresponse.CreateRepresentativeMonthlyGoalResponse{
		ID:                 output.ID,
		RepresentativeUUID: output.RepresentativeUUID,
		Month:              output.Month,
		Year:               output.Year,
		Target:             output.Target,
		Realized:           output.Realized,
		Percentage:         output.Percentage,
		CreatedAt:          output.CreatedAt,
		UpdatedAt:          output.UpdatedAt,
	}
}

// GetRequestToInput converts request to use case input
func (m *RepresentativeMonthlyGoalPayloadMapper) GetRequestToInput(id string) (*representativemonthlygoal.GetRepresentativeMonthlyGoalInput, error) {
	goalID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &representativemonthlygoal.GetRepresentativeMonthlyGoalInput{
		ID: goalID,
	}, nil
}

// GetOutputToResponse converts use case output to response
func (m *RepresentativeMonthlyGoalPayloadMapper) GetOutputToResponse(output *representativemonthlygoal.GetRepresentativeMonthlyGoalOutput) *representativemonthlygoalresponse.GetRepresentativeMonthlyGoalResponse {
	return &representativemonthlygoalresponse.GetRepresentativeMonthlyGoalResponse{
		ID:                 output.ID,
		RepresentativeUUID: output.RepresentativeUUID,
		Month:              output.Month,
		Year:               output.Year,
		Target:             output.Target,
		Realized:           output.Realized,
		Percentage:         output.Percentage,
		Remaining:          output.Remaining,
		IsTargetMet:        output.IsTargetMet,
		CreatedAt:          output.CreatedAt,
		UpdatedAt:          output.UpdatedAt,
	}
}

// UpdateRequestToInput converts request to use case input
func (m *RepresentativeMonthlyGoalPayloadMapper) UpdateRequestToInput(id string, req *representativemonthlygoalrequest.UpdateRepresentativeMonthlyGoalRequest) (*representativemonthlygoal.UpdateRepresentativeMonthlyGoalInput, error) {
	goalID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &representativemonthlygoal.UpdateRepresentativeMonthlyGoalInput{
		ID:       goalID,
		Target:   req.Target,
		Realized: req.Realized,
	}, nil
}

// UpdateOutputToResponse converts use case output to response
func (m *RepresentativeMonthlyGoalPayloadMapper) UpdateOutputToResponse(output *representativemonthlygoal.UpdateRepresentativeMonthlyGoalOutput) *representativemonthlygoalresponse.UpdateRepresentativeMonthlyGoalResponse {
	return &representativemonthlygoalresponse.UpdateRepresentativeMonthlyGoalResponse{
		ID:                 output.ID,
		RepresentativeUUID: output.RepresentativeUUID,
		Month:              output.Month,
		Year:               output.Year,
		Target:             output.Target,
		Realized:           output.Realized,
		Percentage:         output.Percentage,
		UpdatedAt:          output.UpdatedAt,
	}
}

// DeleteRequestToInput converts request to use case input
func (m *RepresentativeMonthlyGoalPayloadMapper) DeleteRequestToInput(id string) (*representativemonthlygoal.DeleteRepresentativeMonthlyGoalInput, error) {
	goalID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return &representativemonthlygoal.DeleteRepresentativeMonthlyGoalInput{
		ID: goalID,
	}, nil
}

// ListRequestToInput converts request to use case input
func (m *RepresentativeMonthlyGoalPayloadMapper) ListRequestToInput(req *representativemonthlygoalrequest.ListRepresentativeMonthlyGoalsRequest) *representativemonthlygoal.ListRepresentativeMonthlyGoalsInput {
	var representativeUUID *uuid.UUID
	if req.RepresentativeID != nil {
		id, _ := uuid.Parse(*req.RepresentativeID)
		representativeUUID = &id
	}

	// Converter mês de string (BaseQueryParams) para int, se necessário
	var month *int
	if req.GetMonth() != nil && *req.GetMonth() != "---" {
		monthVal := monthStringToInt(*req.GetMonth())
		month = &monthVal
	}

	// Desreferenciar SortBy e SortOrder se não forem nil
	var sortBy string
	if req.SortBy != nil {
		sortBy = *req.SortBy
	}

	var sortOrder string
	if req.SortOrder != nil {
		sortOrder = *req.SortOrder
	}

	return &representativemonthlygoal.ListRepresentativeMonthlyGoalsInput{
		RepresentativeUUID: representativeUUID,
		Month:              month,
		Year:               req.GetYear(),
		Page:               req.GetPage(),
		PageSize:           req.GetLimit(),
		SortBy:             sortBy,
		SortOrder:          sortOrder,
	}
}

// ListOutputToResponse converts use case output to response
func (m *RepresentativeMonthlyGoalPayloadMapper) ListOutputToResponse(output *representativemonthlygoal.ListRepresentativeMonthlyGoalsOutput) *representativemonthlygoalresponse.ListRepresentativeMonthlyGoalsResponse {
	data := make([]representativemonthlygoalresponse.RepresentativeMonthlyGoalData, len(output.Data))
	for i, item := range output.Data {
		data[i] = representativemonthlygoalresponse.RepresentativeMonthlyGoalData{
			ID:                 item.ID,
			RepresentativeUUID: item.RepresentativeUUID,
			Month:              item.Month,
			Year:               item.Year,
			Target:             item.Target,
			Realized:           item.Realized,
			Percentage:         item.Percentage,
			Remaining:          item.Remaining,
			IsTargetMet:        item.IsTargetMet,
			CreatedAt:          item.CreatedAt,
			UpdatedAt:          item.UpdatedAt,
		}
	}

	return &representativemonthlygoalresponse.ListRepresentativeMonthlyGoalsResponse{
		Data:       data,
		Total:      output.Total,
		Page:       output.Page,
		PageSize:   output.PageSize,
		TotalPages: output.TotalPages,
	}
}

// GetTableDataRequestToInput converts request to use case input
func (m *RepresentativeMonthlyGoalPayloadMapper) GetTableDataRequestToInput(req *representativemonthlygoalrequest.GetRepresentativeGoalsTableDataRequest) *representativemonthlygoal.GetRepresentativeGoalsTableDataInput {
	// Converter mês de string (BaseQueryParams) para int, se necessário
	var month *int
	if req.GetMonth() != nil && *req.GetMonth() != "---" {
		monthVal := monthStringToInt(*req.GetMonth())
		month = &monthVal
	}

	// Desreferenciar Year pois GetRepresentativeGoalsTableDataInput.Year é int (não ponteiro)
	year := 0
	if req.GetYear() != nil {
		year = *req.GetYear()
	}

	return &representativemonthlygoal.GetRepresentativeGoalsTableDataInput{
		Year:  year,
		Month: month,
	}
}

// GetTableDataOutputToResponse converts use case output to response
func (m *RepresentativeMonthlyGoalPayloadMapper) GetTableDataOutputToResponse(output *representativemonthlygoal.GetRepresentativeGoalsTableDataOutput) *representativemonthlygoalresponse.GetRepresentativeGoalsTableDataResponse {
	months := make([]representativemonthlygoalresponse.MonthData, len(output.Months))
	for i, month := range output.Months {
		months[i] = representativemonthlygoalresponse.MonthData{
			Month:       month.Month,
			MonthName:   getMonthName(month.Month),
			TargetSum:   month.TargetSum,
			RealizedSum: month.RealizedSum,
		}
	}

	rows := make([]representativemonthlygoalresponse.RepresentativeRowData, len(output.Rows))
	for i, row := range output.Rows {
		monthValues := make([]representativemonthlygoalresponse.MonthValue, len(row.MonthValues))
		for j, mv := range row.MonthValues {
			monthValues[j] = representativemonthlygoalresponse.MonthValue{
				Month:      mv.Month,
				Target:     mv.Target,
				Realized:   mv.Realized,
				Percentage: mv.Percentage,
				IsMet:      mv.IsMet,
			}
		}

		rows[i] = representativemonthlygoalresponse.RepresentativeRowData{
			RepresentativeID:   row.RepresentativeID,
			RepresentativeName: row.RepresentativeName,
			Company:            row.Company,
			Region:             row.Region,
			City:               row.City,
			MonthValues:        monthValues,
			TotalTarget:        row.TotalTarget,
			TotalRealized:      row.TotalRealized,
			TotalPercentage:    row.TotalPercentage,
		}
	}

	return &representativemonthlygoalresponse.GetRepresentativeGoalsTableDataResponse{
		Year:   output.Year,
		Months: months,
		Rows:   rows,
		Summary: representativemonthlygoalresponse.TableSummaryData{
			TotalRepresentatives: output.Summary.TotalRepresentatives,
			TotalTarget:          output.Summary.TotalTarget,
			TotalRealized:        output.Summary.TotalRealized,
			OverallPercentage:    output.Summary.OverallPercentage,
			GoalsMetCount:        output.Summary.GoalsMetCount,
			GoalsNotMetCount:     output.Summary.GoalsNotMetCount,
		},
	}
}

// getMonthName returns a Portuguese name of a month
func getMonthName(month int) string {
	months := []string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}
	if month >= 1 && month <= 12 {
		return months[month-1]
	}
	return ""
}

// monthStringToInt converte mês em formato string (JAN, FEV, etc.) para int (1-12)
func monthStringToInt(month string) int {
	monthMap := map[string]int{
		"JAN": 1,
		"FEV": 2,
		"MAR": 3,
		"ABR": 4,
		"MAI": 5,
		"JUN": 6,
		"JUL": 7,
		"AGO": 8,
		"SET": 9,
		"OUT": 10,
		"NOV": 11,
		"DEZ": 12,
	}
	if val, ok := monthMap[month]; ok {
		return val
	}
	return 0
}
