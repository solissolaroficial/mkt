package response

// KpiResponse defines the structure for KPI responses
type KpiResponse struct {
	ID        string                `json:"id"`
	Title     string                `json:"title"`
	Slug      string                `json:"slug"`
	Color     string                `json:"color"`
	Unit      string                `json:"unit"`
	CreatedBy *string               `json:"created_by,omitempty"` // ID do usuário criador (null para KPIs do sistema)
	Data      []MonthlyDataResponse `json:"data"`
}

// KpiListResponse defines the structure for KPI list responses with pagination
type KpiListResponse struct {
	Data       []KpiResponse      `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}
