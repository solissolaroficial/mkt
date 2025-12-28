package response

// KpiResponse defines the structure for KPI responses
type KpiResponse struct {
	ID    string                `json:"id"`
	Title string                `json:"title"`
	Color string                `json:"color"`
	Unit  string                `json:"unit"`
	Data  []MonthlyDataResponse `json:"data"`
}

// KpiListResponse defines the structure for KPI list responses with pagination
type KpiListResponse struct {
	Data       []KpiResponse      `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}
