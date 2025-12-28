package response

// MonthlyDataResponse defines the structure for monthly data responses
type MonthlyDataResponse struct {
	ID            string      `json:"id"`
	KpiCategoryID string      `json:"kpi_category_id"`
	Month         string      `json:"month"`
	Realized      *float64    `json:"realized,omitempty"`
	Meta          *float64    `json:"meta,omitempty"`
	Breakdown     interface{} `json:"breakdown,omitempty"`
}
