package controller

import (
	"github.com/seu-usuario/solis-backend/core/domain/entity"
	"github.com/seu-usuario/solis-backend/entrypoint/http/payload/response"
)

// KpiMapper handles conversion between entities and HTTP responses
type KpiMapper struct{}

// NewKpiMapper creates a new KpiMapper instance
func NewKpiMapper() *KpiMapper {
	return &KpiMapper{}
}

// ToKpiResponse converts entity to response
func (m *KpiMapper) ToKpiResponse(kpi *entity.KpiCategory) *response.KpiResponse {
	return &response.KpiResponse{
		ID:    kpi.ID().String(),
		Title: kpi.Title(),
		Color: kpi.Color(),
		Unit:  kpi.Unit(),
		Data:  m.ToMonthlyDataResponseList(kpi.MonthlyDatas()),
	}
}

// ToKpiResponseList converts slice of entities to slice of responses
func (m *KpiMapper) ToKpiResponseList(kpis []*entity.KpiCategory) []response.KpiResponse {
	responses := make([]response.KpiResponse, len(kpis))
	for i, kpi := range kpis {
		responses[i] = *m.ToKpiResponse(kpi)
	}
	return responses
}

// ToMonthlyDataResponse converts entity to response
func (m *KpiMapper) ToMonthlyDataResponse(monthlyData *entity.MonthlyData) *response.MonthlyDataResponse {
	var breakdown interface{}
	if monthlyData.Breakdown() != nil {
		var err error
		breakdown, err = monthlyData.GetBreakdown()
		if err != nil {
			// If we can't unmarshal breakdown, set it to nil
			breakdown = nil
		}
	}

	return &response.MonthlyDataResponse{
		ID:            monthlyData.ID().String(),
		KpiCategoryID: monthlyData.KpiCategoryID().String(),
		Month:         monthlyData.Month(),
		Realized:      monthlyData.Realized(),
		Meta:          monthlyData.Meta(),
		Breakdown:     breakdown,
	}
}

// ToMonthlyDataResponseList converts slice of entities to slice of responses
func (m *KpiMapper) ToMonthlyDataResponseList(monthlyData []*entity.MonthlyData) []response.MonthlyDataResponse {
	responses := make([]response.MonthlyDataResponse, len(monthlyData))
	for i, data := range monthlyData {
		responses[i] = *m.ToMonthlyDataResponse(data)
	}
	return responses
}

// ToKpiResponseWithMonthlyData converts entity to response with monthly data
func (m *KpiMapper) ToKpiResponseWithMonthlyData(kpi *entity.KpiCategory, monthlyData []*entity.MonthlyData) *response.KpiResponse {
	return &response.KpiResponse{
		ID:    kpi.ID().String(),
		Title: kpi.Title(),
		Color: kpi.Color(),
		Unit:  kpi.Unit(),
		Data:  m.ToMonthlyDataResponseList(monthlyData),
	}
}
