package response

// KpiLogResponse defines the structure for KPI log entries
type KpiLogResponse struct {
	ID        string   `json:"id"`
	Date      string   `json:"date"`
	Timestamp string   `json:"timestamp"`
	User      string   `json:"user"`
	Month     string   `json:"month"`
	OldValue  *float64 `json:"oldValue"`
	NewValue  float64  `json:"newValue"`
	Action    string   `json:"action"`
	Context   string   `json:"context"`
}

// MonthlyDataResponse defines the structure for monthly data responses
type MonthlyDataResponse struct {
	ID            string           `json:"id"`
	KpiCategoryID string           `json:"kpi_category_id"`
	Year          int              `json:"year"` // Year (e.g., 2024, 2025)
	Month         string           `json:"month"`
	Realized      *float64         `json:"realized,omitempty"`
	Meta          *float64         `json:"meta,omitempty"`
	Breakdown     interface{}      `json:"breakdown,omitempty"`
	Logs          []KpiLogResponse `json:"logs"`
}
